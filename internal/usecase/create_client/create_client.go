package createclient

import (
	"github.com/preyclesjr/ms-wallet/internal/entity"
	"github.com/preyclesjr/ms-wallet/internal/repository"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt string
	UpdatedAt string
}

type CreateClientUseCase struct {
	ClientRepository repository.ClientRepository
}

func NewCreateClientUseCase(lientRepository repository.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientRepository: lientRepository,
	}

}

func (uc *CreateClientUseCase) Execute(inputDTO *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(inputDTO.Name, inputDTO.Email)
	if err != nil {
		return nil, err
	}

	err = uc.ClientRepository.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt.String(),
		UpdatedAt: client.UpdatedAt.String(),
	}, nil

}
