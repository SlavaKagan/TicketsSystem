package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB = nil

/* Connect to the wix database */

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./wix.db")
	checkErr(err)

	database = db
}

/* Open the database itself with the data*/

func GetDB() *sql.DB {
	if database == nil {
		panic("database not initialized")
	}
	return database
}

/* function for checking varieties of errors */

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
