package showPage

import (
	"html/template"
	"log"
	"net/http"
	"webb/service/QueueUserMess"
)

// 进入可编辑页面
func CouldeditPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("view/pages/user/ebitenPersonalHome.html")
		if err != nil {
			log.Fatal(err)
			return
		}
		//通过cookie获取用户账号
		c, err := r.Cookie("account")
		if err != nil {
			log.Fatal(err)
		}

		//通过账号获得用户名字, 性别, 生日, 省份, 地址， url
		u := QueueUserMess.PersonalMessQueue(c.Value)
		err = t.Execute(w, u)
		if err != nil {
			return
		}
	}
}
