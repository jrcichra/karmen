package main

import (
	"context"
	"log"

	pb "github.com/jrcichra/karmen/grpc"
)

type karmen struct {
	pb.UnimplementedKarmenServer
}

//Register - register a container
func (k *karmen) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received Register for: %v", in.GetName())
	// you would register the client here and send them a response
	return &pb.RegisterResponse{Request: in, Result: &pb.Result{Code: 200}}, nil
}

func (k *karmen) EmitEvent(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	log.Printf("Received EmitEvent for : %v", in.Event.GetEventName())
	// you would start a goroutine to handle this event here
	return &pb.EventResponse{Request: in, Result: &pb.Result{Code: 200}}, nil
}

func (k *karmen) ActionDispatcher(s pb.Karmen_ActionDispatcherServer) error {
	log.Println("Holding onto an ActionDispatcher")
	// you would dispatch actions and expect responses to those actions in here
	return nil
}
