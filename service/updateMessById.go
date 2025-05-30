package service

import "webb/dao"

// 修改账号密码
func UpdateMessByID(account string, password string, id int) (err error) {
	return dao.MessUpdateByID(account, password, id)
}
