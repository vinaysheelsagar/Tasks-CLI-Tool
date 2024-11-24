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

// categoryCmd represents the category command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "To create a new category",
	Long:  `you can use the category command to create a new category`,

	Run: func(cmd *cobra.Command, args []string) {
		priority, err := cmd.Flags().GetInt("priority")
		utilities.CheckNil(err, "", "")

		name := args[0]

		id, err := database.CreateCategory(name, priority)
		utilities.CheckNil(err, "Could not add category to records", "")

		fmt.Println(id)
	},
}

func init() {
	CreateCmd.Flags().IntP("priority", "p", 0, "category priority value")

}
