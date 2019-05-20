package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type TeacherGroup struct {
	GroupID string
	Group   string
	Course  string
	Info    string
	Events  []Event
}
type InfoMarks struct {
	Name string
}
type StudentMarks struct {
	Login string
	Name  string
	Marks []LessonMarks
}
type LessonMarks struct {
	Lesson    int
	HomeMark  int
	ClassMark int
}
type GroupMarks struct {
	Group         int
	StudentsMarks []StudentMarks
}
type ShortLessonInfo struct {
	LessonNumber int
	Homework     string
	Theme        string
}
type FullTeacherInfo struct {
	Marks      GroupMarks
	LessonInfo ShortLessonInfo
}

func GetStudentMarks(groupID int) GroupMarks {
	var studentsMarks []StudentMarks
	var logins []string
	studentSQL := "SELECT DISTINCT login FROM MARKS WHERE groupID=?"
	log.Print("Getting users for group ", groupID)
	rows := database.query(studentSQL, groupID)
	defer rows.Close()
	for rows.Next() {
		var studentMarks StudentMarks
		var login string
		rows.Scan(&login)
		logins = append(logins, login)
		studentMarks.Login = login
		studentSQL = "SELECT real_name FROM USERS WHERE login=?"
		log.Print("Getting real_name for user ", login)
		rowsLogin := database.query(studentSQL, login)
		defer rowsLogin.Close()
		for rowsLogin.Next() {
			rowsLogin.Scan(&studentMarks.Name)
		}
		markSQL := "SELECT lesson_number, class_mark, home_mark FROM MARKS WHERE (login = ?)AND (groupID=?) ORDER BY lesson_number ASC"
		log.Print("Getting marks for user ", login, groupID)
		rowsMarks := database.query(markSQL, login, groupID)
		defer rowsMarks.Close()
		var lessonsMarks []LessonMarks
		for rowsMarks.Next() {
			var lessonMarks LessonMarks
			rowsMarks.Scan(&lessonMarks.Lesson, &lessonMarks.ClassMark, &lessonMarks.HomeMark)
			lessonsMarks = append(lessonsMarks, lessonMarks)
		}
		studentMarks.Marks = lessonsMarks
		studentsMarks = append(studentsMarks, studentMarks)
	}
	var groupMarks GroupMarks
	groupMarks.Group = groupID
	groupMarks.StudentsMarks = studentsMarks
	return groupMarks
}
func GetCourseID(groupID int) (courseID int, err error) {
	rankSQL := "SELECT courseID FROM GROUPS WHERE groupID=?"
	log.Print("Getting rank for user ", groupID)
	rows := database.query(rankSQL, groupID)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&courseID)
		if err != nil {
			return 0, fmt.Errorf("No such group")
		}
	}
	return courseID, nil
}
func GetLessonInfo(courseID int) []ShortLessonInfo {
	var shortLessonsInfo []ShortLessonInfo
	groupsSQL := "SELECT lesson_number, homework, theme FROM LESSONS WHERE courseID = ?"
	log.Print("Getting course info for course ", courseID)
	rows := database.query(groupsSQL, courseID)
	defer rows.Close()
	for rows.Next() {
		var shortLessonInfo ShortLessonInfo
		rows.Scan(&shortLessonInfo.LessonNumber, &shortLessonInfo.Homework, &shortLessonInfo.Theme)
		shortLessonsInfo = append(shortLessonsInfo, shortLessonInfo)
	}
	return shortLessonsInfo
}
func SaveLessonData(key string, value []string, groupID int, courseID int) error {
	sKey := strings.Split(key, ";")
	dataType := sKey[0]
	if dataType == "home_mark" || dataType == "class_mark" {
		marksSQL := "UPDATE MARKS SET " + dataType + " = ? WHERE (login = ?) AND (lesson_number = ?) AND (groupID = ?)"
		err := insertQuery(marksSQL, value[0], sKey[1], sKey[2], groupID)
		if err != nil {
			return err
		}
	}
	if dataType == "theme" || dataType == "homework" {
		marksSQL := "UPDATE LESSONS SET " + dataType + " = ? WHERE (courseID = ?) AND (lesson_number = ?)"
		err := insertQuery(marksSQL, value[0], courseID, sKey[2])
		if err != nil {
			return err
		}
	}
	return nil
}

//GetTeacherGroupList returns groups for a teacher
func GetTeacherGroupList(login string) []TeacherGroup {
	var teacherGroups []TeacherGroup
	groupsSQL := "SELECT groupID, group_name, COURSES.name, info FROM GROUPS, COURSES WHERE (teacher =?) AND (COURSES.courseID = GROUPS.courseID)"
	log.Print("Getting groups of teacher ", login)
	rows := database.query(groupsSQL, login)
	defer rows.Close()
	for rows.Next() {
		var teacherGroup TeacherGroup
		rows.Scan(&teacherGroup.GroupID, &teacherGroup.Group, &teacherGroup.Course, &teacherGroup.Info)
		eventsSQL := "SELECT eventID, event, date FROM EVENTS WHERE groupID=? ORDER BY rowid DESC LIMIT 10"
		log.Print("Getting events for group ", teacherGroup.GroupID)
		rowsEvent := database.query(eventsSQL, teacherGroup.GroupID)
		defer rowsEvent.Close()
		var eventsInfo []Event
		for rowsEvent.Next() {
			var event Event
			rowsEvent.Scan(&event.EventID, &event.EventText, &event.Date)
			eventsInfo = append(eventsInfo, event)
		}
		teacherGroup.Events = eventsInfo
		teacherGroups = append(teacherGroups, teacherGroup)

	}
	return teacherGroups
}

func DeleteEvent(id string) error {
	eventDelSQL := "DELETE FROM EVENTS WHERE eventID=?"
	err := insertQuery(eventDelSQL, id)
	if err != nil {
		return err
	}
	return nil
}

func SaveEventData(groupIDs []string, eventText string) error {
	for _, groupID := range groupIDs {
		currentTime := time.Now()
		eventSQL := "INSERT INTO EVENTS (groupID, date, event) VALUES (?,?,?)"
		err := insertQuery(eventSQL, groupID, currentTime.Format("2006-01-02"), eventText)
		if err != nil {
			return err
		}
	}
	return nil
}
func SaveTeacherBonusInfo(login string, bonusInfo string) error {
	infoSQL := "UPDATE USERS SET bonus_info=? WHERE login=?"
	err := insertQuery(infoSQL, bonusInfo, login)
	if err != nil {
		return err
	}
	return nil
}
