package user

import (
	"html/template"
	"log"
	"net/http"
	"webb/service/QueueUserMess"
)

// 进入个人主页，还不可编辑
func ShowPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "POST" {
		t, err := template.ParseFiles("view/pages/user/personalHomepage.html")
		if err != nil {
			log.Fatal(err)
			return
		}
		//通过cookie获取用户账号
		c, err := r.Cookie("account")
		if err != nil {
			log.Fatal(err)
		}
		//通过账号获得用户名字, 性别, 生日, 省份, 地址， 头像url
		u := QueueUserMess.PersonalMessQueue(c.Value)
		err = t.Execute(w, u)
		if err != nil {
			return
		}
	}
}
