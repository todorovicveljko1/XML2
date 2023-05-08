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
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAccommodationServiceServer
	cfg *config.Config

	acc_collection *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	acc_collection := client.Database("accommodation_acc").Collection("accommodation")

	return &Server{cfg: cfg, dbClient: client, acc_collection: acc_collection}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) GetAccommodation(parent context.Context, dto *pb.GetAccommodationRequest) (*pb.GetAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodation not implemented")
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
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrice not implemented")
}
func (s *Server) SearchAccommodations(parent context.Context, dto *pb.SearchRequest) (*pb.AccommodationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccommodations not implemented")
}
