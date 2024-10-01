package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL, 
		type TEXT NOT NULL
	)
	`
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL, 
		user_id TEXT, 
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id TEXT PRIMARY KEY,
			event_id TEXT,
			user_id TEXT,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, errReg := DB.Exec(createRegistrationTable)

	if errReg != nil {
		panic(errReg)
	}
	_, errUsers := DB.Exec(createUsersTable)

	if errUsers != nil {
		panic(errUsers.Error())
	}

	_, errCreate := DB.Exec(createEventsTable)

	if errCreate != nil {
		panic(errCreate.Error())
	}
}
