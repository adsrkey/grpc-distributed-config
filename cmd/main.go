package main

import (
	"github.com/go-grpc-course/config/configpb"
	"github.com/go-grpc-course/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	configpb.RegisterConfigServiceServer(s, &server.Server{})
	log.Println("start distributed-config-grpc-service")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
