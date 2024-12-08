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

var IncompleteCmd = &cobra.Command{
	Use:   "uncheck",
	Short: "Mark a completed task as incomplete.",
	Long:  `This command allows you to mark a completed task as incomplete.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		id, err := database.GetTaskID(input)
		if err != nil {
			fmt.Println("ERROR: not a valid task name.")
			os.Exit(0)
		}

		err = database.IncompleteTask(id)
		utilities.CheckNil(err, "", "")
	},
}

func init() {}
