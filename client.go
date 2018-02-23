package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var httpClient = http.Client{}

type Client struct {
	config Config
}

func NewClient() *Client {
	client := new(Client)
	setConfig(client)

	return client
}

func (cl *Client) GetIssue(issueKey string) (*Issue, error) {
	uri := cl.config.IssueURI + "/" + issueKey
	req, err := http.NewRequest("GET", uri, nil)
	showError(err)
	req.SetBasicAuth(cl.config.User, cl.config.Password)
	resp, err := httpClient.Do(req)
	showError(err)

	issue := Issue{}

	if resp.StatusCode == 200 {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&issue)
		showError(err)
		return &issue, nil
	}

	fmt.Println(resp.Status)
	return nil, errors.New("Cannot find the issue with key " + issueKey)
}

func setConfig(cl *Client) {
	config := Config{}
	config.setDefaults()
	cl.config = config
}

func showError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
