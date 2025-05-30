package middleware

import (
	"log"
	"net/http"
	"webb/pkg/global"
	"webb/pkg/response"
	"webb/service"
)

func CheckEditPass(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json")
			err := r.ParseForm()
			if err != nil {
				log.Fatal("解析表单有误", err)
				return
			}

			new_account := r.URL.Query().Get("newAccount")
			ans := service.CheckAccountFormat(new_account)
			if !ans {
				response.Json(w, http.StatusOK, false, "账号格式不正确", global.CodeFailed)
				return
			}
			//判断账号是否重名
			flag := service.CheckAccountExists(new_account)
			if flag {
				response.Json(w, http.StatusOK, false, "该用户名已存在，请重新输入", global.CodeFailed)
				return
			}
			handler(w, r)
		}
	})
}
