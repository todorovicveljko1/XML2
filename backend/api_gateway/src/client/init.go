package client

import (
	"api.accommodation.com/config"
	"api.accommodation.com/pb"
	"google.golang.org/grpc"
)

type Connections struct {
	authConn *grpc.ClientConn
}

func (c *Connections) Close() {
	c.authConn.Close()
}

type Clients struct {
	connections *Connections

	AuthClient pb.AuthClient
}

func InitClients(cfg *config.Config) *Clients {
	authConn, err := grpc.Dial(cfg.AuthAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &Clients{
		connections: &Connections{
			authConn: authConn,
		},
		AuthClient: AuthGRPCClient(authConn),
	}
}

func (c *Clients) Close() {
	c.connections.Close()
}
