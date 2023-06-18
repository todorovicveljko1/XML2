package client

import (
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
)

func RatingGRPCClient(conn *grpc.ClientConn) pb.RatingServiceClient {
	return pb.NewRatingServiceClient(conn)
}
