package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "my test cli",
	Short: "A simple CLI application",
	Long:  `This is a simple CLI application to demonstrate the use of Cobra for building command`,
}

func Execute() {
	if err:= rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	
}