package src

import (
	"context"
	"log"
	"time"

	"acc.accommodation.com/config"
	"acc.accommodation.com/pb"
	"acc.accommodation.com/src/db"
	"acc.accommodation.com/src/manager"
	"acc.accommodation.com/src/model"
	"acc.accommodation.com/src/saga"
	"go.mongodb.org/mongo-driver/bson"
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

	available_interval_manager *manager.AvailableIntervalManager
	price_interval_manager     *manager.PriceIntervalManager

	dbClient *mongo.Client

	SAGA *saga.SAGA
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	acc_collection := client.Database("accommodation_acc").Collection("accommodation")
	prices_collection := client.Database("accommodation_acc").Collection("prices_intervals")
	available_collection := client.Database("accommodation_acc").Collection("available_intevals")

	available_interval_manager := manager.NewAvailableIntervalManager(available_collection, client)
	price_interval_manager := manager.NewPriceIntervalManager(prices_collection, client)

	SAGA, err := saga.CreateSAGA(cfg, acc_collection)
	if err != nil {
		return nil, err
	}

	return &Server{
		cfg:                  cfg,
		dbClient:             client,
		acc_collection:       acc_collection,
		prices_collection:    prices_collection,
		available_collection: available_collection,

		available_interval_manager: available_interval_manager,
		price_interval_manager:     price_interval_manager,

		SAGA: SAGA,
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
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert accommodation id")
	}
	// Find the accommodation
	var acc model.Accommodation
	err = s.acc_collection.FindOne(ctx, bson.M{"_id": accId}).Decode(&acc)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find accommodation: %v", err)
	}
	// Check if the accommodation is deleted
	if acc.DeletedAt != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find accommodation: %v", err)
	}
	// Find the available intervals
	availableIntervals, err := s.available_interval_manager.GetAvailableIntervalsByAccommodationId(ctx, accId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find available intervals: %v", err)
	}
	// Find the price intervals
	priceIntervals, err := s.price_interval_manager.GetPriceIntervalsByAccommodationId(ctx, accId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find price intervals: %v", err)
	}

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

func (s *Server) UpdateAccommodation(parent context.Context, dto *pb.Accommodation) (*pb.ResponseMessage, error) {

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	accId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert accommodation id")
	}
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert user id")
	}

	// Find the accommodation

	res, err := s.acc_collection.UpdateOne(ctx, bson.M{"_id": accId, "user_id": userId}, bson.M{"$set": bson.M{
		"amenity":            dto.Amenity,
		"default_price":      dto.DefaultPrice,
		"location":           dto.Location,
		"max_guests":         dto.MaxGuests,
		"min_guests":         dto.MinGuests,
		"name":               dto.Name,
		"photo_url":          dto.PhotoUrl,
		"is_price_per_night": dto.IsPricePerNight,
		"is_manual":          dto.IsManual,
	}})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find accommodation: %v", err)
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "failed to find accommodation: %v", err)
	}

	return &pb.ResponseMessage{Message: accId.Hex()}, nil
}

