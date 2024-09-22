package main

import (
	"log"
	"net"
	"os"

	pb "ch7proto/todo/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
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

	defer func(list net.Listener) {
		if err := list.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(list)

	log.Printf("Listening at %s\n", addr)

	creds, err := credentials.NewServerTLSFromFile("./../certs/server_cert.pem", "./../certs/server_key.pem")
	if err != nil {
		log.Fatalf("failed to create credential: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(unaryAtuhToken,unaryLog),
		grpc.ChainStreamInterceptor(streamAuthToken,streamLog),
	}

	s := grpc.NewServer(opts...)

	pb.RegisterTodoServiceServer(s, &server{
		d: New(),
	})

	defer s.Stop()

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
