package main

import (
	"io/fs"
	"os"
	"path"

	"github.com/vinaysheelsagar/Tasks-CLI-Tool/cmd"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

func main() {
	homeDir, err := os.UserHomeDir()
	utilities.CheckNil(err, "Could not find user's home dir", "")

	taclDir := path.Join(homeDir, ".tacl")
	os.MkdirAll(taclDir, fs.ModePerm)

	config := getConfig(taclDir)

	database.CheckupDB(config.DbPath)

	cmd.Execute()
}
