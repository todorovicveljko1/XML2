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

	// parse dates
	startDate, err := time.Parse(time.RFC3339, dto.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse start date")
	}
	startDate = startDate.Truncate(24 * time.Hour)
	endDate, err := time.Parse(time.RFC3339, dto.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse end date")
	}
	endDate = endDate.Truncate(24 * time.Hour)

	// Create reservation
	reservation := model.Reservation{
		Id:              primitive.NewObjectID(),
		UserId:          userId,
		AccommodationId: accommodationId,
		Status:          "PENDING",
		Price:           dto.Price,
		StartDate:       startDate,
		EndDate:         endDate,
		NumberOfGuests:  dto.NumberOfGuests,
	}

	//Check if there are reservations in that interval
	affectedIntervals, err := s.checkReservationIntervals(ctx, reservation)
	if err != nil {
		log.Println("Cannot get affected intervals")
		return nil, err
	}

	if len(affectedIntervals) != 0 {
		return nil, status.Error(codes.Internal, "Already exists reservation at the same time")
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
		"accommodation_id": toReserve.AccommodationId,
		"_id":              bson.M{"$ne": id},
		"status":           "PENDING",
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": toReserve.StartDate}},
				{"start_date": bson.M{"$lte": toReserve.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": toReserve.StartDate}},
				{"end_date": bson.M{"$lte": toReserve.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": toReserve.StartDate}},
				{"end_date": bson.M{"$gte": toReserve.EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": toReserve.StartDate}},
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
			"status": "PENDING",
		}, bson.M{
			"$set": bson.M{
				"status": "AUTO_REJECTED",
			},
		})
		if err != nil {
			return nil, status.Error(codes.NotFound, "Reservation that should be cancelled, could not be cancelled")
		}
	}

	res, err := s.res_collection.UpdateOne(ctx, bson.M{
		"_id":    id,
		"status": "PENDING",
	}, bson.M{

		"$set": bson.M{
			"status": "APPROVED",
		},
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
		"status": "PENDING",
	}, bson.M{
		"$set": bson.M{
			"status": "REJECTED",
		},
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

	if reservation.StartDate.Before(time.Now()) {
		return nil, status.Error(codes.InvalidArgument, "Canceling reservation is unavailable due to start date already passing.")
	}

	res, err := s.res_collection.UpdateOne(ctx, bson.M{
		"_id":    id,
		"status": bson.M{"$in": bson.A{"APPROVED", "PENDING"}},
	}, bson.M{
		"$set": bson.M{
			"status": "CANCELED",
		},
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

func (s *Server) GetReservationsForGuest(parent context.Context, dto *pb.IdRequest) (*pb.ReservationList, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	cursor, err := s.res_collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching reservations")
	}
	defer cursor.Close(ctx)

	reservations := []*pb.Reservation{}
	for cursor.Next(ctx) {
		reservation := &model.Reservation{}
		if err = cursor.Decode(reservation); err != nil {
			return nil, status.Error(codes.Internal, "Error while decoding reservations")
		}
		reservations = append(reservations, reservation.ConvertToPbReservation())
	}

	return &pb.ReservationList{
		Reservations: reservations,
	}, nil
}
func (s *Server) GetReservationsForAccommodation(parent context.Context, dto *pb.IdRequest) (*pb.ReservationList, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	accommodationId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid accommodation id")
	}

	cursor, err := s.res_collection.Find(ctx, bson.M{"accommodation_id": accommodationId})
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching reservations")
	}
	defer cursor.Close(ctx)

	reservations := []*pb.Reservation{}
	for cursor.Next(ctx) {
		reservation := &model.Reservation{}
		if err = cursor.Decode(reservation); err != nil {
			return nil, status.Error(codes.Internal, "Error while decoding reservations")
		}
		reservations = append(reservations, reservation.ConvertToPbReservation())
	}

	return &pb.ReservationList{
		Reservations: reservations,
	}, nil

}

func (s *Server) FilterOutTakenAccommodations(parent context.Context, dto *pb.FilterTakenAccommodationsRequest) (*pb.IdList, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	// parse dates
	StartDate, err := time.Parse(time.RFC3339, dto.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse start date")
	}
	StartDate = StartDate.Truncate(24 * time.Hour)

	EndDate, err := time.Parse(time.RFC3339, dto.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse end date")
	}
	EndDate = EndDate.Truncate(24 * time.Hour)

	// convert ids to object ids
	var accommodationIds []primitive.ObjectID

	for _, id := range dto.AccommodationIds {
		accommodationId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid accommodation id")
		}
		accommodationIds = append(accommodationIds, accommodationId)
	}

	cursor, err := s.res_collection.Find(ctx, bson.M{
		"accommodation_id": bson.M{"$in": accommodationIds},
		"status":           "APPROVED",
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": StartDate}},
				{"start_date": bson.M{"$lte": EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": StartDate}},
				{"end_date": bson.M{"$lte": EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": StartDate}},
				{"end_date": bson.M{"$gte": EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": StartDate}},
				{"end_date": bson.M{"$lte": EndDate}},
			}},
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	takenAccommodationIds := []string{}
	for cursor.Next(ctx) {
		reservation := &model.Reservation{}
		if err = cursor.Decode(reservation); err != nil {
			return nil, err
		}
		takenAccommodationIds = append(takenAccommodationIds, reservation.AccommodationId.Hex())
	}

	log.Println(takenAccommodationIds)

	// Filter out taken accommodations
	availableAccommodationIds := []string{}
	for _, accommodationId := range accommodationIds {
		taken := false
		for _, takenAccommodationId := range takenAccommodationIds {
			if accommodationId.Hex() == takenAccommodationId {
				taken = true
				break
			}
		}
		if !taken {
			availableAccommodationIds = append(availableAccommodationIds, accommodationId.Hex())
		}
	}

	return &pb.IdList{
		Ids: availableAccommodationIds,
	}, nil

}

func (s *Server) HasActiveReservationInInterval(parent context.Context, dto *pb.IntervalRequest) (*pb.BoolResponse, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	// parse dates
	StartDate, err := time.Parse(time.RFC3339, dto.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse start date")
	}
	StartDate = StartDate.Truncate(24 * time.Hour)

	EndDate, err := time.Parse(time.RFC3339, dto.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse end date")
	}
	EndDate = EndDate.Truncate(24 * time.Hour)

	// parse accommodation id
	accommodationId, err := primitive.ObjectIDFromHex(dto.AccommodationId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid accommodation id")
	}

	cursor, err := s.res_collection.Find(ctx, bson.M{
		"accommodation_id": accommodationId,
		"status":           bson.M{"$in": bson.A{"APPROVED", "PENDING"}},
		"$or": bson.A{
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": StartDate}},
				{"start_date": bson.M{"$lte": EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"end_date": bson.M{"$gte": StartDate}},
				{"end_date": bson.M{"$lte": EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$lte": StartDate}},
				{"end_date": bson.M{"$gte": EndDate}},
			}},
			bson.M{"$and": []bson.M{
				{"start_date": bson.M{"$gte": StartDate}},
				{"end_date": bson.M{"$lte": EndDate}},
			}},
		},
	})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	hasActiveReservation := false
	for cursor.Next(ctx) {
		hasActiveReservation = true
		break
	}

	return &pb.BoolResponse{
		Value: hasActiveReservation,
	}, nil
}

func (s *Server) HasGuestActiveReservationInFuture(parent context.Context, dto *pb.IdRequest) (*pb.BoolResponse, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	cursor, err := s.res_collection.Find(ctx, bson.M{
		"user_id":  userId,
		"status":   bson.M{"$in": bson.A{"APPROVED", "PENDING"}},
		"end_date": bson.M{"$gte": time.Now()},
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching reservations")
	}

	defer cursor.Close(ctx)

	hasActiveReservation := false
	for cursor.Next(ctx) {
		hasActiveReservation = true
		break
	}

	return &pb.BoolResponse{
		Value: hasActiveReservation,
	}, nil

}
func (s *Server) HasHostActiveReservationInFuture(parent context.Context, dto *pb.IdList) (*pb.BoolResponse, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	// convert ids to object ids
	var accommodationIds []primitive.ObjectID

	for _, id := range dto.Ids {
		accommodationId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid accommodation id")
		}
		accommodationIds = append(accommodationIds, accommodationId)
	}

	cursor, err := s.res_collection.Find(ctx, bson.M{
		"accommodation_id": bson.M{"$in": accommodationIds},
		"status":           bson.M{"$in": bson.A{"APPROVED", "PENDING"}},
		"end_date":         bson.M{"$gte": time.Now()},
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching reservations")
	}

	defer cursor.Close(ctx)

	hasActiveReservation := false
	for cursor.Next(ctx) {
		hasActiveReservation = true
		break
	}

	return &pb.BoolResponse{
		Value: hasActiveReservation,
	}, nil

}

// PRIVATE FUNCTIONS
func (a *Server) checkReservationIntervals(ctx context.Context, reservation model.Reservation) ([]*model.Reservation, error) {

	affectedIntervals := []*model.Reservation{}
	cursor, err := a.res_collection.Find(ctx, bson.M{
		"accommodation_id": reservation.AccommodationId,
		"status":           "APPROVED",
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
