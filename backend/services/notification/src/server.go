package src

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"notification.accommodation.com/config"
	"notification.accommodation.com/pb"
	"notification.accommodation.com/src/db"
	"notification.accommodation.com/src/model"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
	cfg *config.Config

	notification_coll *mongo.Collection
	setting_coll      *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	notification_coll := client.Database("accommodation_notif").Collection("notification")
	setting_coll := client.Database("accommodation_notif").Collection("setting")

	return &Server{cfg: cfg, dbClient: client, notification_coll: notification_coll, setting_coll: setting_coll}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) isAllowedNotification(ctx context.Context, userId primitive.ObjectID, notificationType string) (bool, error) {
	var settings model.NotificationSetting
	// get user settings
	err := s.setting_coll.FindOne(ctx, bson.M{"user_id": userId}).Decode(&settings)
	if err != nil {
		// if user settings not found, create new one
		if err == mongo.ErrNoDocuments {
			settings = model.NotificationSetting{
				Id:     primitive.NewObjectID(),
				UserId: userId,
				Settings: []model.OneNotificationSetting{
					{
						Type:    notificationType,
						Allowed: true,
					},
				},
			}
			_, err := s.setting_coll.InsertOne(ctx, settings)
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}

	// check if notification type is allowed
	for _, setting := range settings.Settings {
		if setting.Type == notificationType {
			// if found, return allowed
			return setting.Allowed, nil
		}
	}
	// allowed by default
	return true, nil
}

func (s *Server) SendNotification(parent context.Context, dto *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// parse user id
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}

	// check if notification type is allowed
	allowed, err := s.isAllowedNotification(ctx, userId, dto.Type)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	if !allowed {
		return &pb.SendNotificationResponse{Success: false}, nil
	}

	notification := model.Notification{
		Id:         primitive.NewObjectID(),
		Type:       dto.Type,
		Body:       dto.Body,
		ResourceId: dto.ResourceId,
		UserId:     userId,
		IsRead:     false,
		CreatedAt:  time.Now(),
	}

	_, err = s.notification_coll.InsertOne(ctx, notification)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pb.SendNotificationResponse{Success: true}, nil
}
func (s *Server) GetNotifications(parent context.Context, dto *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {

	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// parse user id
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}
	// notifications for user
	filter := bson.M{"user_id": userId}
	// do we need only unread notifications
	if !dto.All {
		filter["is_read"] = false
	}
	findOptions := options.FindOptions{Sort: bson.M{"created_at": -1}}

	// get notifications
	cursor, err := s.notification_coll.Find(ctx, filter, &findOptions)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	// parse notifications
	defer cursor.Close(ctx)
	var notifications []*pb.Notification
	for cursor.Next(ctx) {
		var notification model.Notification
		err := cursor.Decode(&notification)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error")
		}
		notifications = append(notifications, notification.ToProto())
	}

	return &pb.GetNotificationResponse{Notifications: notifications}, nil
}
func (s *Server) MarkNotificationAsRead(parent context.Context, dto *pb.MarkNotificationAsReadRequest) (*pb.MarkNotificationAsReadResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// parse user id
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}
	// parse notification id
	notificationId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid notification id")
	}

	// mark notification as read
	res, err := s.notification_coll.UpdateOne(ctx, bson.M{"_id": notificationId, "user_id": userId}, bson.M{"$set": bson.M{"is_read": true}})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	if res.MatchedCount == 0 {
		return &pb.MarkNotificationAsReadResponse{Success: false}, nil
	}

	return &pb.MarkNotificationAsReadResponse{Success: true}, nil

}
func (s *Server) ChangeNotifcationSettings(parent context.Context, dto *pb.ChangeNotifcationSettingsRequest) (*pb.ChangeNotifcationSettingsResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// parse user id
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}

	// convert dto.Settings to []model.OneNotificationSetting
	notifications := make([]model.OneNotificationSetting, len(dto.Settings))
	for i, setting := range dto.Settings {
		notifications[i] = model.OneNotificationSetting{
			Type:    setting.Type,
			Allowed: setting.Enabled,
		}
	}

	_, err = s.setting_coll.UpdateOne(ctx,
		bson.M{"user_id": userId},
		bson.M{"$set": bson.M{"settings": notifications}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &pb.ChangeNotifcationSettingsResponse{Success: true}, nil

}

func (s *Server) GetNotificationSettings(parent context.Context, dto *pb.GetNotificationSettingsRequest) (*pb.GetNotificationSettingsResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// parse user id
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}

	var settings model.NotificationSetting
	err = s.setting_coll.FindOne(ctx, bson.M{"user_id": userId}).Decode(&settings)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	// convert []model.OneNotificationSetting to []*pb.NotificationSetting
	notifications := make([]*pb.NotificationSetting, len(settings.Settings))
	for i, setting := range settings.Settings {
		notifications[i] = &pb.NotificationSetting{
			Type:    setting.Type,
			Enabled: setting.Allowed,
		}
	}

	return &pb.GetNotificationSettingsResponse{Settings: notifications}, nil
}
