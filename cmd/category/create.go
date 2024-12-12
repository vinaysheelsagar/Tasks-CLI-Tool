/*
Copyright © 2024 Vinay Sheel Sagar <vinaysheelsagar@gmail.com>
*/
package category

import (
	"fmt"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new category.",
	Long:  `This command allows you to create a new category to organize your tasks.`,

	Run: func(cmd *cobra.Command, args []string) {
		priority, err := cmd.Flags().GetInt("priority")
		utilities.CheckNil(err, "", "")

		name := args[0]

		db := database.GetDB()
		id, err := database.CreateCategory(db, name, priority)
		defer db.Close()
		utilities.CheckNil(err, "Could not add category to records", "")

		fmt.Println(id)
	},
}

func init() {
	CreateCmd.Flags().IntP("priority", "p", 5, "category priority value")
}
