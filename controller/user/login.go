package user

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"webb/model"
	"webb/model/param"
	"webb/service"
)

// 登录
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("view/pages/user/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		t, err := template.ParseFiles("view/pages/user/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		account := r.FormValue("account")
		password := r.FormValue("password")
		u := model.User{
			Account:  account,
			Password: password,
		}

		// 交给service层处理
		//service.Check(u)
		if !service.CheckAccountFormat(u.Account) {
			data := param.LoginPageData{"账号格式不正确"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
			return
		} else if !service.CheckAccountExists(u.Account) {
			data := param.LoginPageData{"账号不存在，请重新输入"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
			return
		} else if !service.CheckPassword(u.Account, u.Password) {
			data := param.LoginPageData{"账号或者密码不正确"}
			err := t.Execute(w, data)
			if err != nil {
				return
			}
			return
		} else {
			//设置cookie，保存登录的账号
			http.SetCookie(w, &http.Cookie{
				Name:     "account",
				Value:    u.Account,
				Path:     "/",
				Expires:  time.Now().Add(time.Hour * 24),
				HttpOnly: true,
				Secure:   true,
			})
			http.Redirect(w, r, "/home", http.StatusFound) //告诉浏览器要跳转页面
		}
	}

}
