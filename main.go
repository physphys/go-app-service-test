package main

import (
	"fmt"
	"go-app-service-test/application/usecase"
	"go-app-service-test/dicontainer"
	"go-app-service-test/domain/service"
	"go-app-service-test/handler"
	"go-app-service-test/inmemrepo"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("start")
	container := initDIContainer()

	handlerConf := handler.Config{DiContainer: container}
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handlerConf.GetUserByIDHandler).Methods("GET")

	const serverTimeout = 15 * time.Second
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: serverTimeout,
		ReadTimeout:  serverTimeout,
	}

	log.Fatal(srv.ListenAndServe())
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
