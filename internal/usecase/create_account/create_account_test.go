package createaccount

import (
	"testing"

	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientRepositoryMock struct {
	mock.Mock
}

func (m *ClientRepositoryMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientRepositoryMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ClientRepositoryMock) Update(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
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

func TestAccountCreateUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "f1.g.com")
	clientMock := &ClientRepositoryMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &AccountRepositoryMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)

	inputDTO := &CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
