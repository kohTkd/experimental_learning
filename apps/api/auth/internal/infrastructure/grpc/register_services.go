package grpc

import (
	pb "github.com/kohTkd/experimental_learning/internal/adapter/grpc"
	handler "github.com/kohTkd/experimental_learning/internal/adapter/handler"
	"google.golang.org/grpc"
)

func RegisterServices(s *grpc.Server) {
	pb.RegisterHelloWorldServer(s, handler.NewHelloServerHandler())
}
