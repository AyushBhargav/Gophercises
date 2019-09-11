package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of tool",
	Long:  `Released version of the tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GTD v0.1 --HEAD")
	},
}