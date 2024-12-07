package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "To create a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		var categoryID *int64

		priority, err := cmd.Flags().GetInt("priority")
		utilities.CheckNil(err, "", "")

		category, err := cmd.Flags().GetString("category")
		utilities.CheckNil(err, "", "")

		if category != "" {
			_categoryID, err := database.GetCategoryID(category)
			utilities.CheckNil(err, "", "")

			categoryID = &_categoryID
		}

		name := args[0]

		id, err := database.CreateTask(name, priority, categoryID)
		utilities.CheckNil(err, "Could not add task to records", "")

		fmt.Println(id)

	},
}

func init() {
	CreateCmd.Flags().IntP("priority", "p", 10, "sets priority for task")
	CreateCmd.Flags().StringP("category", "c", "", "sets category for task")
}
