package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   string
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {

	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now().String(),
	}

	err := transaction.ValidateTransaction()
	if err != nil {
		return nil, err
	}

	errTransatction := transaction.ProcessTransaction()
	if errTransatction != nil {
		return nil, errTransatction
	}

	return transaction, nil
}

func (t *Transaction) ValidateTransaction() error {

	if t.AccountFrom == nil {
		return errors.New("accountFrom is required")
	}

	if t.AccountTo == nil {
		return errors.New("accountTo is required")
	}

	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}

	return nil
}

func (t *Transaction) ProcessTransaction() error {

	errWithdraw := t.AccountFrom.Withdraw(t.Amount)
	if errWithdraw != nil {
		return errWithdraw
	}

	errDeposit := t.AccountTo.Deposit(t.Amount)
	if errDeposit != nil {
		return errDeposit
	}

	return nil
}
