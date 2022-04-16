package usecase

import (
	"gorm.io/gorm"

	"go_sample_api/domain"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, u domain.UserForCreate) (user domain.User, err error)
	GetUserList(db *gorm.DB) (user domain.Users, err error)
	GetUserByID(db *gorm.DB, id int) (user domain.User, err error)
	UpdateUser(db *gorm.DB, id int, u domain.UserForUpdate) (user domain.User, err error)
	DeleteUser(db *gorm.DB, id int) (err error)
}
