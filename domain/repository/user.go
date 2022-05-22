package repository

import "go-app-service-test/domain/model"

type IUserRepository interface {
	FindByID(id model.UserID) (model.User, error)
	FindByName(name model.UserName) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
}
