package main

import (
	"GoServer/db"
	"GoServer/views"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	check_login := "mysimplelogin"
	check_pass := "mysimplepass"
	hash := db.GetHash(check_login, check_pass)
	println("Check Hash :", hash)
	port := ":8080"
	println("Server listen on port", port)
	r := mux.NewRouter()
	r.HandleFunc("/", views.MainPage)
	r.HandleFunc("/teacher", views.TeacherPage)
	r.HandleFunc("/test", views.TestPage)
	r.HandleFunc("/login", views.CreateHandler)
	r.HandleFunc("/androidlogin", views.AndroidLogin)
	r.HandleFunc("/group/{group}", views.GroupsMarks)

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
