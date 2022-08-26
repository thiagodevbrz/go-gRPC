package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/thiagodevbrz/go-gRPC/pb"
)

type EventService struct {
	pb.UnimplementedEventServiceServer
}

func NewEventService() *EventService {
	return &EventService{}
}

func (*EventService) HandleEvent(ctx context.Context, req *pb.Event) (*pb.HandleEventResponse, error) {

	fmt.Println(req.Title)

	return &pb.HandleEventResponse{
		Status: "gRPC has processed your request successfully - Thiago Pereira",
	}, nil
}

func (*EventService) HandleDetailedEvent(req *pb.Event, stream pb.EventService_HandleDetailedEventServer) error {
	stream.Send(&pb.HandleEventResponse{
		Status: "Status received",
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.HandleEventResponse{
		Status: "Processing",
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.HandleEventResponse{
		Status: "Processed",
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*EventService) ClientStream(stream pb.EventService_ClientStreamServer) error {
	events := []*pb.Event{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Events{
				Event: events,
			})
		}
		if err != nil {
			log.Fatal("Error recieving stream : %v", err)
		}

		events = append(events, &pb.Event{
			Title: "Teste",
		})

		fmt.Println("Adding new Event", req.GetMessage())
	}
}

func (*EventService) BidirectionalEventStream(stream pb.EventService_BidirectionalEventStreamServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal("Error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.HandleEventResponse{
			Status: "Processed with success" + req.Message,
		})

		if err != nil {
			log.Fatal("Error sending stream to the client: %v", err)
		}
	}
}
