package client

import (
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
)

func AuthGRPCClient(conn *grpc.ClientConn) pb.AuthClient {
	return pb.NewAuthClient(conn)
}
