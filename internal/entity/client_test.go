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
