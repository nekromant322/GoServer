package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type TeacherGroup struct {
	GroupID string
	Group   string
	Course  string
	Info    string
}
type InfoMarks struct {
	Name string
}
type StudentMarks struct {
	Name  string
	Marks []LessonMarks
}
type LessonMarks struct {
	Lesson    int
	HomeMark  int
	ClassMark int
}

func GetStudentMarks(groupID int) []StudentMarks {
	var studentsMarks []StudentMarks
	var studentMarks StudentMarks
	var logins []string
	var login string
	studentSQL := "SELECT DISTINCT login FROM MARKS WHERE groupID=?"
	log.Print("Getting users for group ", groupID)
	rows := database.query(studentSQL, groupID)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&login)
		logins = append(logins, login)
		studentSQL = "SELECT real_name FROM USERS WHERE login=?"
		log.Print("Getting real_name for user ", login)
		rows := database.query(studentSQL, login)
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&studentMarks.Name)
		}
		markSQL := "SELECT lesson_number, class_mark, home_mark FROM MARKS WHERE (login = ?)AND (groupID=?) ORDER BY lesson_number ASC"
		log.Print("Getting marks for user ", login, groupID)
		rows = database.query(markSQL, login, groupID)
		defer rows.Close()
		var lessonsMarks []LessonMarks
		var lessonMarks LessonMarks
		for rows.Next() {
			rows.Scan(&lessonMarks.Lesson, &lessonMarks.ClassMark, &lessonMarks.HomeMark)
			lessonsMarks = append(lessonsMarks, lessonMarks)
		}
		studentMarks.Marks = lessonsMarks
		studentsMarks = append(studentsMarks, studentMarks)
	}
	return studentsMarks
}

//GetTeacherGroupList returns groups for a teacher
func GetTeacherGroupList(login string) []TeacherGroup {
	var teacherGroups []TeacherGroup
	groupsSQL := "SELECT groupID, group_name, COURSES.name, info FROM GROUPS, COURSES WHERE (teacher =?) AND (COURSES.courseID = GROUPS.courseID)"
	log.Print("Getting groups of teacher ", login)
	rows := database.query(groupsSQL, login)
	defer rows.Close()
	var teacherGroup TeacherGroup
	for rows.Next() {
		rows.Scan(&teacherGroup.GroupID, &teacherGroup.Group, &teacherGroup.Course, &teacherGroup.Info)
		teacherGroups = append(teacherGroups, teacherGroup)
	}
	return teacherGroups
}
