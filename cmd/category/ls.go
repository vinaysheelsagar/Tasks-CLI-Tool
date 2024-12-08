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

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all categories.",
	Long:  `This command displays a list of all existing categories.`,

	Run: func(cmd *cobra.Command, args []string) {

		categories, err := database.ListCategory()
		utilities.CheckNil(err, "", "")

		fmt.Println(categories)
	},
}

func init() {
	CategoryCmd.AddCommand(lsCmd)

}
