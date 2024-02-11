package main

import (
	"context"
	"log"
	"net"

	pb "go_grpc/pb"
	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", "server:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
