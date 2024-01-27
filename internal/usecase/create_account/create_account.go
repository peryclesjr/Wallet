package createaccount

import (
	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/preyclesjr/ms-wallet/internal/repository"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountRepository repository.AccountRepository
	ClientRepository  repository.ClientRepository
}

func NewCreateAccountUseCase(accountRepository repository.AccountRepository, clientRepository repository.ClientRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository: accountRepository,
		ClientRepository:  clientRepository,
	}
}

func (uc *CreateAccountUseCase) Execute(inputDTO *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientRepository.Get(inputDTO.ClientID)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	err = uc.AccountRepository.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
