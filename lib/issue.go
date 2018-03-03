package lib

// Issue model in Jira
type Issue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields field  `json:"fields"`
}

type field struct {
	Summary   string    `json:"summary"`
	IssueType issuetype `json:"issueType"`
}

type issuetype struct {
	Name string `json:"name"`
}
