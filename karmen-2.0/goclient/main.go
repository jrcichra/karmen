package main

import (
	"context"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	pb "github.com/jrcichra/karmen/goclient/grpc"
	"google.golang.org/grpc"
)

const name = "bob"

func register(client pb.KarmenClient) {
	response, err := client.Register(context.Background(),
		&pb.RegisterRequest{
			Name:      name,
			Timestamp: time.Now().Unix(),
			Events:    map[string]string{"event": "event"},
			Actions:   map[string]string{"action": "action"},
		})
	if err != nil {
		panic(err)
	}
	log.Println(response)
}

func handleActions(client pb.KarmenClient) {
	dispatcher, err := client.ActionDispatcher(context.Background())
	if err != nil {
		panic(err)
	}
	err = dispatcher.Send(&pb.ActionResponse{Hostname: name})
	if err != nil {
		panic(err)
	}
	for {
		msg, err := dispatcher.Recv()
		if err != nil {
			panic(err)
		}
		log.Println(msg.RequesterName, "requested I run", msg.Action.ActionName+". It's going to take me a few seconds...")
		log.Println("Parameters:")
		spew.Dump(msg.Action.Parameters)
		time.Sleep(5 * time.Second)
		log.Println("Finished running", msg.Action.ActionName, "for", msg.RequesterName)
		result := &pb.Result{Code: 200, Parameters: map[string]string{"asdf": "1234"}}
		dispatcher.Send(&pb.ActionResponse{Result: result})
	}
}

func sendEvent(client pb.KarmenClient) {
	event := &pb.Event{EventName: "event", Timestamp: time.Now().Unix(), Parameters: map[string]string{"justin": "rocks"}}
	response, err := client.EmitEvent(context.Background(), &pb.EventRequest{RequesterName: name, Event: event})
	if err != nil {
		panic(err)
	}
	log.Println(response)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewKarmenClient(conn)
	register(client)
	go handleActions(client)
	sendEvent(client)
	select {}
}
