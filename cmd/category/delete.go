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
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To create a new category",
	Long:  `you can use the category command to create a new category`,

	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		id, err := database.GetCategoryID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid category name")
			os.Exit(0)
		}

		err = database.DeleteCategory(id)
		utilities.CheckNil(err, "", "")

	},
}

func init() {
	CategoryCmd.AddCommand(DeleteCmd)
}
