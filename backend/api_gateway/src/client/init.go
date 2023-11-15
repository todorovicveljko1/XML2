package client

import (
	"context"
	"log"
	"time"

	"api.accommodation.com/config"
	"api.accommodation.com/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connections struct {
	authConn *grpc.ClientConn
	accConn  *grpc.ClientConn
	resConn  *grpc.ClientConn
	retConn  *grpc.ClientConn
	notConn  *grpc.ClientConn

	NC *nats.Conn
}

func (c *Connections) Close() {
	c.authConn.Close()
	c.accConn.Close()
	c.resConn.Close()
	c.retConn.Close()
	c.notConn.Close()
	c.NC.Close()
}

type Clients struct {
	Connections *Connections

	AuthClient          pb.AuthClient
	AccommodationClient pb.AccommodationServiceClient
	ReservationClient   pb.ReservationServiceClient
	RatingClient        pb.RatingServiceClient
	NotificationClient  pb.NotificationServiceClient
}

func InitClients(cfg *config.Config) *Clients {
	log.Println("Connecting to servises...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authConn, err := grpc.DialContext(ctx, cfg.AuthAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to auth servise")
		panic(err)
	}

	accConn, err := grpc.DialContext(ctx, cfg.AccAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to accommodation servise")
		panic(err)
	}

	resConn, err := grpc.DialContext(ctx, cfg.ResAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to reservation servise")
		panic(err)
	}

	retConn, err := grpc.DialContext(ctx, cfg.RetAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to rating servise")
		panic(err)
	}

	nc, err := nats.Connect(cfg.NatsAddress)
	if err != nil {
		log.Println("Can't connect to NATS server")
		panic(err)
	}

	notConn, err := grpc.DialContext(ctx, cfg.NotAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Can't connect to notification servise")
		panic(err)
	}

	log.Println("Connected to servises")
	return &Clients{
		Connections: &Connections{
			authConn: authConn,
			accConn:  accConn,
			resConn:  resConn,
			retConn:  retConn,
			notConn:  notConn,
			NC:       nc,
		},
		AuthClient:          AuthGRPCClient(authConn),
		AccommodationClient: AccommodationGRPCClient(accConn),
		ReservationClient:   ReservationGRPCClient(resConn),
		RatingClient:        RatingGRPCClient(retConn),
		NotificationClient:  NotificationGRPCClient(notConn),
	}
}

func (c *Clients) Close() {
	c.Connections.Close()
}
