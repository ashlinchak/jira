package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ashlinchak/jira/lib"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:     "branch",
	Short:   "Create a GIT branch based on Atlassian Jira isssue key",
	Long:    `Create a GIT branch based on Atlassian Jira isssue key.`,
	Example: "jira branch lm-1112 --prefix=capability",
	Run:     run,
}

var prefix string
var master string

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Branch prefix")
	branchCmd.PersistentFlags().StringVarP(&master, "master", "m", "", "Master branch")
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

func run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("You should provide at least one argument")
		return
	}

	config := lib.Config{}

	config.SetDefaults()

	issueKey := args[0]

	client := lib.NewClient()
	issue, err := client.GetIssue(issueKey)

	if err != nil {
		fmt.Println(err)
		return
	}

	if prefix == "" {
		switch issue.Fields.IssueType.Name {
		case "Story":
			prefix = config.Git.Branches.Feature
		case "Bug":
			prefix = config.Git.Branches.Hotfix
		}
	}

	if master == "" {
		master = config.Git.Branches.Master
	}

	createBranch(*issue, config, prefix)
}

func createBranch(issue lib.Issue, config lib.Config, prefix string) {
	branchName := issue.Key + "_" + prepareBranchName(issue.Fields.Summary)
	if prefix != "" {
		branchName = prefix + "/" + branchName
	}

	git := lib.Git{}

	git.CreateBranch(branchName, config.Git.Branches.Master)
}

func prepareBranchName(str string) string {
	name := strings.ToLower(str)
	// Remove special symbols
	reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
	name = reg.ReplaceAllString(name, "_")
	// remove special symbols from the end of the string
	req, _ := regexp.Compile("[^a-zA-Z0-9]+$")
	return req.ReplaceAllString(name, "")
}
