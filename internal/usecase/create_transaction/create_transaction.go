package createtransaction

import (
	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/preyclesjr/ms-wallet/internal/repository"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	TransactionId string
}

type CreateTransactionUseCase struct {
	transactionRepository repository.TransactionRepository
	accountRepository     repository.AccountRepository
}

func NewCreateTransactionUseCase(transactionRepository repository.TransactionRepository, accountRepository repository.AccountRepository) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

func (uc *CreateTransactionUseCase) Execute(inputDTO *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accoutFrom, err := uc.accountRepository.FindByID(inputDTO.AccountIdFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := uc.accountRepository.FindByID(inputDTO.AccountIdTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accoutFrom, accountTo, inputDTO.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		TransactionId: transaction.ID,
	}, nil
}
