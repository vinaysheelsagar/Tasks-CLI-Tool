/*
Copyright © 2024 Vinay Sheel Sagar <vinaysheelsagar@gmail.com>
*/
package category

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var updateCmd = &cobra.Command{
	Use:   "create",
	Short: "To create a new category",
	Long:  `you can use the category command to create a new category`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You called category: create command")
	},
}

func init() {
	CategoryCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().StringP("name", "n", "", "Name of the category")
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
