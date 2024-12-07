package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete a task",
	Long:  `You can use the task command to create a new task`,

	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		id, err := database.GetTaskID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid task name")
			os.Exit(0)
		}

		err = database.DeleteTask(id)
		utilities.CheckNil(err, "", "")
	},
}

func init() {}
