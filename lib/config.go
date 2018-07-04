package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config for Jira
type Config struct {
	Host            string `json:"host"`
	APIPath         string `json:"api_path"`
	IssuePath       string `json:"issue_path"`
	BrowseIssuePath string `json:"browse_path"`
	User            string `json:"user"`
	Branch          branch `json:"branch"`
	Password        string
	IssueURI        string
	BrowseIssueURI  string
}

type branch struct {
	Master   string            `json:"master"`
	Prefixes map[string]string `json:"prefixes"`
}

func (config *Config) SetDefaults() {
	file, err := os.Open(FilePath("config.json"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config.Password == "" {
		config.Password = os.Getenv("JIRA_PASS")
	}
	config.IssueURI = config.Host + config.APIPath + config.IssuePath
	config.BrowseIssueURI = config.Host + config.BrowseIssuePath
}
