package user

import (
	"html/template"
	"log"
	"net/http"
	"webb/model"
	"webb/model/param"
	"webb/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("view/pages/user/register.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		t, err := template.ParseFiles("view/pages/user/register.html")
		if err != nil {
			log.Fatal(err)
		}
		err = r.ParseForm()
		if err != nil {
			log.Fatal(err)
			return
		}

		account := r.FormValue("account")
		password := r.FormValue("password")
		confirm_password := r.FormValue("confirm_password")
		u := model.User{
			Account:          account,
			Password:         password,
			Confrom_Password: confirm_password,
		}
		if !service.CheckAccountFormat(u.Account) {
			data := param.RegisterPageData{"账号格式不正确"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
			return
		} else if u.Password != u.Confrom_Password {
			data := param.RegisterPageData{"两次密码输入不一致"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
			return
		} else if !service.CheckAccountNoExists(u.Account) { //查询账号存不存在
			data := param.RegisterPageData{"该账号已存在"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
			return
		} else if service.CheckAccountNoExists(u.Account) {
			service.AddAccount(u.Account, u.Password)
			data := param.RegisterPageData{"注册成功，请返回登录"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
		}

	}

}
