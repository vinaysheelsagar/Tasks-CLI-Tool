package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "To edit a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You called task: update command")
	},
}

func init() {
	RootCmd.AddCommand(editCmd)
}
