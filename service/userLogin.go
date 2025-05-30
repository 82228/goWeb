package service

import (
	"fmt"
	"regexp"
	"webb/dao"
)

// 检查账号格式问题
func CheckAccountFormat(account string) bool {
	str := account
	pattern := `^[0-9a-zA-Z]{6,18}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(str) {
		fmt.Println("账号格式不符合规范")
		return false
	}

	return true
}

// 检查账号存不存在
func CheckAccountExists(account string) bool {
	err := dao.QueryUserAccount(account)
	if err != nil {
		fmt.Println("账号不存在")
		return false
	}
	return true
}

// 检查密码正不正确
func CheckPassword(account string, password string) bool {
	// 从数据库里面取
	psd, err := dao.QueryUserPasswordByAccount(account)
	if err != nil {
		fmt.Println("查询失败")
		return false
	}
	if psd != password {
		fmt.Println("密码不正确")
		return false
	}
	return true
}
