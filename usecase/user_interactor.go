package usecase

import (
	"go_sample_api/domain"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

type ResultStatus struct {
	StatusCode int
	Error      error
}

func NewResultStatus(code int, err error) *ResultStatus {
	r := ResultStatus{}
	r.StatusCode = code
	r.Error = err
	return &r
}

func (interactor *UserInteractor) CreateUser(u domain.UserForCreate) (user domain.User, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	user, err := interactor.User.CreateUser(db, u)
	if err != nil {
		return domain.User{}, NewResultStatus(404, err)
	}
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) GetUserList() (user domain.UserListForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	found_user_list, err := interactor.User.GetUserList(db)
	if err != nil {
		return domain.UserListForGet{}, NewResultStatus(404, err)
	}
	user = found_user_list.BuildUserListForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) GetUserByID(id int) (user domain.UserForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	found_user, err := interactor.User.GetUserByID(db, id)
	if err != nil {
		return domain.UserForGet{}, NewResultStatus(404, err)
	}
	user = found_user.BuildUserForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) UpdateUser(id int, u domain.UserForUpdate) (user domain.UserForUpdate, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	updated_user, err := interactor.User.UpdateUser(db, id, u)
	if err != nil {
		return domain.UserForUpdate{}, NewResultStatus(404, err)
	}
	user = domain.UserForUpdate(updated_user.BuildUserForGet())
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) DeleteUser(id int) (resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	err := interactor.User.DeleteUser(db, id)
	if err != nil {
		return NewResultStatus(404, err)
	}
	return NewResultStatus(200, nil)
}
