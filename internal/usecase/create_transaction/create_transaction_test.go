package createtransaction

import (
	"testing"

	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionRepository struct {
	mock.Mock
}

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountRepositoryMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *TransactionRepository) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCAse_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "f1@g.com")
	account1 := entity.NewAccount(client1)
	account1.Deposit(1000)

	client2, _ := entity.NewClient("Petter", "f2@g.com")
	account2 := entity.NewAccount(client2)
	account2.Deposit(1000)

	mockAccountRepository := &AccountRepositoryMock{}
	mockAccountRepository.On("FindByID", account1.ID).Return(account1, nil)
	mockAccountRepository.On("FindByID", account2.ID).Return(account2, nil)

	mockTransactionRepository := &TransactionRepository{}
	mockTransactionRepository.On("Create", mock.Anything).Return(nil)

	inputdto := &CreateTransactionInputDTO{
		AccountIdFrom: account1.ID,
		AccountIdTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransactionRepository, mockAccountRepository)
	outputDTO, err := uc.Execute(inputdto)

	assert.Nil(t, err)
	assert.NotNil(t, outputDTO)
	assert.NotEmpty(t, outputDTO.TransactionId)
	mockTransactionRepository.AssertExpectations(t)
	mockTransactionRepository.AssertNumberOfCalls(t, "Create", 1)
	mockAccountRepository.AssertExpectations(t)
	mockAccountRepository.AssertNumberOfCalls(t, "FindByID", 2)

}
