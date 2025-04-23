// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure of table t_user for DAO operations like Where/Data.
type TUser struct {
	g.Meta      `orm:"table:t_user, do:true"`
	Id          interface{} //
	Name        interface{} //
	Password    interface{} //
	Email       interface{} //
	LastLoginAt *gtime.Time //
	Sex         interface{} //
	Age         interface{} //
}
