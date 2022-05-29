package usecase

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-app-service-test/application/dto"
	"go-app-service-test/domain/model"
	"go-app-service-test/domain/service"
	"go-app-service-test/inmemrepo"
	"testing"

	_ "github.com/stretchr/testify"
)

func TestUserAppService_Get(t *testing.T) {
	type args struct {
		userID string
	}

	var users []model.User
	user1, err := model.NewUser("user_1")
	if err != nil {
		return
	}
	users = append(users, user1)

	var tests = []struct {
		name    string
		users   []model.User
		args    args
		want    dto.UserData
		wantErr bool
	}{
		{
			name:  "not found user",
			users: []model.User{},
			args: args{
				userID: string(user1.ID()),
			},
			want:    dto.UserData{},
			wantErr: true,
		},
		{
			name:  "found user",
			users: users,
			args: args{
				userID: string(user1.ID()),
			},
			want:    dto.NewUserData(user1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := inmemrepo.UserRepository{
				Users: tt.users,
			}

			userDomainService, err := service.NewUserDomainService(&repo)
			if err != nil {
				fmt.Println(err)

				return
			}
			u := UserAppService{
				userRepository:    &repo,
				userDomainService: userDomainService,
			}

			got, err := u.Get(tt.args.userID)
			assert.Equal(t, tt.want, got)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
