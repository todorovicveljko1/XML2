package src

import (
	"context"
	"time"

	"acc.accommodation.com/config"
	"acc.accommodation.com/pb"
	"acc.accommodation.com/src/db"
	"acc.accommodation.com/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAccommodationServiceServer
	cfg *config.Config

	acc_collection       *mongo.Collection
	prices_collection    *mongo.Collection
	available_collection *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	acc_collection := client.Database("accommodation_acc").Collection("accommodation")
	prices_collection := client.Database("accommodation_acc").Collection("prices_intervals")
	available_collection := client.Database("accommodation_acc").Collection("available_intevals")

	return &Server{
		cfg:                  cfg,
		dbClient:             client,
		acc_collection:       acc_collection,
		prices_collection:    prices_collection,
		available_collection: available_collection,
	}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) GetAccommodation(parent context.Context, dto *pb.GetAccommodationRequest) (*pb.GetAccommodationResponse, error) {

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	accId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert accommodation id")
	}

	// Find the accommodation
	var acc model.Accommodation
	err = s.acc_collection.FindOne(ctx, model.Accommodation{Id: accId}).Decode(&acc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find accommodation: %v", err)
	}

	// Find the available intervals
	var availableIntervals []model.AvailableInterval
	cursor, err := s.available_collection.Find(ctx, model.AvailableInterval{AccommodationId: accId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find available intervals: %v", err)
	}
	cursor.All(ctx, &availableIntervals)

	// Find the price intervals
	var priceIntervals []model.PriceInterval
	cursor, err = s.prices_collection.Find(ctx, model.PriceInterval{AccommodationId: accId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find price intervals: %v", err)
	}
	cursor.All(ctx, &priceIntervals)

	// Convert to proto
	accProto := acc.ToProto()
	availableIntervalsProto := make([]*pb.AvailableInterval, len(availableIntervals))
	for i, availableInterval := range availableIntervals {
		availableIntervalsProto[i] = availableInterval.ToProto()
	}
	priceIntervalsProto := make([]*pb.PriceInterval, len(priceIntervals))
	for i, priceInterval := range priceIntervals {
		priceIntervalsProto[i] = priceInterval.ToProto()
	}

	return &pb.GetAccommodationResponse{
		Accommodation:      accProto,
		AvailableIntervals: availableIntervalsProto,
		PriceIntervals:     priceIntervalsProto,
	}, nil
}
func (s *Server) CreateAccommodation(parent context.Context, dto *pb.CreateAccommodationRequest) (*pb.ResponseMessage, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert user id")
	}

	acc := model.Accommodation{
		Id:              primitive.NewObjectID(),
		Amenity:         dto.Amenity,
		DefaultPrice:    dto.DefaultPrice,
		Location:        dto.Location,
		MaxGuests:       int(dto.MaxGuests),
		MinGuests:       int(dto.MinGuests),
		Name:            dto.Name,
		PhotoURL:        dto.PhotoUrl,
		UserId:          userId,
		IsPricePerNight: dto.IsPricePerNight,
		IsManual:        dto.IsManual,
	}

	// Insert the accommodation
	res, err := s.acc_collection.InsertOne(ctx, acc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert accommodation: %v", err)
	}

	return &pb.ResponseMessage{Message: res.InsertedID.(primitive.ObjectID).Hex()}, nil

}
func (s *Server) UpdateAvailability(parent context.Context, dto *pb.UpdateAvailabilityRequest) (*pb.ResponseMessage, error) {

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	accId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert accommodation id")
	}

	// parse the dates
	startDate, err := time.Parse(time.RFC3339, dto.Availability.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse start date")
	}
	endDate, err := time.Parse(time.RFC3339, dto.Availability.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse end date")
	}

	// Save the available interval
	res, err := s.available_collection.InsertOne(ctx, model.AvailableInterval{
		Id:              primitive.NewObjectID(),
		AccommodationId: accId,
		StartDate:       startDate,
		EndDate:         endDate,
		IsAvailable:     dto.Availability.IsAvailable,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert available interval: %v", err)
	}

	return &pb.ResponseMessage{Message: res.InsertedID.(primitive.ObjectID).Hex()}, nil
}
func (s *Server) UpdatePrice(parent context.Context, dto *pb.UpdatePriceRequest) (*pb.ResponseMessage, error) {

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	accId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert accommodation id")
	}

	// parse the dates
	startDate, err := time.Parse(time.RFC3339, dto.Price.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse start date")
	}
	endDate, err := time.Parse(time.RFC3339, dto.Price.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse end date")
	}

	// Save the price interval
	res, err := s.prices_collection.InsertOne(ctx, model.PriceInterval{
		Id:              primitive.NewObjectID(),
		AccommodationId: accId,
		StartDate:       startDate,
		EndDate:         endDate,
		Price:           dto.Price.Price,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert price interval: %v", err)
	}

	return &pb.ResponseMessage{Message: res.InsertedID.(primitive.ObjectID).Hex()}, nil
}
func (s *Server) SearchAccommodations(parent context.Context, dto *pb.SearchRequest) (*pb.AccommodationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccommodations not implemented")
}
