package db

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type GroupInfo struct {
	Lessons    []LessonInfo `json:"Lessons"`
	Group      string       `json:"GroupName"`
	CourseName string       `json:"CourseName"`
	Teacher    string       `json:"Teacher"`
	Amount     int          `json:"Amount"`
}
type LessonInfo struct {
	LessonNumber int    `json:"LessonNumber"`
	ClassMark    int    `json:"ClassMark"`
	HomeMark     int    `json:"HomeMark"`
	Homework     string `json:"Homework"`
	Theme        string `json:"Theme"`
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
func GetRank(login string) (rankFromDB int8, err error) {
	rankSQL := "SELECT rank FROM USERS WHERE login =?"
	log.Print("Getting rank for user ", login)
	rows := database.query(rankSQL, login)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&rankFromDB)
		if err != nil {
			return 0, fmt.Errorf("No such user")
		}
	}
	return rankFromDB, nil
}

//GetGroups returns a slice of groupIDs for a user
func GetGroups(login string) []int {
	var groups []int
	var groupID int
	groupsSQL := "SELECT groupID FROM MARKS WHERE login =? GROUP BY groupID"
	log.Print("Getting groups of user ", login)
	rows := database.query(groupsSQL, login)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&groupID)
		groups = append(groups, groupID)
	}
	return groups
}

//GetGroupInfo returns all info related to the user in certain group (marks, homework, course info)
func GetGroupInfo(group int, login string) GroupInfo {
	var groupInfo GroupInfo
	var lessonInfo LessonInfo
	var lessonsInfo []LessonInfo
	lessonSQL := "SELECT MARKS.lesson_number, theme, homework, class_mark, home_mark FROM LESSONS, MARKS, GROUPS WHERE (GROUPS.groupID = ?) AND (GROUPS.courseID=LESSONS.courseID) AND (MARKS.groupID =?) AND (LESSONS.lesson_number = MARKS.lesson_number) AND (login = ?);"
	log.Print("Getting lessons for group ", group)
	rows := database.query(lessonSQL, group, group, login)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&lessonInfo.LessonNumber, &lessonInfo.Theme, &lessonInfo.Homework, &lessonInfo.ClassMark, &lessonInfo.HomeMark)
		lessonsInfo = append(lessonsInfo, lessonInfo)
	}
	groupInfo.Lessons = lessonsInfo
	groupSQL := "SELECT group_name, name, teacher, amount FROM GROUPS, COURSES WHERE (groupID =?) AND (GROUPS.courseID = COURSES.courseID)"
	log.Print("Getting group info for user ", group)
	rows = database.query(groupSQL, group)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&groupInfo.Group, &groupInfo.CourseName, &groupInfo.Teacher, &groupInfo.Amount)
	}
	return groupInfo
}
