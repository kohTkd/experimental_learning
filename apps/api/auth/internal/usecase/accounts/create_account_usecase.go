package accounts

import "github.com/kohTkd/experimental_learning/internal/domain/repository"

type CreateAccountUsecase struct {
	store             repository.Store
	accountRepository repository.IAccountsRepository
}

func NewCreateAccountUsecase(store repository.Store, accountRepository repository.IAccountsRepository) *CreateAccountUsecase {
	return &CreateAccountUsecase{
		store:             store,
		accountRepository: accountRepository,
	}
}
