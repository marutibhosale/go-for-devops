package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var memCmd = &cobra.Command {
	Use:   "mem",
	Short: "Get Memory usage statistics",
	Long:  `This command retrieves and displays the memory usage statistics of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := GetMemoryUsage()
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(memCmd)
}