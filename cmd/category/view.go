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

var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View category details.",
	Long:  `This command displays detailed information about a specific category, including its name and associated tasks.`,

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

		fmt.Println(category)
	},
}

func init() {}
