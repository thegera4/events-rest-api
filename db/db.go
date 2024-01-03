package db

import (
	"database/sql"
	_ "modernc.org/sqlite" //underscore means we are not using it directly
)

var DB *sql.DB

func InitDB() {
	var initDBErr error
	DB, initDBErr = sql.Open("sqlite", "api.db")
	if initDBErr != nil {
		panic("Could not connect to the database!: " + initDBErr.Error())
	}

	DB.SetMaxOpenConns(10) //prevent too many connections to the database (pool)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create the users table!")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create the events table!")
	}
}