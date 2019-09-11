package cmd

import (
	"ayushbhargav/task_manager/db"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists incomplete tasks",
	Long:  `Retrieves and lists incomplete tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.GetIncompleteTasks()
		fmt.Println("You have the following tasks:")
		for id, task := range tasks {
			fmt.Printf("%s. %s", id, task)
			fmt.Println()
		}
	},
}
