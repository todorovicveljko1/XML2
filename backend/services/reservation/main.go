package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"reservation.accommodation.com/config"
	"reservation.accommodation.com/pb"
	"reservation.accommodation.com/src"
)

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Get execution time
	start := time.Now()
	resp, err := handler(ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Printf("ERROR %s - %dms %s | %v", info.FullMethod, time.Since(start).Milliseconds(), s.Code().String(), s.Message())
		} else {

			log.Printf("ERROR %s - %dms UNKNOWN | %v", info.FullMethod, time.Since(start).Milliseconds(), err.Error())
		}
	} else {
		log.Printf("SUCCESS %s - %dms", info.FullMethod, time.Since(start).Milliseconds())
	}
	return resp, err
}

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

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	reflection.Register(grpcServer)
	// Register service to expose
	pb.RegisterReservationServiceServer(grpcServer, server)

	go func() {
		log.Printf("gRPC reservation server listening on %s", cfg.Address)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	log.Println("Shutting down reservation server...")
	grpcServer.Stop()
	server.Stop()
	log.Println("Reservation server stopped.")
}
