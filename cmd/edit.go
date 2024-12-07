package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "To edit a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		id, err := database.GetTaskID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid task name")
			os.Exit(0)
		}

		task, err := database.ReadTask(id)
		utilities.CheckNil(err, "", "")

		updatePriority, err := cmd.Flags().GetInt("priority")
		utilities.CheckNil(err, "", "")
		if updatePriority != -1 {
			task.Priority = updatePriority
		}

		updateName, err := cmd.Flags().GetString("name")
		utilities.CheckNil(err, "", "")

		if updateName != "" {
			task.Name = updateName
		}

		category, err := cmd.Flags().GetString("category")
		utilities.CheckNil(err, "", "")

		if category != "" {
			categoryID, err := database.GetCategoryID(category)
			utilities.CheckNil(err, "", "")

			task.CategoryID = &categoryID
		}

		task, err = database.UpdateTask(task)
		utilities.CheckNil(err, "", "")

		fmt.Println(task)
	},
}

func init() {
	EditCmd.Flags().IntP("priority", "p", -1, "set priority")
	EditCmd.Flags().StringP("name", "n", "", "set name")
	EditCmd.Flags().StringP("category", "c", "", "set category")
}
