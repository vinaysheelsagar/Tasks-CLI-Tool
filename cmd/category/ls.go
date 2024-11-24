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
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "To create a new category",
	Long:  `you can use the category command to create a new category`,

	Run: func(cmd *cobra.Command, args []string) {

		categories, err := database.ListCategory()
		utilities.CheckNil(err, "", "")

		fmt.Println(categories)
	},
}

func init() {
	CategoryCmd.AddCommand(lsCmd)

}
