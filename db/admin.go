package db

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type StudentShort struct {
	Login string
	Name  string
}

func GetAllGroups() []TeacherGroup {
	var teacherGroups []TeacherGroup
	groupsSQL := "SELECT groupID, group_name, COURSES.name, info FROM GROUPS, COURSES WHERE (COURSES.courseID = GROUPS.courseID)"
	log.Print("Getting groups for admin")
	rows := database.query(groupsSQL)
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

func GetAllStudents(groupID int) []StudentShort {
	var studentsShort []StudentShort
	studentsSQL := "SELECT login, real_name from USERS where (rank = 0) AND login NOT IN (SELECT login from MARKS where groupID=?)"
	log.Print("Getting students for admin")
	rows := database.query(studentsSQL, groupID)
	defer rows.Close()
	for rows.Next() {
		var studentShort StudentShort
		rows.Scan(&studentShort.Login, &studentShort.Name)
		studentsShort = append(studentsShort, studentShort)

	}
	return studentsShort
}

func AddStudent(groupID int, login string) error {
	log.Print("Getting amount of hours for group " + strconv.Itoa(groupID))
	amountSQL := "SELECT amount from COURSES where courseID = (SELECT courseID from groups where groupID=?)"
	rows := database.query(amountSQL, groupID)
	var amount int
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&amount)
		if err != nil {
			return err
		}
	}
	rows.Close()
	log.Print("Adding student " + login + " to group " + strconv.Itoa(groupID))
	for i := 1; i <= amount; i++ {
		studentSQL := "INSERT INTO MARKS (login, lesson_number, groupID) VALUES (?,?,?)"
		err := insertQuery(studentSQL, login, i, groupID)
		if err != nil {
			return err
		}
	}
	return nil
}

func DelStudentFromGroup(groupID int, login string) error {
	log.Print("Deleting student " + login + " from group " + strconv.Itoa(groupID))

	studentDelSQL := "DELETE FROM MARKS WHERE (groupID=?) AND (login=?)"
	err := insertQuery(studentDelSQL, groupID, login)
	if err != nil {
		return err
	}
	return nil
}
