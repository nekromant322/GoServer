package db

import (
	"database/sql"
	"log"
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
