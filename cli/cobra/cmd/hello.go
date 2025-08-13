package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  `This command allows you to say hello to someone.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from my first cli app")
},
}

func init() {
	rootcmd.AddCommand(helloCmd)
}