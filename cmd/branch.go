package cmd

import (
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a GIT branch based on Atlassian Jira isssue key",
	Long: `Create a GIT branch based on Atlassian Jira isssue key.`,
	Example: "jira branch lm-1112 --prefix=capability",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO:
	},
}


func init() {
	rootCmd.AddCommand(branchCmd)

	branchCmd.PersistentFlags().String("prefix", "", "Branch prefix")
	branchCmd.SetUsageTemplate(`Usage:
  jira branch <issue-key> [flags]
{{ if .HasLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}
{{if .HasExample}}
Examples:
	{{.Example}}{{end}}
`)
}	
