package cmd

import (
	"ayushbhargav/task_manager/db"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do [task_ref_num]",
	Short: "Mark task as done",
	Long:  `Check task as done and remove from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, task := range args {
			db.MarkComplete(task)
			fmt.Printf("Task: '%s' has been done!\n", task)
		}
	},
}
