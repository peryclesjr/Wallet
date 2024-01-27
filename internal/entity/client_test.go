package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, error := NewClient("John Doe", "f_f@gmail.com")
	assert.Nil(t, error)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "f_f@gmail.com", client.Email)

}

func TestCreateNewClientWithEmptyName(t *testing.T) {
	client, err := NewClient("", "f_f@gmail.com")
	assert.NotNil(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "name is required", err.Error())
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "f@f.com")
	err := client.Update("John Doe Update", "f1@f.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "f1@f.com", client.Email)
}

func TestUpdateClientWithEmptyName(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	err := client.Update("", "f@f.com")
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())
}

func TestUpdateClientWithEmptyEmail(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	err := client.Update("John Doe Update", "")
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountWithOtherClientAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "f1@f.com")
	client2, _ := NewClient("Peter Park", "spider@gmail.com")
	account2 := NewAccount(client2)
	err := client.AddAccount(account2)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(client.Accounts))
	assert.Equal(t, "account does not belong to other client", err.Error())
}
