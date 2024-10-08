package registry

import (
	"github.com/kohTkd/experimental_learning/internal/adapter/handler"
	"github.com/kohTkd/experimental_learning/internal/infrastructure/repository"
)

type Registry struct {
	HelloHandler *handler.HelloServerHandler
}

func NewRegistry(store *repository.Store) *Registry {
	r := &Registry{}
	r.register(store)
	return r
}

func (r *Registry) register(store *repository.Store) {
	r.registerHello()
}
