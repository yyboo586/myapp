package user

import (
	"context"
	"errors"
	"strings"

	v1 "myapp/api/user/v1"
	"myapp/internal/service"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}

	userInfo, err := service.User().SignIn(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &v1.UserLoginRes{
		UserID: userInfo.Id,
	}, nil
}
