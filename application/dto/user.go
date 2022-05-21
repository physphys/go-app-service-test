package dto

import "go-app-service-test/domain/model"

type UserData struct {
	id   string
	name string
}

func NewUserData(domainUser model.User) UserData {
	return UserData{
		id:   string(domainUser.ID()),
		name: string(domainUser.Name()),
	}
}
