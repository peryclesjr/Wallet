package repository

import "github.com/preyclesjr/ms-wallet/internal/entity"

type AccountRepository interface {
	FindByID(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
