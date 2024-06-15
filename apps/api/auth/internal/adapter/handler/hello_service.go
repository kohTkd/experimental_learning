package handler

import (
	"context"
	"fmt"

	pb "github.com/kohTkd/experimental_learning/internal/adapter/grpc"
)

type HelloServerHandler struct {
	pb.UnimplementedHelloWorldServer
}

func NewHelloServerHandler() *HelloServerHandler {
	return &HelloServerHandler{}
}

func (s *HelloServerHandler) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", req.GetName()),
	}, nil
}
