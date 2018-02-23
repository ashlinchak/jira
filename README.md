# Jira Go plugin

### Environment variables:
You can specify a Jira pass in the `config.json` or via `JIRA_PASS` environment variable.

### Examples

#### Create branch from master basing on the jira issue summary

`jira branch issue-key`

This will create a branch by this command:

`git checkout -b feature/ISSUE-KEY_your_issue_summary_here master`

Currently implemented two types of prefix:
* Story - will create a branch with the prefix **feature**
* Bug - will create a branch with the prefix **bugfix**

### Options
* **--prefix** - custom branch prefix 

### TODO
* Remove not alphabet symbols at the end of the branch name