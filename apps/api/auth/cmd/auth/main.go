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
	c, _ := repository.NewStore()
	r := newRegistry(c)
	s := newServer()

	s.RegisterServices(r)

	s.Listen()
}

func newRegistry(c *repository.Store) *registry.Registry {
	return registry.NewRegistry(c)
}

func newServer() *server.GrpcServer {
	return server.NewGrpcServer(config.Conf.Server.Port)
}
