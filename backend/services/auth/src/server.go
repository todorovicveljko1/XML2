package src

import (
	"context"
	"log"
	"time"

	"auth.accommodation.com/config"
	pb "auth.accommodation.com/pb"
	"auth.accommodation.com/src/db"
	"auth.accommodation.com/src/helper"
	"auth.accommodation.com/src/model"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServer
	cfg *config.Config

	user_collection *mongo.Collection

	dbClient *mongo.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	client, _ := db.DbInit(cfg)

	user_collection := client.Database("accommodation_auth").Collection("users")

	return &Server{cfg: cfg, dbClient: client, user_collection: user_collection}, nil
}

func (s *Server) Stop() {
	if err := s.dbClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Server) Login(parent context.Context, dto *pb.LoginRequest) (*pb.LoginResponse, error) {

	var user model.User
	// cancle context after 5 seconds
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()
	//Get user by username
	err := s.user_collection.FindOne(ctx, bson.M{"username": dto.Username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.Unauthenticated, "Bad Credentials")
		}
		return nil, status.Error(codes.Internal, "Error while fetching user")
	}
	//Check if passwords are same
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Bad Credentials")
	}
	//generate token
	token, err := helper.GenerateJWT(&user, s.cfg.Secret)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while generating token")
	}
	return &pb.LoginResponse{Token: *token}, nil
}

func (s *Server) Register(parent context.Context, dto *pb.RegisterRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while hashing password")
	}


	user := model.User{
		Id:            primitive.NewObjectID(),
		Username:      dto.Username,
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Email:         dto.Email,
		Password:      string(hashedPassword),
		PlaceOfLiving: dto.PlaceOfLiving,
		Role:          model.Role(dto.Role),
	}

	_, err = s.user_collection.InsertOne(ctx, user)
	if err != nil {
		if db.IsDup(err) {
			return nil, status.Error(codes.AlreadyExists, "User already exists")
		}
		log.Println(err)
		return nil, status.Error(codes.Internal, "Error while inserting user")
	}
	return user.ConvertToPbUser(), nil
}

// TODO: get token from metadata
func (s *Server) AuthUser(parent context.Context, dto *pb.AuthUserRequest) (*pb.User, error) {
	var user model.User

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	//Decode token and check if is valid
	token, err := helper.DecodeJWT(dto.Token, s.cfg.Secret)
	if err != nil || !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}
	// teke user id from token
	userId, err := primitive.ObjectIDFromHex(token.Claims.(jwt.MapClaims)["user"].(string))
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching user")
	}
	// find that user
	err = s.user_collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, "Error while fetching user")
	}

	return user.ConvertToPbUser(), nil
}

func (s *Server) GetUser(parent context.Context, dto *pb.GetUserRequest) (*pb.User, error) {
	var user model.User

	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	userId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error while fetching user")
	}

	err = s.user_collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, "Error while fetching user")
	}

	return user.ConvertToPbUser(), nil
}
