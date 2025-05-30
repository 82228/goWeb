package user

import (
	"html/template"
	"log"
	"net/http"
	"webb/model"
	"webb/model/param"
	"webb/service/QueueUserMess"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/pages/public/home.html")

	//用cookie获得登录用户的账号
	c, e := r.Cookie("account")
	if e != nil {
		log.Fatal(e)
	}

	//用一个值接收查询所有用户账号返回的切片
	users := QueueUserMess.QueryAllMess()

	//将cookie值放到结构体中
	user := model.User{
		Account: c.Value,
	}

	CurRole := QueueUserMess.QueryRoleByCount(c.Value)
	//创建一个结构体存放当前账号和所有用户的信息
	data := param.HomePageData{
		CurUser:    user,  //当前账号
		UserList:   users, //所有账号的信息
		CurserRole: CurRole,
	}

	err := t.Execute(w, data)
	if err != nil {
		return
	}
}
