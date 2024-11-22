/*
Copyright © 2024 Vinay Sheel Sagar <vinaysheelsagar@gmail.com>
*/
package category

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var CategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "To categorize tasks in a category",
	Long: `You can link tasks to other tasks so that in case parent task completes 
	then the sub tasks are already market complete or when tasks lead to subtasks`,

	Run: func(cmd *cobra.Command, args []string) {
		// toggle_value, _ := cmd.Flags().GetBool("toggle")
		name, _ := cmd.Flags().GetString("name")

		if name != "" {
			fmt.Println(name)
		}

		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	CategoryCmd.Flags().StringP("name", "n", "", "Name of the category")
	CategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
