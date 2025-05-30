package middleware

import (
	"net/http"
	"webb/pkg/global"
	"webb/pkg/response"
	"webb/service/QueueUserMess"
)

func AuthorizationAuth(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取cookie
		account, err := r.Cookie("account")
		if err != nil {
			response.Json(w, http.StatusOK, false, "未获取到当前账号", global.CodeFailed)
		}
		// 查询角色,根据account查询role
		role := QueueUserMess.QueryRoleByCount(account.Value)
		// 判断
		if role != 1 {
			// 数据发送给前端
			response.Json(w, http.StatusOK, false, "不是管理员，没有权限", global.CodeFailed)
			return
		}
		handler(w, r)
	})
}
