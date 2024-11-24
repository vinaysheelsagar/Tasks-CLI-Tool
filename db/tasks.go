package database

import (
	"database/sql"
	"log"

	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

type Task struct {
	id       int
	name     string
	priority int
	category Category
}

func TaskTableCheckup(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "tasks" (
		"id"			INTEGER,
		"name"			TEXT,
		"priority"		INTEGER,
		"category_id"	INTEGER,
		
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("category_id") REFERENCES categories("id")
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateTask(db *sql.DB) {
	_, err := db.Exec(`INSERT INTO 'tasks' (name, priority, category_id) VALUES(%s, %s, %s);`)
	utilities.CheckNil(err, "", "")
}
