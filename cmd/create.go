package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "To create a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You called task: create command")
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
