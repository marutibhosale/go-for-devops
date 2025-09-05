package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command {
	Use:   "cpu",
	Short: "Get CPU usage statistics",
	Long:  `This command retrieves and displays the CPU usage statistics of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := GetCPUUsage()
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}