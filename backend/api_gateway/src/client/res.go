package client

import (
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
)

func ReservationGRPCClient(conn *grpc.ClientConn) pb.ReservationServiceClient {
	return pb.NewReservationServiceClient(conn)
}
