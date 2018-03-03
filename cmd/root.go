package cmd

import (
	"fmt"
	"os"

	"github.com/ashlinchak/jira/lib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "jira",
	Short:   "Jira is CLI for working with Atlassian Jira",
	Version: lib.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to specify a command")
		fmt.Println("Run `jira help` for getting more information")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
