package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/kohTkd/experimental_learning/internal/adapter/grpc"
	"github.com/kohTkd/experimental_learning/internal/registry"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	port   int
	server *grpc.Server
}

func NewGrpcServer(port int) *GrpcServer {
	return &GrpcServer{
		port:   port,
		server: grpc.NewServer(),
	}
}

func (g *GrpcServer) RegisterServices(r registry.Registry) {
	pb.RegisterHelloWorldServer(g.server, r.HelloHandler)
}

func (g *GrpcServer) Listen() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		panic(err)
	}

	go func() {
		log.Printf("Start gRPC server port: %d", g.port)
		g.server.Serve(l)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Stopping gRPC server...")
	g.server.GracefulStop()
}
