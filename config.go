package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config for Jira
type Config struct {
	Host      string `json:"host"`
	APIPath   string `json:"api_path"`
	IssuePath string `json:"issue_path"`
	User      string `json:"user"`
	Git       git    `json:"git"`
	Password  string
	IssueURI  string
}

type git struct {
	MasterBranch string `json:"master_branch"`
}

func (config *Config) setDefaults() {
	file, _ := os.Open(FilePath("config.json"))
	decoder := json.NewDecoder(file)
	err := decoder.Decode(config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config.Password == "" {
		config.Password = os.Getenv("JIRA_PASS")
	}

	config.IssueURI = config.Host + config.APIPath + config.IssuePath
}
