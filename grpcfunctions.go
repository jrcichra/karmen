package main

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	pb "github.com/jrcichra/karmen/grpc"
)

type karmen struct {
	pb.UnimplementedKarmenServer
	Config *Config
	State  State
}

//Register - register a container
func (k *karmen) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received Register for: %v", in.GetName())
	k.State.Hosts[HostName(in.GetName())] = &Host{Online: true, Events: in.Events, Actions: in.Actions}
	debugPrintln("Dumping the state:")
	debugSpew(k.State.Hosts)
	return &pb.RegisterResponse{Request: in, Result: &pb.Result{Code: 200}}, nil
}

func (k *karmen) EmitEvent(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	log.Printf("Received EmitEvent for : %v", in.Event.GetEventName())

	// Get the event as parsed from the yaml
	event := k.Config.Events[in.RequesterName+"."+in.Event.GetEventName()]

	if event == nil {
		err := errors.New("Event '" + in.GetEvent().EventName + "' was not found in the YAML. Cannot launch it")
		log.Println(err)
		return &pb.EventResponse{Request: in, Result: &pb.Result{Code: 500}}, err
	}

	// generate a UUID for this event
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
	}
	k.State.Events = make(map[UUID]Results)

	// Run through the blocks
	for _, block := range event.Blocks {
		k.runBlock(block, in.RequesterName, uuid)
	}

	// When the event is done, delete the result history
	delete(k.State.Events, UUID(uuid.String()))

	return &pb.EventResponse{Request: in, Result: &pb.Result{Code: 200}}, nil
}

func (k *karmen) ActionDispatcher(s pb.Karmen_ActionDispatcherServer) error {
	// log.Println("Holding onto an ActionDispatcher. Waiting for a stub ActionResponse telling us who this is...")
	who, err := s.Recv()
	if err != nil {
		log.Println(err)
	}

	// hostname is the name of the host. We'll use this hostname to map a hostname with this dispatcher
	hostname := who.Hostname

	// Assign the dispatcher for this host
	log.Println("ActionDispatcher is open for", who.Hostname)
	k.State.Hosts[HostName(hostname)].Dispatcher = s

	// Keep the dispatcher alive so we can send actions later
	select {}

	return nil
}

func (k *karmen) PingPong(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	log.Println("Got a ping! Message:", in.Message)
	return &pb.Pong{Message: "Hello there, " + in.Message}, nil
}
