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

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a category.",
	Long:  `This command allows you to update the name of an existing category.`,

	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		db := database.GetDB()
		id, err := database.GetCategoryID(db, input)
		if err != nil {
			defer db.Close()
			fmt.Println("ERROR: not a valid category name")
			os.Exit(0)
		}

		category, err := database.ReadCategory(db, id)
		defer db.Close()
		utilities.CheckNil(err, "", "")

		updatePriority, err := cmd.Flags().GetInt("priority")
		utilities.CheckNil(err, "", "")
		if updatePriority != -1 {
			category.Priority = updatePriority
		}

		updateName, err := cmd.Flags().GetString("name")
		utilities.CheckNil(err, "", "")

		if updateName != "" {
			category.Name = updateName
		}

		db = database.GetDB()
		category, err = database.UpdateCategory(db, category)
		defer db.Close()
		utilities.CheckNil(err, "", "")

		fmt.Println(category)
	},
}

func init() {
	UpdateCmd.Flags().IntP("priority", "p", -1, "set priority")
	UpdateCmd.Flags().StringP("name", "n", "", "set name")
}
