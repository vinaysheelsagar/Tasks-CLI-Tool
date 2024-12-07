package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/cmd/category"
)

var RootCmd = &cobra.Command{
	Use:     "tacl",
	Version: "0.1.0",
	Short:   "A command-line interface for manage tasks",
	Long:    "A simple CLI tool for generating tasks using the github.com/vinaysheelsagar/Tasks-CLI-Tool library.",

	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Println(cmd.Version)
			os.Exit(0)
		}

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
	RootCmd.AddCommand(CreateCmd)
	RootCmd.AddCommand(ReadCmd)
	RootCmd.AddCommand(EditCmd)
	RootCmd.AddCommand(DeleteCmd)
	RootCmd.AddCommand(LsCmd)
	RootCmd.AddCommand(ShowCmd)
	RootCmd.AddCommand(IncompleteCmd)
	RootCmd.AddCommand(CompleteCmd)

	RootCmd.Flags().BoolP("version", "v", false, "Prints the version number")
}
