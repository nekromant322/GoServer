package main

import (
<<<<<<< HEAD
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

=======
	"fmt"
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

>>>>>>> 171e61a6edc275d644a69c51e770d23f7d9ee105
func main() {

	port := ":8080"
	println("Server listen on port", port)
<<<<<<< HEAD
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
=======
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/test", testPage)
	http.HandleFunc("/login", CreateHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
func mainPage(w http.ResponseWriter, r *http.Request) {

	marks := []Mark{Mark{"8160327", 1, 6, 2, "SB1230"}, Mark{"8160327", 2, 5, 0, "SB1230"}}
	tmpl, err := template.ParseFiles("static/main.html")
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
func testPage(w http.ResponseWriter, r *http.Request) {

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
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		login := r.FormValue("login")
		password := r.FormValue("password")
		fmt.Errorf(login, password)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "static/new_login.html")
	}
}

//ссылка на файл со скриптом
//https://github.com/thewhitetulip/Tasks/blob/master/main.go
//https://github.com/thewhitetulip/Tasks/blob/master/views/addViews.go
>>>>>>> 171e61a6edc275d644a69c51e770d23f7d9ee105
