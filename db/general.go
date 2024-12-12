package database

import (
	"database/sql"

	configuration "github.com/vinaysheelsagar/Tasks-CLI-Tool/config"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

func GetDB() *sql.DB {
	config := configuration.GetConfig()

	db, err := sql.Open("sqlite3", config.DbPath)
	utilities.CheckNil(err, "", "")

	return db
}
