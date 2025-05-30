package showPage

import (
	"net/http"
)

func BackHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound) //告诉浏览器要跳转页面
}
