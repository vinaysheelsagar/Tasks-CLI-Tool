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
	Short: "Delete an existing task.",
	Long:  `This command allows you to delete an existing task.`,

	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		id, err := database.GetTaskID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid task name or ID.")
			os.Exit(0)
		}

		err = database.DeleteTask(id)
		utilities.CheckNil(err, "", "")
	},
}

func init() {}
