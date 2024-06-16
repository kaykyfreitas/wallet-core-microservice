package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john@email.com", client.Email)
	assert.NotNil(t, client.CreatedAt)
	assert.NotNil(t, client.UpdatedAt)
}

func TestCreateNewClientWithInvalidArgs(t *testing.T) {
	client, err := NewClient("", "john@email.com")
	assert.Error(t, err, "name is required")
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@email.com")
	err := client.Update("John Doe Update", "john@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "john@email.com", client.Email)
	assert.NotNil(t, client.CreatedAt)
	assert.NotNil(t, client.UpdatedAt)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "john@email.com")
	err := client.Update("", "john@email.com")
	assert.Error(t, err, "name is required")
}
