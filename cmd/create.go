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
		// get flags
		priority, _ := cmd.Flags().GetInt("priority")

		// get args
		task_name := args[0]

		// add task
		fmt.Println(task_name)
		fmt.Println(priority)

		// show task id

	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().IntP("priority", "p", 10, "sets priority for task")
}
