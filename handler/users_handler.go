package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-app-service-test/application/usecase"
	"go-app-service-test/dicontainer"
	"net/http"
)

type Config struct {
	DiContainer *dicontainer.Container
}

func (conf Config) GetUserByIDHandler(_ http.ResponseWriter, r *http.Request) {
	get, err := conf.DiContainer.Get(dicontainer.DefNameUserAppService)
	if err != nil {
		fmt.Println(err)

		return
	}
	userAppService, ok := get.(usecase.IUserAppService)
	if !ok {
		fmt.Println("DI container error")

		return
	}

	userID := mux.Vars(r)["id"]
	userData, err := userAppService.Get(userID)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(userData)
}
