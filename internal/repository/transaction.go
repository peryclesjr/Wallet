package repository

import "github.com/preyclesjr/ms-wallet/internal/entity"

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
}
