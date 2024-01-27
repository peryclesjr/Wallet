package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, float64(0), account.Balance)
}

func TestCreateNewAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestDeposit(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	account := NewAccount(client)
	err := account.Deposit(100)
	assert.Nil(t, err)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDepositWithNegativeAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@.com")
	account := NewAccount(client)
	err := account.Deposit(-100)
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), account.Balance)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}

func TestWithdraw(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	account := NewAccount(client)
	account.Deposit(100)
	err := account.Withdraw(50)
	assert.Nil(t, err)
	assert.Equal(t, float64(50), account.Balance)
}

func TestWithdrawWithNegativeAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	account := NewAccount(client)
	account.Deposit(100)
	err := account.Withdraw(-50)
	assert.NotNil(t, err)
	assert.Equal(t, float64(100), account.Balance)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}
