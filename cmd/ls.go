/*
Copyright © 2024 Vinay Sheel Sagar <vinaysheelsagar@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all tasks.",
	Long:  `This command displays a list of all tasks, including their status (completed or incomplete).`,

	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := database.ListTask()
		utilities.CheckNil(err, "", "")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "ID\tName\tPriority\tCategory\tStatus\n")
		for _, task := range tasks {
			categoryName := ""
			if task.CategoryID != nil {
				category, err := database.ReadCategory(*task.CategoryID)
				utilities.CheckNil(err, "", "")
				categoryName = category.Name
			}
			fmt.Fprintf(
				w,
				"%d\t%s\t%d\t%v\t%v\n",
				task.ID,
				task.Name,
				task.Priority,
				categoryName,
				task.Status,
			)
		}

		w.Flush()

	},
}

func init() {}
