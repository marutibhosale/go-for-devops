package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

func GetCPUUsage() string {
	pecentages, err := cpu.Percent(0, false)
	if err != nil {
		return fmt.Sprintf("Error getting CPU usage: %s", err)
	}
	return fmt.Sprintf("CPU Usage: %.2f%%", pecentages[0])
}

func GetMemoryUsage() string {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Sprintf("Error getting memory usage: %s", err)
	}
	return fmt.Sprintf("Memory Usage: Total: %v, Used: %v, Free: %v, Used Percent: %.2f%%",
		vmStat.Total, vmStat.Used, vmStat.Free, vmStat.UsedPercent)
}

func GetDiskUsage() string {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return fmt.Sprintf("Error getting Disk Usage: %s", err)
	}
	return fmt.Sprintf("Disk Usage: Total: %v, Free: %v, UsedPercent: %.2f%%", diskStat.Total, diskStat.Free, diskStat.UsedPercent)
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get system statistics",
	Long:  `This command retrieves and displays various system statistics such as CPU usage, memory usage, and disk usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		cpuUsage := GetCPUUsage()
		memUsage := GetMemoryUsage()
		diskUsage := GetDiskUsage()

		results := strings.Join([]string{cpuUsage,memUsage,diskUsage,}, "\n")
		fmt.Println(results)
	},
}	

func init() {
	rootCmd.AddCommand(statsCmd)
}