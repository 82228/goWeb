package response

import (
	"encoding/json"
	"net/http"
	"webb/pkg/global"
)

func Json(w http.ResponseWriter, statusCode int, success bool, mess string, code global.AppCode) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    int    `json:"code"`
	}{
		Success: success,
		Message: mess,
		Code:    int(code),
	})
	if err != nil {
		return
	}
}
