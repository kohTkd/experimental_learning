package main

import (
	"log"
	"os"

	"github.com/kohTkd/experimental_learning/internal/config"
	"github.com/kohTkd/experimental_learning/internal/infrastructure/repository"
	"github.com/kohTkd/experimental_learning/internal/infrastructure/server"
	"github.com/kohTkd/experimental_learning/internal/registry"
	"github.com/kohTkd/experimental_learning/internal/util/environment"
)

func init() {
	envName := os.Getenv("ENV_NAME")
	if envName == "" {
		envName = "development"
	}
	var env *environment.EnvName
	if env = environment.Get(envName); env == nil {
		log.Panic("invalid environment name: " + envName)
	}
	config.Read(config.ReadConfigOption{Env: *env})
}

func main() {
	c, _ := repository.NewClient()
	r := newRegistry(c) // Pass the dereferenced value of c
	s := newServer(r)

	s.Listen()
}

func newRegistry(c *store.Client) *registry.Registry {
	return registry.NewRegistry(c)
}

func newServer(r *registry.Registry) *server.GrpcServer {
	s := server.NewGrpcServer(config.Conf.Server.Port)
	s.RegisterServices(*r)
	return s
}
