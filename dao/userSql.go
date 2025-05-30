package dao

import (
	"database/sql"
	"log"
	"webb/model"
	"webb/model/param"
)

// 查询密码是否对应当前账号
func QueryUserPasswordByAccount(account string) (psd string, err error) {
	sqlStr := "select password from user where account = ?"
	var u model.User
	if err := db.QueryRow(sqlStr, account).Scan(&u.Password); err != nil {
		return "", err
	}
	return u.Password, nil
}

// 查询账号存不存在
func QueryUserAccount(account string) (err error) {
	sqlStr := "select account from user where account = ?"
	var u model.User
	err = db.QueryRow(sqlStr, account).Scan(&u.Account)
	if err != nil {
		return err
	}
	return nil
}

// 添加用户
func AccountInsert(account string, password string) (err error) {
	sqlStr := "insert into user (account, password, role) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, account, password, 0)
	if err != nil {
		return err
	}
	return nil
}

// 查询所有用户的账号信息
func QueryAllAccount() ([]model.User, error) {
	sqlStr := "select * from user"
	var u model.User
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var user []model.User
	//循环读取数据
	for rows.Next() {
		e := rows.Scan(&u.Id, &u.Account, &u.Password, &u.Role, &u.AvatarURL, &u.UserName, &u.Gender, &u.Birthday, &u.Province, &u.Address, &u.City)
		if e != nil {
			log.Fatal("读取数据失败-----------", e)
			return user, e
		}
		user = append(user, u)
	}
	return user, nil
}

// 删除账号
func DeleteById(id int) error {
	sqlStr := "delete from user where id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 根据账号寻找Role
func QueryRoleByAccount(account string) (role int) {
	sqlStr := "select role from user where account = ?"
	var h param.HomePageData
	err := db.QueryRow(sqlStr, account).Scan(&h.CurserRole)
	if err != nil {
		return
	}
	return h.CurserRole
}

// 添加用户头像的url
func AddUserAvatar(account string, avatarUrl string) (err error) {
	sqlStr := "update user set avatarURL = ? where account = ?"
	_, err = db.Exec(sqlStr, avatarUrl, account)
	if err != nil {
		log.Fatal("修改原始头像失败", err)
		return err
	}
	return nil
}

// 通过用户账号查询对应的url
func QueueUrl(account string) (url string) {
	sqlStr := "select avatarURL from user where account = ?"
	var u model.User
	err := db.QueryRow(sqlStr, account).Scan(&u.AvatarURL)
	if err != nil {
		return
	}
	return u.AvatarURL
}

// 在数据库中添加userName,用户名字
//func AddUserName(account string, name string) (err error) {
//	sqlStr := "update user set userName = ? where account = ?"
//	_, err = db.Exec(sqlStr, name, account)
//	if err != nil {
//		return err
//	} else {
//		log.Fatal("添加成功")
//	}
//	return nil
//}

// 通过账号获取信息
func MessSelect(account string) (u model.User) {
	sqlStr := "select userName , gender, avatarURL, birthday, province, address, city from user where account = ?"
	err := db.QueryRow(sqlStr, account).Scan(&u.UserName, &u.Gender, &u.AvatarURL, &u.Birthday, &u.Province, &u.Address, &u.City)
	if err != nil {
		log.Fatal("查询错误")
		return
	}
	return u
}

// 管理员根据id修改信息
func MessUpdateByID(account string, password string, id int) (err error) {
	sqlStr := "update user set account = ?, password = ? where id = ?"
	_, err = db.Exec(sqlStr, account, password, id)
	if err != nil {
		return err
	}
	return nil
}

// 添加信息
func AddMess(account string, u model.User) (err error) {
	sqlStr := "update user set userName=?, gender=?, birthday=?, province=?, city=?, address=? where account = ?"
	_, err = db.Exec(sqlStr, u.UserName, u.Gender, u.Birthday, u.Province, u.City, u.Address, account)
	if err != nil {
		return err
	}
	return nil
}
