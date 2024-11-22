package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/cmd"
	db_controller "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
)

type Task struct {
	ID       int
	name     string
	desc     string
	category string
	status   string
	priority int
}

func main() {

	var DbPath string = "./tasks.db"

	db_controller.CheckupDB(DbPath)

	// CheckupDB(DbPath)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter task title: ")
	// title, _ := reader.ReadString('\n')

	// var task Task

	// task.task_title = title[:len(title)-1]
	// task.task_status = "Pending"

	// // createTask(db, task)

	cmd.Execute()
}
