package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get Disk usage statistics",
	Long:  `This command retrieves and displays the disk usage statistics of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := GetDiskUsage()
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
}