package repository

import (
	"context"

	"github.com/kohTkd/experimental_learning/internal/domain/model"
	"github.com/kohTkd/experimental_learning/internal/domain/repository"
)

type (
	AccountsRepository struct{}
)

func (r *AccountsRepository) Create(ctx context.Context, c repository.Client, a *model.Account) (*model.Account, error) {
	ac, err := c.Client.Accounts.Create().SetEmail(a.Email).SetPassword(a.Password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return ac, nil
}
