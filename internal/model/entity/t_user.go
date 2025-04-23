// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure for table t_user.
type TUser struct {
	Id          int64       `json:"id"          orm:"id"            description:""` //
	Name        string      `json:"name"        orm:"name"          description:""` //
	Password    string      `json:"password"    orm:"password"      description:""` //
	Email       string      `json:"email"       orm:"email"         description:""` //
	LastLoginAt *gtime.Time `json:"lastLoginAt" orm:"last_login_at" description:""` //
	Sex         int         `json:"sex"         orm:"sex"           description:""` //
	Age         int         `json:"age"         orm:"age"           description:""` //
}
