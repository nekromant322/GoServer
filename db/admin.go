package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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
