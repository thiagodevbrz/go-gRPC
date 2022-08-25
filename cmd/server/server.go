package main

import (
	"log"
	"net"

	"github.com/thiagodevbrz/go-gRPC/pb"
	"github.com/thiagodevbrz/go-gRPC/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterEventServiceServer(grpcServer, services.NewEventService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not serve gRPC %v", err)
	}
}