func (s *Server) AddAccommodationAvailability(parent context.Context, dto *pb.AddAvailabilityRequest) (*pb.ResponseMessage, error) {

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
	startDate = startDate.Truncate(24 * time.Hour)
	endDate, err := time.Parse(time.RFC3339, dto.Availability.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse end date")
	}
	endDate = endDate.Truncate(24 * time.Hour)

	res, err := s.available_interval_manager.AddAvailableInterval(ctx, &model.AvailableInterval{
		Id:              primitive.NewObjectID(),
		AccommodationId: accId,
		StartDate:       startDate,
		EndDate:         endDate,
		IsAvailable:     dto.Availability.IsAvailable,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert available interval: %v", err)
	}

	return &pb.ResponseMessage{Message: res.(string)}, nil
}
func (s *Server) AddAccommodationPrice(parent context.Context, dto *pb.AddPriceRequest) (*pb.ResponseMessage, error) {

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
	startDate = startDate.Truncate(24 * time.Hour)
	endDate, err := time.Parse(time.RFC3339, dto.Price.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse end date")
	}
	endDate = endDate.Truncate(24 * time.Hour)

	// Save the price interval
	res, err := s.price_interval_manager.AddPriceInterval(ctx, &model.PriceInterval{
		Id:              primitive.NewObjectID(),
		AccommodationId: accId,
		StartDate:       startDate,
		EndDate:         endDate,
		Price:           dto.Price.Price,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert price interval: %v", err)
	}

	return &pb.ResponseMessage{Message: res.(primitive.ObjectID).Hex()}, nil
}
func (s *Server) SearchAccommodations(parent context.Context, dto *pb.SearchRequest) (*pb.AccommodationList, error) {

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	// parser dates
	startDate, err := time.Parse(time.RFC3339, *dto.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse start date")
	}
	startDate = startDate.Truncate(24 * time.Hour)
	endDate, err := time.Parse(time.RFC3339, *dto.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse end date")
	}
	endDate = endDate.Truncate(24 * time.Hour)
	// create bson filter
	filter := bson.M{
		"deleted_at": bson.M{"$exists": false},
	}

	if dto.Location != nil && *dto.Location != "" {
		// regex contains
		filter["location"] = primitive.Regex{Pattern: *dto.Location, Options: "i"}

	}
	if dto.NumGuests != nil && *dto.NumGuests != 0 {
		filter["min_guests"] = bson.M{"$lte": *dto.NumGuests}
		filter["max_guests"] = bson.M{"$gte": *dto.NumGuests}
	}
	if dto.Amenity != nil && len(dto.Amenity) != 0 && dto.Amenity[0] != "" {
		// filter so that all amenities are in the array
		filter["amenity"] = bson.M{"$all": dto.Amenity}
	}
	if dto.ShowMy {
		userId, err := primitive.ObjectIDFromHex(dto.UserId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to convert user id")
		}
		filter["user_id"] = userId
	}
	// log filter
	log.Println(filter)
	// Find the accommodations
	cursor, err := s.acc_collection.Find(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find accommodations: %v", err)
	}

	var accommodations []*model.Accommodation
	// read cursor
	if err = cursor.All(ctx, &accommodations); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read accommodations: %v", err)
	}

	if dto.ShowMy {
		// convert to proto
		accommodationsProto := make([]*pb.Accommodation, len(accommodations))
		for i, accommodation := range accommodations {
			accommodationsProto[i] = accommodation.ToProto()
		}
		return &pb.AccommodationList{Accommodations: accommodationsProto}, nil
	}

	if len(accommodations) == 0 {
		return &pb.AccommodationList{Accommodations: []*pb.Accommodation{}}, nil
	}

	// filter accommodations by available intervals
	var accommodationIds []primitive.ObjectID
	for _, accommodation := range accommodations {
		accommodationIds = append(accommodationIds, accommodation.Id)
	}
	accommodationIds, err = s.available_interval_manager.FilterAccommodationIdsByAvailableIntervals(ctx, accommodationIds, startDate, endDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to filter accommodations: %v", err)
	}

	// select only accommodations with ids
	var accommodationsFiltered []*model.Accommodation
	for _, accommodation := range accommodations {
		for _, accommodationId := range accommodationIds {
			if accommodation.Id == accommodationId {
				accommodationsFiltered = append(accommodationsFiltered, accommodation)
			}
		}
	}

	// calculate price
	accommodationsFiltered, err = s.price_interval_manager.PriceAccommodationByPriceIntervals(ctx, accommodationsFiltered, startDate, endDate, int(*dto.NumGuests))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to calculate price: %v", err)
	}

	// convert to proto
	accommodationsProto := make([]*pb.Accommodation, len(accommodationsFiltered))
	for i, accommodation := range accommodationsFiltered {
		accommodationsProto[i] = accommodation.ToProto()
	}
	return &pb.AccommodationList{Accommodations: accommodationsProto}, nil

}

func (s *Server) DeleteAccommodation(parent context.Context, dto *pb.GetAccommodationRequestWithUser) (*pb.ResponseMessage, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	accId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert accommodation id")
	}
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert user id")
	}

	// Find the accommodation
	res, err := s.acc_collection.DeleteOne(ctx, bson.M{"_id": accId, "user_id": userId})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find accommodation: %v", err)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "failed to find accommodation: %v", err)
	}

	return &pb.ResponseMessage{Message: accId.Hex()}, nil
}

func (s *Server) GetAccommodationsForHost(parent context.Context, dto *pb.GetAccommodationRequest) (*pb.AccommodationList, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert user id")
	}

	// Find the accommodations
	cursor, err := s.acc_collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find accommodations: %v", err)
	}

	var accommodations []*model.Accommodation
	// read cursor
	if err = cursor.All(ctx, &accommodations); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read accommodations: %v", err)
	}

	// convert to proto
	accommodationsProto := make([]*pb.Accommodation, len(accommodations))
	for i, accommodation := range accommodations {
		accommodationsProto[i] = accommodation.ToProto()
	}

	return &pb.AccommodationList{Accommodations: accommodationsProto}, nil
}
