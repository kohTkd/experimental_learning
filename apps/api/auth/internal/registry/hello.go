package registry

import "github.com/kohTkd/experimental_learning/internal/adapter/handler"

func (r *Registry) registerHello() {
	handler.NewHelloServerHandler()
}
