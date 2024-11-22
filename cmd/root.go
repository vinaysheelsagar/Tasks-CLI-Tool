package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/cmd/category"
)

var RootCmd = &cobra.Command{
	Use:   "tacl",
	Short: "A command-line interface for manage tasks",
	Long:  "A simple CLI tool for generating tasks using the github.com/vinaysheelsagar/Tasks-CLI-Tool library.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(category.CategoryCmd)
}
