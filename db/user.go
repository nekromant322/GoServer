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
type Mark struct {
	Login         string `json:"login"`
	Lesson_number int    `json:"lesson_number"`
	Class_mark    int    `json:"class_mark"`
	Home_mark     int    `json:"home_mark"`
	Group         string `json:"group"`
}

//ValidUser will check if the user exists in db and if it does, checks if the login/password combination is valid
func ValidUser(login, password string) bool {
	//password = GetHash(login,password); раскоментить когда в БД будут хранится хеши
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
		log.Print("successfully validated")
		return true
	}
	log.Print("username and password don't match")
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
	groupSQL := "SELECT group_name, name, real_name, amount FROM GROUPS, COURSES, USERINFO WHERE (groupID =?) AND (GROUPS.courseID = COURSES.courseID) AND (teacher=USERINFO.login)"
	log.Print("Getting group info for user ", group)
	rows = database.query(groupSQL, group)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&groupInfo.Group, &groupInfo.CourseName, &groupInfo.Teacher, &groupInfo.Amount)
	}
	return groupInfo
}
func GetMarkInfo(login string) []Mark{
	markSQL := "SELECT login, lesson_number, class_mark, home_mark, groupID FROM MARKS WHERE login = ?";
	log.Print("Getting marks for user ", login)
	rows := database.query(markSQL, login);
	defer rows.Close()
	var marks []Mark;
	var mark Mark;
	for rows.Next() {
		rows.Scan(&mark.Login, &mark.Lesson_number, &mark.Class_mark, &mark.Home_mark , &mark.Group)
		marks = append(marks, mark)
	}
	return marks;

}
