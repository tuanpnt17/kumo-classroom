package entities

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}

func (u *User) TableName() string {
	return "users"
}

type IUserRepository interface {
	CheckExistByEmail(ctx context.Context, email string) (bool, error)
}
