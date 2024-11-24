package main

import (
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/cmd"
	configuration "github.com/vinaysheelsagar/Tasks-CLI-Tool/config"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
)

func main() {

	config := configuration.GetConfig()

	database.CheckupDB(config.DbPath)

	cmd.Execute()
}
