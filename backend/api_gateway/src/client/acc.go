package client

import (
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
)

func AccommodationGRPCClient(conn *grpc.ClientConn) pb.AccommodationServiceClient {
	return pb.NewAccommodationServiceClient(conn)
}
