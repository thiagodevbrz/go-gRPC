package main

import (
	"context"
	"fmt"
	"log"

	"github.com/thiagodevbrz/go-gRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewEventServiceClient(connection)

	sendEvent(client)
}

func sendEvent(client pb.EventServiceClient) {
	req := &pb.Event{
		Title:   "My first gRPC call",
		Message: "Message for gRPC call",
	}

	res, err := client.HandleEvent(context.Background(), req)

	if err != nil {
		log.Fatal("Could call HandleEvent gRPC: %v", err)
	}

	fmt.Println(res)

}
