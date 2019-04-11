package main

import (
	"GoServer/db"
	"GoServer/views"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	check_login := "mysimplelogin";
	check_pass := "mysimplepass";
	hash :=  db.GetHash(check_login,check_pass);
	println("Check Hash :",hash);
	port := ":8080"
	println("Server listen on port", port)
	http.HandleFunc("/", views.MainPage)
	http.HandleFunc("/test", views.TestPage)
	http.HandleFunc("/login", views.CreateHandler)
	http.HandleFunc("/androidlogin", views.AndroidLogin)
	http.Handle("/static/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go
