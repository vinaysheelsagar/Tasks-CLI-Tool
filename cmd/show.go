package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "To show a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.ShowUncheckedTasks()
		utilities.CheckNil(err, "", "")

		fmt.Println(tasks)
	},
}

func init() {}
