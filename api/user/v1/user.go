package v1

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta   `path:"/user-login" tags:"User" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	UserID int64
}
