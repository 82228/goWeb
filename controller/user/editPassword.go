package user

import (
	"log"
	"net/http"
	"strconv"
	"webb/model"
	"webb/pkg/global"
	"webb/pkg/response"
	"webb/service"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.URL.Query().Get("id")
		w.Header().Set("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil {
			log.Fatal("解析表单有误", err)
			return
		}

		new_account := r.URL.Query().Get("newAccount")
		new_password := r.URL.Query().Get("password")

		//判断账号是否重名
		flag := service.CheckAccountExists(new_account)
		if flag {
			response.Json(w, http.StatusOK, true, "该用户名已存在，请重新输入", global.CodeSuccess)
			return
		}

		var u model.User
		//将字符串转换为整数
		u.Id, err = strconv.Atoi(id)
		if err != nil {
			response.Json(w, http.StatusBadRequest, false, "无效的请求参数", global.CodeFailed)
			return
		}

		//根据id更改账号密码,交给service层处理
		err = service.UpdateMessByID(new_account, new_password, u.Id)
		if err != nil {
			response.Json(w, http.StatusOK, true, "修改失败", global.CodeSuccess)
		} else {
			response.Json(w, http.StatusOK, true, "更新成功", global.CodeSuccess)
		}
	}
}
