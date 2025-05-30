package AddUserMess

import (
	"webb/dao"
	"webb/model"
)

// 添加用户个人信息
func PersonalMessAdd(account string, u model.User) (err error) {
	return dao.AddMess(account, u)
}
