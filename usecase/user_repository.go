package usecase

import (
	"github.com/jinzhu/gorm"

	"go_sample_api/domain"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
}
