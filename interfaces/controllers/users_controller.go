package controllers

import (
	"strconv"

	"go_sample_api/interfaces/database"
	"go_sample_api/usecase"

	"go_sample_api/domain"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (controller *UsersController) CreateUser(c Context) {
	user_payload := domain.UserForCreate{}
	err := c.Bind(&user_payload)
	if err != nil {
		c.JSON(400, NewH(err.Error(), nil))
		return
	}
	user, res := controller.Interactor.CreateUser(user_payload)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}

func (controller *UsersController) GetUserList(c Context) {
	user_list, res := controller.Interactor.GetUserList()
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user_list))
}

func (controller *UsersController) GetUserByID(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.GetUserByID(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}

func (controller *UsersController) UpdateUser(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user_payload := domain.UserForUpdate{}
	err := c.Bind(&user_payload)
	if err != nil {
		c.JSON(400, NewH(err.Error(), nil))
		return
	}

	user, res := controller.Interactor.UpdateUser(id, user_payload)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}

func (controller *UsersController) DeleteUser(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := controller.Interactor.DeleteUser(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", nil))
}
