package src

import (
	"context"
	"fmt"
	"time"

	"acc.accommodation.com/config"
	"acc.accommodation.com/pb"
	"acc.accommodation.com/src/db"
	"acc.accommodation.com/src/helper"
	"acc.accommodation.com/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// *PREPRAVITI client.Database("Accommodations") u nesto trece.*
type Server struct {
	pb.UnimplementedAccommodationServiceServer
	cfg *config.Config

	acc_collection      *mongo.Collection
	price_collection    *mongo.Collection
	interval_collection *mongo.Collection
	dbClient            *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)
	//*PREPRAVITI client.Database("Accommodations") u nesto trece.*
	acc_collection := client.Database("accommodation_acc").Collection("accommodation")
	prices_collection := client.Database("accommodation_acc").Collection("prices")
	unavailable_collection := client.Database("accommodation_acc").Collection("unavailable_intevals")
	return &Server{cfg: cfg, dbClient: client, acc_collection: acc_collection, price_collection: prices_collection, interval_collection: unavailable_collection}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) GetAccommodation(parent context.Context, dto *pb.GetAccommodationRequest) (*pb.GetAccommodationResponse, error) {
	var accommodation model.Accommodation
	accommodationID, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching user")
	}

	err = s.acc_collection.FindOne(parent, bson.M{"_id": accommodationID}).Decode(&accommodation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "Accommodation not found")
		}
		return nil, status.Error(codes.Internal, "Error while fetching accommodation")
	}
	fmt.Println("found UserID: ", accommodation.UserId)
	//unavailableDates := make([]*primitive.DateTime, 1)
	datePrices := make([]*pb.DatePrice, 0)

	amenities := make([]pb.Amenity, len(accommodation.Amenity))
	for i, a := range accommodation.Amenity {
		amenities[i] = pb.Amenity(a)
	}

	response := &pb.GetAccommodationResponse{
		Accomodation: &pb.Accommodation{
			Id:           accommodation.Id.Hex(),
			Name:         accommodation.Name,
			Location:     accommodation.Location,
			Amenity:      amenities,
			PhotoUrl:     accommodation.PhotoURL,
			MaxGuests:    int32(accommodation.MaxGuests),
			MinGuests:    int32(accommodation.MinGuests),
			DefaultPrice: accommodation.DefaultPrice,
			UserId:       accommodation.UserId.Hex(),
		},
		UnavailableDates: nil,
		DatePrice:        datePrices,
	}

	return response, nil
}

func (s *Server) AddAccommodation(parent context.Context, dto *pb.AddAccommodationRequest) (*pb.AddAccommodationResponse, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	// get user from metadata
	md, ok := metadata.FromIncomingContext(parent)
	if !ok {
		return nil, status.Errorf(codes.Internal, "failed to get metadata")
	}
	user := md.Get("user")[0]

	userId, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert user id")
	}

	temp := make([]model.Amenity, len(dto.Amenity))
	for i, v := range dto.Amenity {
		temp[i] = model.Amenity(v)
	}

	acc := model.Accommodation{
		Id:           primitive.NewObjectID(),
		Amenity:      temp,
		DefaultPrice: dto.DefaultPrice,
		Location:     dto.Location,
		MaxGuests:    int(dto.MaxGuests),
		MinGuests:    int(dto.MinGuests),
		Name:         dto.Name,
		PhotoURL:     dto.PhotoUrl,
		UserId:       userId,
	}

	// Insert the accommodation
	_, err = s.acc_collection.InsertOne(ctx, acc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert accommodation: %v", err)
	}

	return &pb.AddAccommodationResponse{Success: true}, nil

}
func (s *Server) UpdateAvailability(parent context.Context, dto *pb.UpdateAvailabilityRequest) (*pb.UpdateAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvailability not implemented")
}
func (s *Server) UpdatePrice(parent context.Context, dto *pb.UpdatePriceRequest) (*pb.UpdatePriceResponse, error) {
	// Convert the ID string to an ObjectID
	id, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error parsing ID: %v", err)
	}

	// Create a new Price object
	days := helper.GetDaysInInterval(dto.Price.StartDate.AsTime(), dto.Price.EndDate.AsTime())
	for _, day := range days {
		price := model.Price{
			AccommodationId: id,
			Date:            primitive.NewDateTimeFromTime(day),
			PricePerNight:   dto.Price.PricePerNight,
		}
		s.price_collection.UpdateOne(
			parent, bson.M{"accommodation_id": price.AccommodationId, "date": price.Date},
			bson.M{
				"$set": bson.M{"price_per_night": price.PricePerNight},
				"$setOnInsert": bson.M{
					"accommodation_id": price.AccommodationId,
					"date":             price.Date,
					"price_per_night":  price.PricePerNight,
					"_id":              primitive.NewObjectID(),
				},
			})
	}
	return &pb.UpdatePriceResponse{Success: true}, nil
}

func (s *Server) SearchAccommodations(parent context.Context, dto *pb.SearchRequest) (*pb.AccommodationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccommodations not implemented")
}
