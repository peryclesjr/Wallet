package repository

import "github.com/preyclesjr/ms-wallet/internal/entity"

type ClientRepository interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
	Update(client *entity.Client) error
	Delete(id string) error
}
