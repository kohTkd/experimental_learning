package repository

import (
	"context"

	"github.com/kohTkd/experimental_learning/ent"
	"github.com/kohTkd/experimental_learning/internal/domain/repository"
)

func client(ctx context.Context) *repository.Client {
	return repository.NewClient(ent.FromContext(ctx))
}
