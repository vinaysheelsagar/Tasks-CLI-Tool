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

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a category.",
	Long:  `This command deletes an existing category and all its associated tasks.`,

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

func init() {}
