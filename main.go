package main

import (
	"GoServer/views"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	bdtest()
	port := ":8080"
	println("Server listen on port", port)
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/test", testPage)
	http.HandleFunc("/login", CreateHandler)
	http.HandleFunc("/androidLogin", views.AndroidLogin)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go
