package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var EditCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task.",
	Long:  `This command allows you to update the details of an existing task. You can modify the task's description, dependencies, or other attributes.`,

	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		db := database.GetDB()
		id, err := database.GetTaskID(db, input)

		if err != nil {
			defer db.Close()
			fmt.Println("ERROR: not a valid task name")
			os.Exit(0)
		}

		task, err := database.ReadTask(db, id)
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
			categoryID, err := database.GetCategoryID(db, category)
			utilities.CheckNil(err, "", "")

			task.CategoryID = &categoryID
		}

		task, err = database.UpdateTask(db, task)
		utilities.CheckNil(err, "", "")

		defer db.Close()
		fmt.Println(task)
	},
}

func init() {
	EditCmd.Flags().IntP("priority", "p", -1, "set priority")
	EditCmd.Flags().StringP("name", "n", "", "set name")
	EditCmd.Flags().StringP("category", "c", "", "set category")
}
