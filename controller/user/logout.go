package user

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	//清楚会话cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1, //立即过期
	})

	//重定向到登录页面
	http.Redirect(w, r, "/login", http.StatusFound)
}
