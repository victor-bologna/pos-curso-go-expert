package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openUserDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	if err = db.AutoMigrate(&entity.User{}); err != nil {
		t.Error(err)
	}
	return db
}

func TestCreateUser(t *testing.T) {
	db := openUserDB(t)
	user, err := entity.NewUser("Joao", "joao@gmail.com", "test")
	assert.NoError(t, err)
	userDB := NewUserDB(db)

	err = userDB.CreateUser(user)
	assert.NoError(t, err)

	var userResult entity.User
	err = db.Where("id = ?", user.ID).First(&userResult).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userResult.ID)
	assert.Equal(t, user.Email, userResult.Email)
	assert.Equal(t, user.Name, userResult.Name)
	assert.NotNil(t, userResult.Password)
}

func TestFindByEmail(t *testing.T) {
	db := openUserDB(t)
	user, err := entity.NewUser("Joao", "joao@gmail.com", "test")
	assert.NoError(t, err)
	userDB := NewUserDB(db)

	err = userDB.CreateUser(user)
	assert.NoError(t, err)

	userResult, err := userDB.FindByEmail("joao@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userResult.ID)
	assert.Equal(t, user.Email, userResult.Email)
	assert.Equal(t, user.Name, userResult.Name)
	assert.NotNil(t, userResult.Password)
}
