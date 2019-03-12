package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":8080"
	println("Server listen on port", port)
	http.HandleFunc("/", mainPage);
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
	if err := tmpl.Execute(w,nil) ; err !=nil{
	    http.Error(w,err.Error(),400)
	    return
	}


}




func getUser(c *gin.Context) {
	id := c.Query("id")
	user := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{id, "Вася"}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func postUser(c *gin.Context) {
	user := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{}
	c.BindJSON(&user)

	user.Name = "Петя"

	c.JSON(200, gin.H{
		"user": user,
	})
}
