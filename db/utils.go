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
func GetHash(login string, password string) string{

	key := 1234567;
	sol := login[0:4];

	hash := password + sol;
	arr := make([] int, len(hash));
	for i := 0; i < len(hash); i++ {
		arr[i] = int(hash[i]);
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < key; j++ {
			arr[i] = arr[i] * arr[j%len(arr)];
			arr[i] = arr[i]%key;
		}
		if(arr[i] % 5 == 0 || arr[i] %11 ==0){
			arr[i] = 65 + arr[i] % 26;
		} else {
			arr[i] = 47 + arr[i] % 10;
		}

	}
	hash = "";
	for i := 0; i < len(hash); i++ {
		character := string(arr[i]);
		hash =hash + character;
	}

	return hash;
}