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

// linkCmd represents the link command
var IncompleteCmd = &cobra.Command{
	Use:   "incomplete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
