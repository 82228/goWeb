package main

import (
	"net/http"
	file2 "webb/controller/file"
	"webb/controller/middleware"
	"webb/controller/showPage"
	"webb/controller/user"
	"webb/dao"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		return
	}
	server := http.Server{Addr: ":8080"}
	http.HandleFunc("/user/profile", user.Profile)
	http.HandleFunc("/logout", user.Logout)
	http.HandleFunc("/user/backHomePage", showPage.BackHomePage)
	http.HandleFunc("/edit-user", middleware.CheckEditPass(user.EditUser))
	http.HandleFunc("/user/saveChanges", user.ShowPage)
	http.HandleFunc("/user/register", user.Register)
	http.HandleFunc("/user/delete", middleware.AuthorizationAuth(user.DeleteUser))
	http.HandleFunc("/user/personalHomePage", user.ShowPage)
	http.HandleFunc("/user/CouldeditPage", showPage.CouldeditPage)
	http.HandleFunc("/user/editMess", file2.UploadHeadPhoto)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/home", user.Home)
	//http.HandleFunc("/user/profile", file.Profile)

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	err = server.ListenAndServe()
	if err != nil {
		return
	}
}
