package main

import (
	"GoServer/views"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := ":8080"
	println("Server listen on port", port)
	http.HandleFunc("/", views.MainPage)
	http.HandleFunc("/test", views.TestPage)
	http.HandleFunc("/login", views.CreateHandler)
	http.HandleFunc("/androidlogin", views.AndroidLogin)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go
