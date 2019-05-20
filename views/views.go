package views

import (
	"GoServer/db"
	"GoServer/sessions"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

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
		isUser, _ := db.ValidUser(login, password)
		if !isUser {
			log.Println("Unable to log user in")
			http.Error(w, "Unable to log user in", http.StatusInternalServerError)
		} else {
			var userInfo db.UserInfo
			realname, err := db.GetRealName(login)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			userInfo.RealName = realname
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
				userInfo.GroupsInfo = groupsInfo
				data, _ := json.Marshal(userInfo)
				w.Write([]byte(data))
			}
		}
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/css")
	session, _ := sessions.Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		tmpl, err := template.ParseFiles("templates/main.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		login := fmt.Sprintf("%v", session.Values["username"])
		rank, err := db.GetRank(login)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var userInfo db.UserInfo
		realname, err := db.GetRealName(login)
		userInfo.RealName = realname
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
			userInfo.GroupsInfo = groupsInfo
			err = tmpl.Execute(w, userInfo)
		} else {
			if rank == 1 {
				http.Redirect(w, r, "/teacher", 302)
			}
		}

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func GroupsMarks(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("templates/group.html")
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			login := fmt.Sprintf("%v", session.Values["username"])

			_, err = db.GetRank(login)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			var groupMarks db.GroupMarks
			groupNum, _ := strconv.Atoi(mux.Vars(r)["group"])
			courseID, _ := db.GetCourseID(groupNum)
			var shortLessonsInfo []db.ShortLessonInfo
			shortLessonsInfo = db.GetLessonInfo(courseID)
			groupMarks = db.GetStudentMarks(groupNum)
			markData := struct {
				CourseID   int
				Marks      db.GroupMarks
				LessonInfo []db.ShortLessonInfo
			}{courseID, groupMarks, shortLessonsInfo}
			err = tmpl.Execute(w, markData)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		} else {
			if r.Method == "POST" {
				err := r.ParseForm()
				if err != nil {
					log.Println(err)
				}
				groupNum, _ := strconv.Atoi(mux.Vars(r)["group"])
				courseID, _ := db.GetCourseID(groupNum)
				if r.FormValue("submit") == "save_marks" {
					for key, values := range r.Form {
						if strings.Split(key, ";")[0] == "home_mark" || strings.Split(key, ";")[0] == "class_mark" {
							err := db.SaveLessonData(key, values, groupNum, courseID)
							if err != nil {
								log.Println(err)
							}
						}
					}
				} else {
					for key, values := range r.Form {
						if (strings.Split(key, ";")[0] == "homework" || strings.Split(key, ";")[0] == "theme") && strings.Split(key, ";")[2] == r.FormValue("submit") {
							err := db.SaveLessonData(key, values, groupNum, courseID)
							if err != nil {
								log.Println(err)
							}
						}
					}
				}
				http.Redirect(w, r, "/group/"+strconv.Itoa(groupNum), 301)
			}
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
func Header(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/css")
	session, _ := sessions.Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		tmpl, err := template.ParseFiles("templates/header.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		login := fmt.Sprintf("%v", session.Values["username"])

		_, err = db.GetRank(login)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var teacherGroups []db.TeacherGroup
		teacherGroups = db.GetTeacherGroupList(login)
		err = tmpl.Execute(w, teacherGroups)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
func Groups(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/css")
	session, _ := sessions.Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		tmpl, err := template.ParseFiles("templates/groups.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		login := fmt.Sprintf("%v", session.Values["username"])

		_, err = db.GetRank(login)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var teacherGroups []db.TeacherGroup
		teacherGroups = db.GetTeacherGroupList(login)
		err = tmpl.Execute(w, teacherGroups)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func Profile(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/css")
	session, _ := sessions.Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("templates/profile.html")
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			login := fmt.Sprintf("%v", session.Values["username"])

			_, err = db.GetRank(login)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			var teacherGroups []db.TeacherGroup
			teacherGroups = db.GetTeacherGroupList(login)
			name, _ := db.GetRealName(login)
			info, _ := db.GetBonusInfo(login)
			teacherInfoAndGroups := struct {
				TeacherGroups []db.TeacherGroup
				Teacher       string
				TeacherInfo   string
			}{
				teacherGroups,
				name,
				info,
			}
			err = tmpl.Execute(w, teacherInfoAndGroups)

			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		} else {
			if r.Method == "POST" {
				err := r.ParseForm()
				if err != nil {
					log.Println(err)
				}
				switch r.FormValue("submit") {
				case "save_info":
					teacherInfo := r.FormValue("teacher_info")
					login := fmt.Sprintf("%v", session.Values["username"])
					db.SaveTeacherBonusInfo(login, teacherInfo)
					http.Redirect(w, r, "/profile", 301)
				case "save_event":
					var groupIDs []string
					var eventText string
					for key, values := range r.Form { // range over map
						for _, value := range values { // range over []string
							if key == "groupCheck" {
								groupIDs = append(groupIDs, value)
							}
							if key == "eventText" {
								eventText = value
							}
						}
					}
					db.SaveEventData(groupIDs, eventText)
					http.Redirect(w, r, "/profile", 301)
				default:
					err = db.DeleteEvent(r.FormValue("submit"))
					if err != nil {
						log.Println(err)
					}
					http.Redirect(w, r, "/profile", 302)
				}
			}
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func TeacherPage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/css")
	session, _ := sessions.Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
		tmpl, err := template.ParseFiles("templates/frames.html") //need to parse frames.html
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		login := fmt.Sprintf("%v", session.Values["username"])

		_, err = db.GetRank(login)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var teacherGroups []db.TeacherGroup
		teacherGroups = db.GetTeacherGroupList(login)
		err = tmpl.Execute(w, teacherGroups)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
func TestPage(w http.ResponseWriter, r *http.Request) {
	/*marks := []Mark{Mark{"8160327", 1, 6, 2, "SB1230"}, Mark{"8160327", 2, 5, 0, "SB1230"}}
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
	*/
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

		var check bool
		var rank int
		check, rank = db.ValidUser(login, password)
		if check {
			session.Values["loggedin"] = "true"
			session.Values["username"] = login
			session.Save(r, w)
			if rank == 0 {
				http.Redirect(w, r, "/", 301)
			}
			//Need to fix
			if rank == 1 {
				http.Redirect(w, r, "/teacher", 301)
			}
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
