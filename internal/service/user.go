package service

import (
	"context"
	"myapp/internal/model/entity"
)

type IUser interface {
	SignIn(ctx context.Context, username, password string) (user *entity.TUser, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
