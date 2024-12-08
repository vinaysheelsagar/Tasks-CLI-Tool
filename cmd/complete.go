/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var CompleteCmd = &cobra.Command{
	Use:   "check",
	Short: "Mark a task as complete.",
	Long:  `This command allows you to mark an existing task as complete.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		id, err := database.GetTaskID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid task name.")
			os.Exit(0)
		}

		err = database.CompleteTask(id)
		utilities.CheckNil(err, "", "")
	},
}

func init() {}
