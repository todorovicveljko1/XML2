package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"reservation.accommodation.com/config"
	"reservation.accommodation.com/pb"
	"reservation.accommodation.com/src"
)

func main() {
	cfg := config.GetConfig()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	server, err := src.NewServer(&cfg)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	// Register service to expose
	pb.RegisterReservationServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
	server.Stop()
}
