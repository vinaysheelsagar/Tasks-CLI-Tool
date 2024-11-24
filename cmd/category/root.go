/*
Copyright © 2024 Vinay Sheel Sagar <vinaysheelsagar@gmail.com>
*/
package category

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

// categoryCmd represents the category command
var CategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "To categorize tasks in a category",
	Long: `You can link tasks to other tasks so that in case parent task completes 
	then the sub tasks are already market complete or when tasks lead to subtasks`,

	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		id, err := database.GetCategoryID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid category name")
			os.Exit(0)
		}

		category, err := database.ReadCategory(id)
		utilities.CheckNil(err, "", "")

		fmt.Println(category)

	},
}

func init() {
	CategoryCmd.AddCommand(CreateCmd)
	CategoryCmd.AddCommand(UpdateCmd)
}
