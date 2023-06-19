package client

import (
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
)

func NotificationGRPCClient(conn *grpc.ClientConn) pb.NotificationServiceClient {
	return pb.NewNotificationServiceClient(conn)
}
