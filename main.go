package main

import (
	"log"
	"net"
	"os"

	pb "github.com/jrcichra/karmen/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"
const debug = false

func serveGRPC(c *Config) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKarmenServer(s, &karmen{Config: c, State: State{Hosts: make(map[HostName]*Host)}})
	reflection.Register(s)
	log.Println("Serving gRPC on port " + port + "...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadConfig(filename string) *Config {
	config := &Config{}
	config.LoadConfig(filename)
	if debug {
		config.dumpConfig()
	}
	return config
}

func main() {
	// get the first argument as the config file
	if len(os.Args) > 1 {
		configFile := os.Args[1]
		serveGRPC(loadConfig(configFile))
	} else {
		log.Println("Usage: karmen <config file>")
	}
}
