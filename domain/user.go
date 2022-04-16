package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ScreenName  string
	DisplayName string
	Password    string
	Email       *string
}

type Users []User

type UserForCreate struct {
	ScreenName  string  `json:"screen_name" binding:"required"`
	DisplayName string  `json:"display_name" binding:"required"`
	Password    string  `json:"password_name" binding:"required"`
	Email       *string `json:"email_name" binding:"required"`
}

type UserForGet struct {
	ID          uint    `json:"id"`
	ScreenName  string  `json:"screen_name"`
	DisplayName string  `json:"display_name"`
	Email       *string `json:"email"`
}

type UserForUpdate struct {
	ID          uint    `json:"id"`
	ScreenName  string  `json:"screen_name"`
	DisplayName string  `json:"display_name"`
	Email       *string `json:"email"`
}

type UserListForGet []UserForGet

func (u *User) BuildUserForCreate() UserForCreate {
	user := UserForCreate{}
	user.ScreenName = u.ScreenName
	user.DisplayName = u.DisplayName
	if u.Email != nil {
		user.Email = u.Email
	} else {
		empty := ""
		user.Email = &empty
	}
	return user
}

func (u *User) BuildUserForGet() UserForGet {
	user := UserForGet{}
	user.ID = u.ID
	user.ScreenName = u.ScreenName
	user.DisplayName = u.DisplayName
	if u.Email != nil {
		user.Email = u.Email
	} else {
		empty := ""
		user.Email = &empty
	}
	return user
}

func (u Users) BuildUserListForGet() UserListForGet {
	var user_list UserListForGet
	for _, v := range u {
		user := UserForGet{}
		user.ID = v.ID
		user.ScreenName = v.ScreenName
		user.DisplayName = v.DisplayName
		if v.Email != nil {
			user.Email = v.Email
		} else {
			empty := ""
			user.Email = &empty
		}
		user_list = append(user_list, user)
	}
	return user_list
}
