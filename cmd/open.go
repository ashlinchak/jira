package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/ashlinchak/jira/lib"
	"github.com/spf13/cobra"
)

var openTicketCmd = &cobra.Command{
	Use:     "open",
	Short:   "Open Atlassian Jira isssue",
	Long:    `Open Atlassian Jira isssue by default browser`,
	Example: "jira open lm-1112",
	Run:     runOpenTicketCmd,
}

func init() {
	rootCmd.AddCommand(openTicketCmd)
	openTicketCmd.SetUsageTemplate(`Usage:
  jira open <issue-key> [flags]
{{ if .HasLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}
{{if .HasExample}}
Examples:
	{{.Example}}{{end}}
`)
}

func runOpenTicketCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("You should provide at least one argument")
		return
	}

	issueKey := args[0]
	client := lib.NewClient()
	uri := client.BrowseIssueURI(issueKey)
	openBrowser(uri)
}

func openBrowser(uri string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", uri).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", uri).Start()
	case "darwin":
		err = exec.Command("open", uri).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
