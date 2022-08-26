package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	SendDetailedEvent(client)
	SendStreamEvent(client)
	SendEventsBidirectionalStream(client)

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

func SendDetailedEvent(client pb.EventServiceClient) {
	req := &pb.Event{
		Title:   "My event for process",
		Message: "Please send me a stream",
	}

	responseStream, err := client.HandleDetailedEvent(context.Background(), req)

	if err != nil {
		log.Fatal("Cold not process the stream request %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Cold not receive stream %v", err)
		}

		fmt.Println("Status", stream.String())
	}
}

func SendStreamEvent(client pb.EventServiceClient) {

	reqs := []*pb.Event{
		&pb.Event{
			Message: "Event 1",
		},
		&pb.Event{
			Message: "Event 2",
		},
		&pb.Event{
			Message: "Event 3",
		},
		&pb.Event{
			Message: "Event 4",
		},
		&pb.Event{
			Message: "Event 5",
		},
		&pb.Event{
			Message: "Event 6",
		},
	}

	stream, err := client.ClientStream(context.Background())

	if err != nil {
		log.Fatal("Error sendind event stream: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Error closing response: %v", err)
	}

	fmt.Println(res)
}

func SendEventsBidirectionalStream(client pb.EventServiceClient) {
	stream, err := client.BidirectionalEventStream(context.Background())

	if err != nil {
		log.Fatal("Error creating request: %v", err)
	}

	reqs := []*pb.Event{
		&pb.Event{
			Message: "Event 1",
		},
		&pb.Event{
			Message: "Event 2",
		},
		&pb.Event{
			Message: "Event 3",
		},
		&pb.Event{
			Message: "Event 4",
		},
		&pb.Event{
			Message: "Event 5",
		},
		&pb.Event{
			Message: "Event 6",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending event: ", req)
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal("Error receiving data: %v", err)
				break
			}

			fmt.Printf("Recebendo evento %v", res)
		}
		close(wait)
	}()

	<-wait
}
