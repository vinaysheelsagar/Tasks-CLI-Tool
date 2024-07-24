package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "To show a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You called task: read command")
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
}
