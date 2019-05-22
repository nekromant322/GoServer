package main

import (
	"GoServer/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := ":8080"
	println("Server listen on port", port)
	r := mux.NewRouter()
	r.HandleFunc("/", views.MainPage)
	r.HandleFunc("/teacher", views.TeacherPage)
	r.HandleFunc("/login", views.CreateHandler)
	r.HandleFunc("/androidlogin", views.AndroidLogin)
	r.HandleFunc("/group/{group}", views.GroupsMarks)
	r.HandleFunc("/header", views.Header)
	r.HandleFunc("/groups", views.Groups)
	r.HandleFunc("/profile", views.Profile)
	r.HandleFunc("/logout", views.LogoutFunc)
	r.HandleFunc("/forgotpass", views.ForgotPassword)
	r.HandleFunc("/admin", views.Admin)
	r.HandleFunc("/admin_header", views.HeaderAdmin)
	r.HandleFunc("/admin_groups", views.GroupsAdmin)
	r.HandleFunc("/admin_main", views.MainAdmin)
	r.HandleFunc("/admin_group/{group}", views.GroupAdmin)
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.Handle("/", r)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go
