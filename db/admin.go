package db

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type UserShort struct {
	Login string
	Name  string
}
type CourseShort struct {
	CourseID   int
	CourseName string
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

func GetAllStudents(groupID int) []UserShort {
	var studentsShort []UserShort
	studentsSQL := "SELECT login, real_name from USERS where (rank = 0) AND login NOT IN (SELECT login from MARKS where groupID=?)"
	log.Print("Getting students for admin")
	rows := database.query(studentsSQL, groupID)
	defer rows.Close()
	for rows.Next() {
		var UserShort UserShort
		rows.Scan(&UserShort.Login, &UserShort.Name)
		studentsShort = append(studentsShort, UserShort)

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

func AddUser(rank, name, login, bday, bonus string) error {
	userSQL := "INSERT INTO USERS(login, password, rank, real_name, birthday, bonus_info) values (?,'-',?,?,?,?)"
	err := insertQuery(userSQL, login, rank, name, bday, bonus)
	if err != nil {
		return err
	}
	err = SendPassword(login)
	if err != nil {
		return err
	}
	return nil
}

func AddCourse(courseName string, amount int) error {
	userSQL := "INSERT INTO COURSES(name, amount) VALuES (?, ?)"
	err := insertQuery(userSQL, courseName, amount)
	if err != nil {
		return err
	}
	courseSQL := "SELECT MAX(courseID) from COURSES"
	var courseID int
	rows := database.query(courseSQL)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&courseID)
		if err != nil {
			return err
		}
	}
	rows.Close()
	for i := 1; i <= amount; i++ {
		lessonSQL := "INSERT INTO LESSONS (courseID, lesson_number) VALUES (?, ?)"
		err := insertQuery(lessonSQL, courseID, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTeachersAndCourses() ([]UserShort, []CourseShort) {
	var teachersShort []UserShort
	teacherSQL := "SELECT login, real_name from USERS where rank = 1"
	log.Print("Getting all teachers for admin")
	rows := database.query(teacherSQL)
	defer rows.Close()
	for rows.Next() {
		var userShort UserShort
		rows.Scan(&userShort.Login, &userShort.Name)
		teachersShort = append(teachersShort, userShort)
	}

	var coursesShort []CourseShort
	courseSQL := "SELECT courseID, name from COURSES"
	log.Print("Getting all courses for admin")
	rowsCourse := database.query(courseSQL)
	defer rowsCourse.Close()
	for rowsCourse.Next() {
		var courseShort CourseShort
		rowsCourse.Scan(&courseShort.CourseID, &courseShort.CourseName)
		coursesShort = append(coursesShort, courseShort)
	}

	return teachersShort, coursesShort
}

func DeleteCourse(courseID int) error {
	log.Print("Deleting course " + strconv.Itoa(courseID))

	deleteCourseSQL := "DELETE FROM MARKS WHERE groupID in (select groupID from groups where courseID=?)"
	err := insertQuery(deleteCourseSQL, courseID)
	if err != nil {
		return err
	}

	deleteCourseSQL = "DELETE FROM events WHERE groupID in (select groupID from courses where courseID=?)"
	err = insertQuery(deleteCourseSQL, courseID)
	if err != nil {
		return err
	}
	deleteCourseSQL = "DELETE FROM courses WHERE (courseID=?)"
	err = insertQuery(deleteCourseSQL, courseID)
	if err != nil {
		return err
	}
	deleteCourseSQL = "DELETE FROM groups WHERE (courseID=?)"
	err = insertQuery(deleteCourseSQL, courseID)
	if err != nil {
		return err
	}
	deleteCourseSQL = "DELETE FROM lessons WHERE (courseID=?)"
	err = insertQuery(deleteCourseSQL, courseID)
	if err != nil {
		return err
	}
	return nil
}

func AddGroup(name string, courseID int, teacher, info string) error {
	log.Println("Adding group")

	groupSQL := "INSERT INTO GROUPS(group_name, courseID, teacher, info) VALuES (?, ?, ?, ?)"
	err := insertQuery(groupSQL, name, courseID, teacher, info)
	if err != nil {
		return err
	}

	return nil
}
