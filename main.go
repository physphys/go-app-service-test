package main

import (
	"fmt"
	"go-app-service-test/application/usecase"
	"go-app-service-test/dicontainer"
	"go-app-service-test/domain/service"
	"go-app-service-test/inmemrepo"
	"log"
)

func main() {
	fmt.Println("start")

	container := initDIContainer()
	get, err := container.Get(dicontainer.DefNameUserAppService)
	if err != nil {
		log.Fatal(err)
	}
	userAppService, ok := get.(usecase.IUserAppService)
	if !ok {
		log.Fatal(err)
	}
	_, err = userAppService.Get("tekitou")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("end")
}

func initDIContainer() *dicontainer.Container {
	container := dicontainer.NewContainer()

	urDefinition := &dicontainer.Definition{
		Name: dicontainer.DefNameUserRepository,
		Builder: func(con *dicontainer.Container) (interface{}, error) {
			return inmemrepo.UserRepository{}, nil
		},
	}
	container.Register(urDefinition)

	udsDefinition := &dicontainer.Definition{
		Name: dicontainer.DefNameUserDomainService,
		Builder: func(con *dicontainer.Container) (interface{}, error) {
			repo, err := con.Get(dicontainer.DefNameUserRepository)
			if err != nil {
				return nil, err
			}

			castedRepo, ok := repo.(inmemrepo.UserRepository)
			if !ok {
				return nil, dicontainer.CastError
			}
			userDomainService, err := service.NewUserDomainService(&castedRepo)
			if err != nil {
				return nil, dicontainer.CastError
			}

			return userDomainService, nil
		},
	}
	container.Register(udsDefinition)

	uasDefinition := &dicontainer.Definition{
		Name: dicontainer.DefNameUserAppService,
		Builder: func(con *dicontainer.Container) (interface{}, error) {
			repo, err := con.Get(dicontainer.DefNameUserRepository)
			if err != nil {
				return nil, err
			}
			castedRepo, ok := repo.(inmemrepo.UserRepository)
			if !ok {
				return nil, dicontainer.CastError
			}

			domainService, err := con.Get(dicontainer.DefNameUserDomainService)
			if err != nil {
				return nil, err
			}
			castedDomainService, ok := domainService.(service.UserDomainService)
			if !ok {
				return nil, dicontainer.CastError
			}

			userAppService, err := usecase.NewUserAppService(&castedRepo, castedDomainService)
			if err != nil {
				return nil, err
			}

			return userAppService, nil
		},
	}
	container.Register(uasDefinition)

	return container
}
