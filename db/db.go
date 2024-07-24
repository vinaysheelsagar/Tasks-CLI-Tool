package db_controller

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func DoesDBFileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		log.Println("db file not found. Creating new one...")
		return false, nil

	} else {
		log.Fatal(err)
		return false, err

	}

}

func createDBFile(path string) {
	os.Create(path)
}

func CreateTaskTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "tasks" (
		"id"	INTEGER,
		"name"	TEXT,
		"desc"	TEXT,
		"priority"	TEXT,
		"category"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func CheckupDB(path string) {
	exists, err := DoesDBFileExists(path)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		createDBFile(path)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	CreateTaskTable(db)
	// CreateCategoryTable(db)
	// CreateSubTaskTable(db)

	defer db.Close()

}
