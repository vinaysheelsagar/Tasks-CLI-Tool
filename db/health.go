package database

import (
	"errors"
	"log"
	"os"

	_ "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DbFilePath string
}

func doesDbFileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil

	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil

	} else {
		log.Println(err)
		return false, err
	}
}

func createDbFile(path string) {
	log.Println("Database file not found. Creating new one.")
	os.Create(path)
}

func CheckupDB(path string) {

	exists, err := doesDbFileExists(path)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		createDbFile(path)
	}

	// db, err := sql.Open("sqlite3", path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	db := GetDB()

	CategoryTableCheckup(db)
	TaskTableCheckup(db)

	defer db.Close()

}
