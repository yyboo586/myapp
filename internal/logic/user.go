package logic

import (
	"context"
	"errors"
	"myapp/internal/dao"
	"myapp/internal/model/entity"
	"myapp/internal/service"
	"strings"
)

type UserLogic struct{}

func init() {
	service.RegisterUser(NewUserLogic())
}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (u *UserLogic) SignIn(ctx context.Context, username, password string) (userInfo *entity.TUser, err error) {
	userInfo, err = dao.TUser.GetUserByName(username)
	if err != nil {
		return nil, err
	}

	if strings.Compare(userInfo.Password, password) != 0 {
		return nil, errors.New("账户名或密码错误")
	}
	// if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password)); err != nil {
	// 	return nil, errors.New("账户名或密码错误")
	// }

	return
}
