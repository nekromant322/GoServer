package main

import (
	"html/template"
	"log"
	"net/http"
)

type Mark struct {
	Login         string `json:"login"`
	Lesson_number int    `json:"lesson_number"`
	Class_mark    int    `json:"class_mark"`
	Home_mark     int    `json:"home_mark"`
	Group         string `json:"group"`
}

func main() {

	port := ":8080"
	println("Server listen on port", port)
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/login", loginPage)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

func mainPage(w http.ResponseWriter, r *http.Request) {

	marks := []Mark{Mark{"8160327", 1, 6, 2, "SB1230"}, Mark{"8160327", 2, 5, 0, "SB1230"}}
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = tmpl.Execute(w, marks)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

}
func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/new_login.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		print("qwe")
		return
	}

}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go
