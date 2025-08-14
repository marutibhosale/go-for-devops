package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var printcmd = &cobra.Command{
	Use:   "print",
	Short: "Print a message",
	Long:  `This command prints a message to the console.`,
	Args: cobra.MinimumNArgs(1), // Ensure at least one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print " + strings.Join(args, " "))
},
}

func init() {
	rootcmd.AddCommand(printcmd)
}