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
	Run:     runBranchCmd,
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

func runBranchCmd(cmd *cobra.Command, args []string) {
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
		prefix = branchPrefix(config, issue.Fields.IssueType.Name)
	}

	if master == "" {
		master = config.Branch.Master
	}

	createBranch(*issue, config, prefix)
}

func branchPrefix(config lib.Config, issueTypeName string) string {
	prefix := ""

	for issueType, branchPrefix := range config.Branch.Prefixes {
		if issueTypeName == issueType {
			prefix = branchPrefix
			break
		}
	}

	return prefix
}

func createBranch(issue lib.Issue, config lib.Config, prefix string) {
	branchName := issue.Key + "_" + prepareBranchName(issue.Fields.Summary)
	if prefix != "" {
		branchName = prefix + "/" + branchName
	}

	git := lib.Git{}

	git.CreateBranch(branchName, config.Branch.Master)
}

func prepareBranchName(str string) string {
	name := strings.ToLower(str)

	// Accept only alphabet, numbers and underscore
	reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
	name = reg.ReplaceAllString(name, "_")

	// remove not valid symbols from head and tail of string
	reg = regexp.MustCompile("(^[^a-zA-Z0-9]+|[^a-zA-Z0-9]+$)")
	name = reg.ReplaceAllString(name, "")

	// remove duplicate underscores
	reg = regexp.MustCompile("_{2,}")
	name = reg.ReplaceAllString(name, "_")

	return name
}
