package main

import (
	server "github.com/kohTkd/experimental_learning/internal/infrastructure/server"
)


func main() {
	server.NewGrpcServer(8080).Listen()
}
