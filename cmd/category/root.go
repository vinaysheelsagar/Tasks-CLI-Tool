/*
Copyright © 2024 Vinay Sheel Sagar <vinaysheelsagar@gmail.com>
*/
package category

import (
	"github.com/spf13/cobra"
)

var CategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Manage task categories.",
	Long:  `This command allows you to manage task categories. You can create, list, and delete categories to organize your tasks effectively.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CategoryCmd.AddCommand(CreateCmd)
	CategoryCmd.AddCommand(UpdateCmd)
	CategoryCmd.AddCommand(DeleteCmd)
	CategoryCmd.AddCommand(ViewCmd)
}
