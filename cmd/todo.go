package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	database "github.com/vinaysheelsagar/Tasks-CLI-Tool/db"
	"github.com/vinaysheelsagar/Tasks-CLI-Tool/utilities"
)

// TODO: filter by categories flag
var ShowCmd = &cobra.Command{
	Use:   "todo",
	Short: "List incomplete tasks by priority.",
	Long:  `This command displays a list of all incomplete tasks, sorted by their priority level. Higher priority tasks will appear at the top of the list.`,

	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDB()
		tasks, err := database.ShowUncheckedTasks(db)
		defer db.Close()

		utilities.CheckNil(err, "", "")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "ID\tTask\tCategory\n")

		for _, task := range tasks {
			categoryName := ""
			if task.CategoryID != nil {
				db := database.GetDB()
				category, err := database.ReadCategory(db, *task.CategoryID)
				defer db.Close()
				utilities.CheckNil(err, "", "")
				categoryName = category.Name
			}
			fmt.Fprintf(
				w,
				"%d\t%s\t%v\n",
				task.ID,
				task.Name,
				categoryName,
			)
		}

		w.Flush()

	},
}

func init() {}
