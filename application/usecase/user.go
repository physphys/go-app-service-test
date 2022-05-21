package usecase

import (
	"fmt"
	"go-app-service-test/application/dto"
	"go-app-service-test/domain/model"
	"go-app-service-test/domain/repository"
	"go-app-service-test/domain/service"
)

type UserAppService struct {
	userRepository    repository.IUserRepository
	userDomainService service.UserDomainService
}

func (u UserAppService) NewUser(
	userRepository repository.IUserRepository,
	userDomainService service.UserDomainService,
) (UserAppService, error) {
	return UserAppService{
		userRepository:    userRepository,
		userDomainService: userDomainService,
	}, nil
}

func (u UserAppService) Register(name string) error {
	newUser, err := model.NewUser(name)
	if err != nil {
		return err
	}

	if u.userDomainService.Exists(newUser) {
		return fmt.Errorf("user already exists for user_name: %s", newUser.Name())
	}

	if _, err := u.userRepository.Create(newUser); err != nil {
		return err
	}

	return nil
}

func (u UserAppService) Get(userID string) (dto.UserData, error) {
	domainUser, err := u.userRepository.FindByID(model.UserID(userID))
	if err != nil {
		return dto.UserData{}, fmt.Errorf("not found domainUser for id: %s", userID)
	}

	return dto.NewUserData(domainUser), nil
}

func (u UserAppService) Update(userID string, name string) error {
	domainUser, err := u.userRepository.FindByID(model.UserID(userID))
	if err != nil {
		return fmt.Errorf("not found domainUser for id: %s", userID)
	}

	if err := (&domainUser).ChangeName(name); err != nil {
		return err
	}
	if !u.userDomainService.Exists(domainUser) {
		return fmt.Errorf("user already exists")
	}

	if _, err := u.userRepository.Update(domainUser); err != nil {
		return err
	}

	return nil
}
