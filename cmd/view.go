package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var ReadCmd = &cobra.Command{
	Use:   "view",
	Short: "View detailed information about a specific task.",
	Long:  `This command allows you to display detailed information about a particular task. It provides a comprehensive overview of the task, including its description, dependencies, and other relevant details.`,

	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		db := database.GetDB()
		id, err := database.GetTaskID(db, input)
		if err != nil {
			defer db.Close()
			fmt.Println("ERROR: not a valid task name.")
			os.Exit(0)
		}

		task, err := database.ReadTask(db, id)
		defer db.Close()

		utilities.CheckNil(err, "", "")

		println(task.ID, task.Name, task.Priority, task.CategoryID)
	},
}

func init() {}
