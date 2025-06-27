package user

import "sample/app/model"

type IUserRepository interface {
	Register(user model.User) (model.User, error)
	Login(username string, password string) (model.User, error)
	GetUserByID(id int) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(id int) error
}