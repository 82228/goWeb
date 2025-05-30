package QueueUserMess

import (
	"webb/dao"
	"webb/model"
)

// 查询所有账号信息、
func QueryAllMess() []model.User {
	user, _ := dao.QueryAllAccount()
	return user
}

// 根据账号寻找role
func QueryRoleByCount(account string) (role int) {
	r := dao.QueryRoleByAccount(account)
	return r
}
