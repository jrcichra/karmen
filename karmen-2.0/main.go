package main

import (
	"log"
	"net"

	pb "github.com/jrcichra/karmen/grpc"
	"google.golang.org/grpc"
)

const port = ":8080"

func serveGRPC() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKarmenServer(s, &karmen{})
	log.Println("Serving gRPC on port " + port + "...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// go serveGRPC()
	config := &Config{}
	config.LoadConfig("example.yml")
	config.dumpConfig()
	// Just to hold us until we have something that holds main - keep things in parallel that can be
	// for {
	// 	time.Sleep(1 * time.Second)
	// }
}
