package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var ReadCmd = &cobra.Command{
	Use:   "read",
	Short: "To show a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		id, err := database.GetTaskID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid task name.")
			os.Exit(0)
		}

		task, err := database.ReadTask(id)
		utilities.CheckNil(err, "", "")

		println(task.ID, task.Name, task.Priority, task.CategoryID)
		// if task.CategoryID != nil {
		// 	fmt.Println(*task.CategoryID)
		// }
	},
}

func init() {}
