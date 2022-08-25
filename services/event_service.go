package services

import (
	"context"
	"fmt"

	"github.com/thiagodevbrz/go-gRPC/pb"
)

type EventService struct {
	pb.UnimplementedEventServiceServer
}

func NewEventService() *EventService {
	return &EventService{}
}

//rpc handleUser (Event) returns (HandleEventResponse)
func (*EventService) HandleEvent(ctx context.Context, req *pb.Event) (*pb.HandleEventResponse, error) {

	fmt.Println(req.Title)

	return &pb.HandleEventResponse{
		Message: "gRPC has processed your request successfully - Thiago Pereira",
	}, nil
}
