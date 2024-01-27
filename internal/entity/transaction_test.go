package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	client, errC1 := NewClient("John Doe", "f1@f.com")
	assert.Nil(t, errC1)
	accountFrom := NewAccount(client)
	accountFrom.Balance = 1000
	//(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {

	client2, errC2 := NewClient("John Doe", "f1@f.com")
	assert.Nil(t, errC2)
	accountTo := NewAccount(client2)
	accountTo.Balance = 1000

	//Create transaction
	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(900), transaction.AccountFrom.Balance)
	assert.Equal(t, float64(1100), transaction.AccountTo.Balance)
}

func TestCreateNewTransactionWithEmptyAccountFrom(t *testing.T) {
	client, errC1 := NewClient("John Doe", "f1@f.com")
	assert.Nil(t, errC1)
	accountFrom := NewAccount(client)
	accountFrom.Balance = 0

	client2, errC2 := NewClient("Peter Park", "spider@gmail.com")
	assert.Nil(t, errC2)
	accountTo := NewAccount(client2)
	accountTo.Balance = 1000

	//Create transaction
	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient funds", err.Error())
}

func TestCreateNewTransactionWithEmptyAccountTo(t *testing.T) {
	client, errC1 := NewClient("John Doe", "f1@f.com ")
	assert.Nil(t, errC1)
	accountFrom := NewAccount(client)
	accountFrom.Balance = 1000

	client2, errC2 := NewClient("Peter Park", "spider@gmail.com")
	assert.Nil(t, errC2)
	accountTo := NewAccount(client2)
	accountTo.Balance = 0

	//Create transaction
	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(900), transaction.AccountFrom.Balance)
	assert.Equal(t, float64(100), transaction.AccountTo.Balance)
}

func TestCreateNewTransactionWithNegativeAmount(t *testing.T) {
	client, errC1 := NewClient("John Doe", "f1.gmail.com")
	assert.Nil(t, errC1)
	accountFrom := NewAccount(client)
	accountFrom.Balance = 1000

	client2, errC2 := NewClient("Peter Park", "spider@#gmail.com")
	assert.Nil(t, errC2)
	accountTo := NewAccount(client2)
	accountTo.Balance = 1000

	//Create transaction
	transaction, err := NewTransaction(accountFrom, accountTo, -100)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}

func TestCreateNewTransactionWithInsufficientFunds(t *testing.T) {
	client, errC1 := NewClient("John Doe", "f1.gmail.com")
	assert.Nil(t, errC1)
	accountFrom := NewAccount(client)
	accountFrom.Balance = 1000

	client2, errC2 := NewClient("Peter Park", "spider@#gmail.com")
	assert.Nil(t, errC2)
	accountTo := NewAccount(client2)
	accountTo.Balance = 1000

	//Create transaction
	transaction, err := NewTransaction(accountFrom, accountTo, 2000)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient funds", err.Error())
}
