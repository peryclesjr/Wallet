package createclient

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

func TestCreateClientUseCase_Execute(t *testing.T) {

	m := new(ClientRepositoryMock)
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	inputDTO := &CreateClientInputDTO{
		Name:  "John Doe",
		Email: "f1@f.com",
	}
	outputDTO, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, outputDTO)
	assert.NotEmpty(t, outputDTO.ID)
	assert.Equal(t, "John Doe", outputDTO.Name)
	assert.Equal(t, "f1@f.com", outputDTO.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}

func TestCreateClientUseCase_ExecuteWithInvalidInputName(t *testing.T) {

	m := new(ClientRepositoryMock)
	uc := NewCreateClientUseCase(m)

	inputDTO := &CreateClientInputDTO{
		Name:  "",
		Email: "",
	}
	outputDTO, err := uc.Execute(inputDTO)
	assert.NotNil(t, err)
	assert.Nil(t, outputDTO)
	assert.Equal(t, "name is required", err.Error())

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 0)
}

func TestCreateClientUseCase_ExecuteWithInvalidInputEmail(t *testing.T) {

	m := new(ClientRepositoryMock)
	uc := NewCreateClientUseCase(m)

	inputDTO := &CreateClientInputDTO{
		Name:  "John Doe",
		Email: "",
	}
	outputDTO, err := uc.Execute(inputDTO)
	assert.NotNil(t, err)
	assert.Nil(t, outputDTO)
	assert.Equal(t, "email is required", err.Error())

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 0)
}
