package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Joao", "joao@gmail.com", "test")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Joao", user.Name)
	assert.Equal(t, "joao@gmail.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("Joao", "joao@gmail.com", "test")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("test"))
	assert.False(t, user.ValidatePassword("false"))
}
