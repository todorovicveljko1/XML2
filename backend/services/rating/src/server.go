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
	"rating.accommodation.com/config"
	"rating.accommodation.com/pb"
	"rating.accommodation.com/src/db"
	"rating.accommodation.com/src/db/model"
)

type Server struct {
	pb.UnimplementedRatingServiceServer
	cfg *config.Config

	rating_coll *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	rating_coll := client.Database("accommodation_rating").Collection("rating")

	return &Server{cfg: cfg, dbClient: client, rating_coll: rating_coll}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) HostRating(parent context.Context, dto *pb.RatingIdRequest) (*pb.RatingResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// get object id from dto
	hostId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}

	// get ratings
	cursor, err := s.rating_coll.Find(ctx, bson.M{"host_id": hostId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get ratings")
	}

	// get ratings from cursor
	var ratings []*model.Rating
	if err = cursor.All(ctx, &ratings); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get ratings")
	}

	// calculate average rating
	var sum float64 = 0
	for _, rating := range ratings {
		sum += float64(rating.HostRating)
	}
	avg := sum / float64(len(ratings))

	return &pb.RatingResponse{Rating: avg}, nil

}
func (s *Server) AccommodationRating(parent context.Context, dto *pb.RatingIdRequest) (*pb.RatingResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// get object id from dto
	accommodationId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}

	// get ratings
	cursor, err := s.rating_coll.Find(ctx, bson.M{"accommodation_id": accommodationId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get ratings")
	}

	// get ratings from cursor
	var ratings []*model.Rating
	if err = cursor.All(ctx, &ratings); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get ratings")
	}

	// calculate average rating
	var sum float64 = 0
	for _, rating := range ratings {
		sum += float64(rating.AccommodationRating)
	}
	avg := sum / float64(len(ratings))

	return &pb.RatingResponse{Rating: avg}, nil
}
func (s *Server) Rate(parent context.Context, dto *pb.RateRequest) (*pb.RateResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// get object id from dto
	accommodationId, err := primitive.ObjectIDFromHex(dto.AccommodationId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}
	hostId, err := primitive.ObjectIDFromHex(dto.HostId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}
	reservationId, err := primitive.ObjectIDFromHex(dto.ReservationId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}

	// create rating object
	rating := &model.Rating{
		ID:                  primitive.NewObjectID(),
		AccommodationId:     accommodationId,
		HostId:              hostId,
		UserId:              userId,
		ReservationId:       reservationId,
		HostRating:          dto.HostRating,
		AccommodationRating: dto.AccommodationRating,
	}

	// upsert rating
	res, err := s.rating_coll.UpdateOne(ctx, bson.M{"reservation_id": reservationId}, bson.M{"$set": rating}, &options.UpdateOptions{Upsert: &[]bool{true}[0]})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot insert rating")
	}

	// check if rating was inserted
	if res.UpsertedCount == 0 {
		return &pb.RateResponse{Updated: false}, nil
	}
	return &pb.RateResponse{Updated: true}, nil
}
func (s *Server) RemoveRating(parent context.Context, dto *pb.RemoveRatingRequest) (*pb.RemoveRatingResponse, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// get object id from dto
	reservationId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}
	userId, err := primitive.ObjectIDFromHex(dto.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}

	res, err := s.rating_coll.DeleteOne(ctx, bson.M{"reservation_id": reservationId, "user_id": userId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot delete rating")
	}

	if res.DeletedCount == 0 {
		return &pb.RemoveRatingResponse{Removed: false}, nil
	}
	return &pb.RemoveRatingResponse{Removed: true}, nil
}
func (s *Server) GetMyRatings(parent context.Context, dto *pb.RatingIdRequest) (*pb.RatingList, error) {
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle()

	// get object id from dto
	userId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id")
	}

	// get ratings
	cursor, err := s.rating_coll.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get ratings")
	}

	// get ratings from cursor
	var ratings []*model.Rating
	if err = cursor.All(ctx, &ratings); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get ratings")
	}

	// convert ratings to pb
	var pbRatings []*pb.Rating
	for _, rating := range ratings {
		pbRatings = append(pbRatings, rating.ConvertToPbAccommodationRating())
	}

	return &pb.RatingList{Ratings: pbRatings}, nil

}
