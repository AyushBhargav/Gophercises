package cmd

import (
	"ayushbhargav/task_manager/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [task to be added]",
	Short: "Add new task",
	Long:  `Add new task to database`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		db.CreateNewTask(task)
		fmt.Printf("Task added: %s\n", task)
	},
}
