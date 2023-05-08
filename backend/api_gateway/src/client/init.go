package client

import (
	"log"

	"api.accommodation.com/config"
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connections struct {
	authConn *grpc.ClientConn
	accConn  *grpc.ClientConn
}

func (c *Connections) Close() {
	c.authConn.Close()
	c.accConn.Close()
}

type Clients struct {
	connections *Connections

	AuthClient          pb.AuthClient
	AccommodationClient pb.AccommodationServiceClient
}

func InitClients(cfg *config.Config) *Clients {
	authConn, err := grpc.Dial(cfg.AuthAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to auth servise")
		panic(err)
	}

	accConn, err := grpc.Dial(cfg.AccAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to accommodation servise")
		panic(err)
	}

	return &Clients{
		connections: &Connections{
			authConn: authConn,
			accConn:  accConn,
		},
		AuthClient:          AuthGRPCClient(authConn),
		AccommodationClient: AccommodationGRPCClient(accConn),
	}
}

func (c *Clients) Close() {
	c.connections.Close()
}
