package database

import (
	"errors"

	"gorm.io/gorm"

	"go_sample_api/domain"
)

type UserRepository struct{}

func Authenticate() {

}

func (repo *UserRepository) CreateUser(db *gorm.DB, u domain.UserForCreate) (user domain.User, err error) {
	user.Email = u.Email
	user.ScreenName = u.ScreenName
	user.DisplayName = u.DisplayName
	user.Password = u.Password
	result := db.Create(&user)
	if result.Error != nil {
		return domain.User{}, errors.New("fail on user creation")
	}
	return user, nil
}

func (repo *UserRepository) GetUserList(db *gorm.DB) (user domain.Users, err error) {
	users := domain.Users{}
	db.Find(&users)
	if users == nil {
		return domain.Users{}, errors.New("user is not found")
	}
	return users, nil
}

func (repo *UserRepository) GetUserByID(db *gorm.DB, id int) (user domain.User, err error) {
	user = domain.User{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) UpdateUser(db *gorm.DB, id int, u domain.UserForUpdate) (user domain.User, err error) {
	user = domain.User{}
	user.ID = uint(id)
	user.ScreenName = u.ScreenName
	user.DisplayName = u.DisplayName
	user.Email = u.Email
	result := db.Model(&user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) DeleteUser(db *gorm.DB, id int) (err error) {
	user := domain.User{}
	result := db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
