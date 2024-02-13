package handler

import (
	"context"
	"fmt"

	pb "github.com/kohTkd/experimental_learning/internal/adapter/grpc"
)

type helloServerHandler struct {
	pb.UnimplementedHelloWorldServer
}

func NewHelloServerHandler() *helloServerHandler {
	return &helloServerHandler{}
}

func (s *helloServerHandler) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s", req.GetName()),
	}, nil
}
