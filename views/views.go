package views

import (
	"GoServer/db"
	"GoServer/sessions"
	"encoding/json"
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

//AndroidLogin is used to get all data about the user for android app
func AndroidLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	login := r.Form.Get("login")
	password := r.Form.Get("password")

	log.Println(login, password)
	if (login == "") || (password == "") {
		log.Println("Bad request")
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		isUser := db.ValidUser(login, password)
		if !isUser {
			log.Println("Unable to log user in")
			http.Error(w, "Unable to log user in", http.StatusInternalServerError)
		} else {
			rank, err := db.GetRank(login)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			if rank == 0 {
				var groupsInfo []db.GroupInfo
				groups := db.GetGroups(login)
				for _, group := range groups {
					groupsInfo = append(groupsInfo, db.GetGroupInfo(group, login))
				}
				data, _ := json.Marshal(groupsInfo)
				w.Write([]byte(data))
			}
		}
	}
}
func MainPage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/css")
	if sessions.IsLoggedIn(r) {
		tmpl, err := template.ParseFiles("templates/main.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), 400)

			return
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func TestPage(w http.ResponseWriter, r *http.Request) {
	marks := []Mark{Mark{"8160327", 1, 6, 2, "SB1230"}, Mark{"8160327", 2, 5, 0, "SB1230"}}
	tmpl, err := template.ParseFiles("templates/index.html")
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
	tmpl, err := template.ParseFiles("templates/new_login.html")
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
		log.Println(login, " ", password)
		if err != nil {
			log.Println(err)
		}
		session, err := sessions.Store.Get(r, "session")
		if db.ValidUser(login, password) {
			session.Values["loggedin"] = "true"
			session.Values["username"] = login
			session.Save(r, w)
			http.Redirect(w, r, "/", 301)
		} else {
			http.ServeFile(w, r, "templates/login_fail.html")
		}

	} else {
		http.ServeFile(w, r, "templates/new_login.html")
	}
}

//LogoutFunc Implements the logout functionality.
//WIll delete the session information from the cookie store
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	if err == nil {
		if session.Values["loggedin"] != "false" {
			session.Values["loggedin"] = "false"
			session.Save(r, w)
		}
	}
	http.Redirect(w, r, "/login", 302)
}
