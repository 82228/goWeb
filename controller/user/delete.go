package user

import (
	"net/http"
	"strconv"
	"webb/model"
	"webb/pkg/global"
	"webb/pkg/response"
	"webb/service/DelUserMess"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var u model.User
		id := r.FormValue("id")

		var err error
		u.Id, err = strconv.Atoi(id)
		if err != nil {
			response.Json(w, http.StatusBadRequest, false, "无效的请求参数", global.CodeFailed)
			return
		}

		err = DelUserMess.Delete(u.Id)
		if err != nil {
			response.Json(w, http.StatusBadRequest, false, "删除失败", global.CodeFailed)
			return
		}
		// 返回成功响应
		response.Json(w, http.StatusOK, true, "删除成功", global.CodeSuccess)
	}
}
