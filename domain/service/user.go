package service

import (
	"go-app-service-test/domain/model"
	"go-app-service-test/domain/repository"
)

type UserDomainService struct {
	UserRepository repository.IUserRepository
}

func NewUserDomainService(r repository.IUserRepository) (UserDomainService, error) {
	return UserDomainService{UserRepository: r}, nil
}

func (us UserDomainService) Exists(user model.User) bool {
	_, err := us.UserRepository.FindByName(user.Name())
	if err != nil {
		return false
	}

	return true
}
