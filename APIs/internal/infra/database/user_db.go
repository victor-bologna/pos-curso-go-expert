package database

import (
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.First(&user, "email = ?", email).Error
	return &user, err
}
