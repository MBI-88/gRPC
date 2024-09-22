package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("usage: server [IP_ADDR]")
	}

	addr := args[0]
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer func (list net.Listener)  {
		if err := list.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(list)

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer([]grpc.ServerOption{}...)

	defer s.Stop()

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}