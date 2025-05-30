package QueueUserMess

import (
	"webb/dao"
	"webb/model"
)

// 查询用户个人信息
func PersonalMessQueue(account string) (u model.User) {
	return dao.MessSelect(account)
}
