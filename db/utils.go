package db

import (
	"database/sql"
	"errors"
	"log"
	"net/smtp"
	"os"

	"github.com/sethvargo/go-password/password"
)

var err error
var database Database

//Database encapsulates database
type Database struct {
	db *sql.DB
}

func init() {
	database.db, err = sql.Open("sqlite3", "./serverDB.db")
	if err != nil {
		log.Fatal(err)
	}
}
func (db Database) query(q string, args ...interface{}) (rows *sql.Rows) {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rows
}
func (db Database) prepare(q string) (stmt *sql.Stmt) {
	stmt, err := db.db.Prepare(q)
	if err != nil {
		log.Println(err)
		return nil
	}
	return stmt
}
func (db Database) begin() (tx *sql.Tx) {
	tx, err := db.db.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	return tx
}
func insertQuery(sql string, args ...interface{}) error {
	log.Print("inside insert query")
	SQL := database.prepare(sql)
	tx := database.begin()
	_, err = tx.Stmt(SQL).Exec(args...)
	if err != nil {
		log.Println("inserQuery: ", err)
		tx.Rollback()
	} else {
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("Commit successful")
	}
	return err
}
func GetHash(login string, password string) (string, error) {

	key := 1234567
	if len(login) < 5 {
		err := errors.New("email is too short")
		return "", err
	}

	sol := login[0:4]

	hash := password + sol
	arr := make([]int, len(hash))
	for i := 0; i < len(hash); i++ {
		arr[i] = int(hash[i])
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < key; j++ {
			arr[i] = arr[i] * arr[j%len(arr)]
			arr[i] = arr[i] % key
		}
		if arr[i]%5 == 0 || arr[i]%11 == 0 || arr[i]%7 == 0 || arr[i]%13 == 0 {
			arr[i] = 48 + arr[i]%10
		} else {

			arr[i] = 65 + arr[i]%26
		}

	}
	hash = ""
	for i := 0; i < len(arr); i++ {
		character := string(arr[i])
		hash = hash + character
	}

	return hash, nil
}

func Send(email string, body string) error {
	from := "mietcko@gmail.com"
	pass := os.Getenv("MAIL_PASS")
	to := email
	log.Printf(pass)
	log.Printf(email)
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: ЦКО МИЭТ\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}

	log.Print("email sent")
	return nil
}

func SendPassword(login string) error {
	pass, err := password.Generate(6, 3, 0, false, false)
	if err != nil {
		return err
	}
	hash, err := GetHash(login, pass)
	if err != nil {
		return err
	}
	log.Printf(pass)
	err = savePassword(login, hash)
	if err != nil {
		return err
	}
	message := "Ваш пароль для доступа к ckomiet.ru: " + pass
	Send(login, message)
	if err != nil {
		return err
	}
	return nil
}
