package karmen

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	pb "github.com/jrcichra/karmen/goclient/grpc"
	"google.golang.org/grpc"
)

type KarmenClient struct {
	Name    string
	pb      pb.KarmenClient
	actions map[string]func(parameters map[string]string) *Result
}

type Result pb.Result

func (k *KarmenClient) Init(name string) error {
	k.Name = name
	k.actions = make(map[string]func(parameters map[string]string) *Result)
	return nil
}

func (k *KarmenClient) connect(hostname string, port int) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(hostname+":"+strconv.Itoa(port), opts...)
	if err != nil {
		return err
	}
	k.pb = pb.NewKarmenClient(conn)
	return nil
}

func (k *KarmenClient) Register(hostname string, port int) error {
	k.connect(hostname, port)
	response, err := k.pb.Register(context.Background(),
		&pb.RegisterRequest{
			Name:      k.Name,
			Timestamp: time.Now().Unix(),
		})
	if err != nil {
		panic(err)
	}
	//handle incoming actions
	go k.handleActions()

	if response.Result.Code != 200 {
		return errors.New("[Karmen] - register got a non-OK return code: " + strconv.FormatInt(response.Result.Code, 10))
	}
	return nil
}

func (k *KarmenClient) handleActions() {
	dispatcher, err := k.pb.ActionDispatcher(context.Background())
	if err != nil {
		panic(err)
	}
	err = dispatcher.Send(&pb.ActionResponse{Hostname: k.Name})
	if err != nil {
		panic(err)
	}

	for {
		// wait for action requests
		msg, err := dispatcher.Recv()
		if err != nil {
			panic(err)
		}
		// run the action
		go func() {
			result := k.actions[msg.Action.GetActionName()](msg.Action.GetParameters())

			// send the action result
			log.Println("Finished running", msg.Action.ActionName, "for", msg.RequesterName)

			// result := &pb.Result{Code: 200, Parameters: map[string]string{"asdf": "1234"}}

			dispatcher.Send(&pb.ActionResponse{Result: (*pb.Result)(result)})
		}()
	}
}

func (k *KarmenClient) RunEvent(name string, params map[string]string) (*pb.EventResponse, error) {
	event := &pb.Event{EventName: name, Timestamp: time.Now().Unix()}
	return k.pb.EmitEvent(context.Background(), &pb.EventRequest{RequesterName: k.Name, Event: event, Parameters: params})
}

func (k *KarmenClient) AddAction(function func(parameters map[string]string) *Result, name string) {
	k.actions[name] = function
}
