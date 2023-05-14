package src

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reservation.accommodation.com/config"
	"reservation.accommodation.com/pb"
	"reservation.accommodation.com/src/db"
	"reservation.accommodation.com/src/model"
)

type Server struct {
	pb.UnimplementedReservationServiceServer

	cfg *config.Config

	res_collection *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	res_collection := client.Database("accommodation_res").Collection("reservation")

	return &Server{cfg: cfg, dbClient: client, res_collection: res_collection}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) GetReservation(context.Context, *pb.GetReservationRequest) (*pb.Reservation, error) {
	return &pb.Reservation{
		Id:              "1",
		UserId:          "1",
		AccommodationId: "1",
		StartDate:       timestamppb.Now(),
		EndDate:         timestamppb.Now(),
		Status:          "CREATED",
		Price:           1000,
	}, nil
	//return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (s *Server) CreateReservation(parent context.Context, dto *pb.CreateReservationRequest) (*pb.Reservation, error) {
	// Create timeout context
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	accommodationId, err := primitive.ObjectIDFromHex(dto.AccommodationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid accommodation id")
	}

	// Create reservation
	reservation := model.Reservation{
		Id:              primitive.NewObjectID(),
		UserId:          userId,
		AccommodationId: accommodationId,
		StartDate:       primitive.NewDateTimeFromTime(dto.StartDate.AsTime()),
		EndDate:         primitive.NewDateTimeFromTime(dto.EndDate.AsTime()),
		Status:          "PENNDING",
		Price:           dto.Price,
	}

	// Insert reservation
	_, err = s.res_collection.InsertOne(ctx, reservation)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while inserting reservation")
	}

	return reservation.ConvertToPbReservation(), nil
}

func (s *Server) ApproveReservation(parent context.Context, dto *pb.GetReservationRequest) (*pb.ReservationStatus, error) {
	var toReserve model.Reservation
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(dto.ReservationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid reservation id")
	}

	err = s.res_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&toReserve)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}

	cursor, err := s.res_collection.Find(parent, bson.M{
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"_id": bson.M{"$ne": id}},
				{"start_date": bson.M{"$gte": toReserve.StartDate}},
				{"start_date": bson.M{"$lte": toReserve.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"_id": bson.M{"$ne": id}},
				{"end_date": bson.M{"$gte": toReserve.StartDate}},
				{"end_date": bson.M{"$lte": toReserve.EndDate}},
			}},
		},
	})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		reservationCancel := &model.Reservation{}
		if err = cursor.Decode(reservationCancel); err != nil {
			return nil, err
		}

		_, err := s.res_collection.UpdateOne(ctx, bson.M{
			"_id":    reservationCancel.Id,
			"status": "PENNDING",
		}, bson.M{
			"status": "REJECTED",
		})
		if err != nil {
			return nil, status.Error(codes.NotFound, "Reservation that should be cancelled, could not be cancelled")
		}
	}

	res, err := s.res_collection.UpdateOne(ctx, bson.M{
		"_id":    id,
		"status": "PENNDING",
	}, bson.M{
		"status": "APPROVED",
	})

	// handle error and see res if somthing is updated
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while updating reservation")
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}

	return &pb.ReservationStatus{
		Status: "APPROVED",
	}, nil
}
func (s *Server) RejectReservation(context.Context, *pb.GetReservationRequest) (*pb.ReservationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectReservation not implemented")
}
func (s *Server) CancelReservation(context.Context, *pb.GetReservationRequest) (*pb.ReservationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
