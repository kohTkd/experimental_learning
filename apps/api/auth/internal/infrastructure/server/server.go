package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	igrpc "github.com/kohTkd/experimental_learning/internal/infrastructure/grpc"
)

type grpcServer struct {
	port int
}

func NewGrpcServer(port int) *grpcServer {
	return &grpcServer{
		port: port,
	}
}

func (g *grpcServer) Listen() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	igrpc.RegisterServices(s)

	go func() {
		log.Printf("Start gRPC server port: %d", g.port)
		s.Serve(l)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Stopping gRPC server...")
	s.GracefulStop()
}
