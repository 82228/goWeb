package user

import (
	"encoding/json"
	"log"
	"net/http"
	"webb/model"
	"webb/model/param"
	"webb/service/AddUserMess"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		//获取用户的信息
		username := r.FormValue("username") //用户名
		gender := r.FormValue("gender")     //性别
		birthday := r.FormValue("birthday") //生日
		province := r.FormValue("province") //省份
		city := r.FormValue("city")         //市
		address := r.FormValue("address")   //详细地址

		//绑定到结构体里面
		var u model.User
		u = model.User{
			UserName: username,
			Gender:   gender,
			Birthday: birthday,
			Province: province,
			City:     city,
			Address:  address,
		}
		//通过cookie获取用户账号
		c, err := r.Cookie("account")
		if err != nil {
			log.Fatal("cookie获取错误", err)
		}
		//交给service层处理,将账号,性别,生日,省份,详细地址 存到数据库中
		err = AddUserMess.PersonalMessAdd(c.Value, u)
		if err != nil {
			log.Fatal(err)
		}

		data, _ := json.Marshal(param.Response{
			Code: 200,
		})

		_, err = w.Write(data)
		if err != nil {
			return
		}
	}
}
