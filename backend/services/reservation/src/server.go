package src

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) GetReservation(parent context.Context, dto *pb.GetReservationRequest) (*pb.Reservation, error) {
	var reservation model.Reservation

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	reservationId, err := primitive.ObjectIDFromHex(dto.ReservationId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching reservation")
	}

	err = s.res_collection.FindOne(ctx, bson.M{"_id": reservationId}).Decode(&reservation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "Reservation not found")
		}
		return nil, status.Error(codes.Internal, "Error while fetching reservation")
	}

	return reservation.ConvertToPbReservation(), nil
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
		StartDate:       dto.StartDate.AsTime(),
		EndDate:         dto.EndDate.AsTime(),
		Status:          "PENNDING",
		Price:           dto.Price,
	}

	//Check if there are reservations in that interval
	affectedReservations, err := s.CheckReservationIntervals(ctx, reservation)
	if err != nil {
		log.Println("Cannot get affected intervals")
		return nil, err
	}

	for _, i_reservation := range affectedReservations {
		if i_reservation.Status == "APPROVED" {
			return nil, status.Error(codes.Internal, "Already exists reservation at the same time")
		}
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
func (s *Server) RejectReservation(parent context.Context, dto *pb.GetReservationRequest) (*pb.ReservationStatus, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(dto.ReservationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid reservation id")
	}

	res, err := s.res_collection.UpdateOne(ctx, bson.M{
		"_id":    id,
		"status": "PENNDING",
	}, bson.M{
		"status": "REJECTED",
	})

	// handle error and see res if somthing is updated
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while updating reservation")
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}

	return &pb.ReservationStatus{
		Status: "REJECTED",
	}, nil
}
func (s *Server) CancelReservation(parent context.Context, dto *pb.GetReservationRequest) (*pb.ReservationStatus, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(dto.ReservationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid reservation id")
	}

	var reservation model.Reservation
	err = s.res_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&reservation)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}

	if reservation.StartDate.Time().Before(time.Now()) {
		return nil, status.Error(codes.DeadlineExceeded, "Canceling reservation is unavailable due to start date already passing.")
	}

	res, err := s.res_collection.UpdateOne(ctx, bson.M{
		"_id":    id,
		"status": "APPROVED",
	}, bson.M{
		"status": "CANCELED",
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "Error while updating reservation")
	}

	if res.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Reservation not found")
	}

	return &pb.ReservationStatus{
		Status: "CANCELED",
	}, nil
}

func (a *Server) CheckReservationIntervals(ctx context.Context, reservation model.Reservation) ([]*model.Reservation, error) {

	affectedIntervals := []*model.Reservation{}
	cursor, err := a.res_collection.Find(ctx, bson.M{
		"accommodation_id": reservation.AccommodationId,
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": reservation.StartDate}},
				{"start_date": bson.M{"$lte": reservation.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": reservation.StartDate}},
				{"end_date": bson.M{"$lte": reservation.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": reservation.StartDate}},
				{"end_date": bson.M{"$gte": reservation.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": reservation.StartDate}},
				{"end_date": bson.M{"$lte": reservation.EndDate}},
			}},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		reservations := &model.Reservation{}
		if err = cursor.Decode(reservations); err != nil {
			return nil, err
		}
		affectedIntervals = append(affectedIntervals, reservations)
	}
	return affectedIntervals, nil
}
