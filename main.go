package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var prefix string
var client = NewClient()
var config = Config{}

func main() {
	// Set default configs
	config.setDefaults()

	branchCmd := flag.NewFlagSet("branch", flag.ExitOnError)
	branchCmd.StringVar(&prefix, "prefix", "", "Branch prefix")

	if len(os.Args) < 3 {
		fmt.Println("usage: jira <command> [<args>]")
		showValidCommands()
		return
	}

	switch os.Args[1] {
	case "branch":
		branchCmd.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		showValidCommands()
		return
	}

	// Get Issue
	issueKey := os.Args[len(os.Args)-1]
	issue, err := client.GetIssue(issueKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	if prefix == "" {
		switch issue.Fields.IssueType.Name {
		case "Story":
			prefix = "feature"
		case "Bug":
			prefix = "bugfix"
		}
	}

	// create git branch
	createBranch(*issue, config, prefix)
}

func createBranch(issue Issue, config Config, prefix string) {
	branchName := issue.Key + "_" + prepareBranchName(issue.Fields.Summary)
	if prefix != "" {
		branchName = prefix + "/" + branchName
	}

	git := Git{}

	git.CreateBranch(branchName, config.Git.MasterBranch)
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

func showValidCommands() {
	fmt.Println("Valid commands are:")
	fmt.Println("* jira branch lm-1111 <-prefix=branch_prefix>")
}
