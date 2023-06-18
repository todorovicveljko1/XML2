package src

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rating.accommodation.com/config"
	"rating.accommodation.com/pb"
	"rating.accommodation.com/src/db"
)

type Server struct {
	pb.UnimplementedRatingServiceServer
	cfg *config.Config

	host_rating_coll          *mongo.Collection
	accommodation_rating_coll *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	host_rating_coll := client.Database("accommodation_rating").Collection("host_rating")
	accommodation_rating_coll := client.Database("accommodation_rating").Collection("accommodation_rating_coll")

	return &Server{cfg: cfg, dbClient: client, host_rating_coll: host_rating_coll, accommodation_rating_coll: accommodation_rating_coll}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) HostRating(context.Context, *pb.RatingIdRequest) (*pb.RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostRating not implemented")
}
func (s *Server) AccommodationRating(context.Context, *pb.RatingIdRequest) (*pb.RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccommodationRating not implemented")
}
func (s *Server) Rate(context.Context, *pb.RateRequest) (*pb.RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rate not implemented")
}
func (s *Server) RemoveRating(context.Context, *pb.RemoveRatingRequest) (*pb.RemoveRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRating not implemented")
}
func (s *Server) GetMyRatings(context.Context, *pb.RatingIdRequest) (*pb.RatingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyRatings not implemented")
}
