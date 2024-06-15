package usecase

import (
	pb "github.com/kohTkd/experimental_learning/internal/adapter/grpc"
)

type HelloInputPort interface {
	Hello(name string)
}

type HelloOutputPort interface {
	Write(response pb.HelloResponse,message string)
}
