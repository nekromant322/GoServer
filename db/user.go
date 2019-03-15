package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type GroupInfo struct {
	Lessons []LessonInfo `json:"lessons"`
	Group   string       `json:"group_name"`
	Course  string       `json:"course_name"`
	Teacher string       `json:"teacher_name"`
	Amount  string       `json:"amount"`
}
type LessonInfo struct {
	Lesson_number int `json:"lesson_number"`
	Class_mark    int `json:"class_mark"`
	Home_mark     int `json:"home_mark"`
	Homework      int `json:"homework"`
	Theme         int `json:"theme"`
}

//ValidUser will check if the user exists in db and if it does, checks if the login/password combination is valid
func ValidUser(login, password string) bool {
	var passwordFromDB string
	userSQL := "SELECT password FROM USERS WHERE login=?"
	log.Print("validating user ", login)
	rows := database.query(userSQL, login)

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&passwordFromDB)
		if err != nil {
			return false
		}
	}
	//If the password matches, return true
	if password == passwordFromDB {
		return true
	}
	//by default return false
	return false
}

//GetRank will return the rank of a user by his login
func GetRank(login string) int8 {
	//not finished yet
	return 0
}

//GetGroups returns a slice of groupIDs for a user
func GetGroups(login string) []int {
	//not finished yet
	var a []int
	return a
}

//GetGroupInfo returns all info related to the user in certain group (marks, homework, course info)
func GetGroupInfo(group int, login string) GroupInfo {
	//not finished yet
	var a GroupInfo
	return a
}
