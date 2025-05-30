package service

import (
	"fmt"
	"webb/dao"
)

// 查询用户存不存在
func CheckAccountNoExists(account string) bool {
	err := dao.QueryUserAccount(account)
	if err != nil {
		fmt.Println("账号不存在")
		return true
	}
	fmt.Println("该账号已存在")
	return false
}

// 如果不存在就添加到数据库中
func AddAccount(account string, password string) {
	if CheckAccountNoExists(account) {
		err := dao.AccountInsert(account, password)
		if err != nil {
			fmt.Println(err)
			fmt.Println("添加失败")
		} else {
			fmt.Println("添加成功")
		}
	}
}
