package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	port := ":8080"
	println("Server listen on port", port)
	http.HandleFunc("/", mainPage);
	http.HandleFunc("/login", loginPage);

	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}


}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl , err := template.ParseFiles("static/index.html")
	if err != nil{
		http.Error(w,err.Error(),400)
		return
	}
	err = tmpl.Execute(w,nil)
	if  err !=nil{
	    http.Error(w,err.Error(),400)

	    return
	}


}
func loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl , err := template.ParseFiles("static/new_login.html")
	if err != nil{
		http.Error(w,err.Error(),400)
		return
	}
	if err := tmpl.Execute(w,nil) ; err !=nil{
		http.Error(w,err.Error(),400)
		print("qwe");
		return
	}


}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go


