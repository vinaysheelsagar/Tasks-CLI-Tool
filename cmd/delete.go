package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You called task: delete command")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
