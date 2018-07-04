# Jira Go plugin

### Installation
* Install Go language
* Clone this repository
* Build executable file by the command:

```
$ go build
```

* Change the `config.json` file according to your settings. Also, your Atlassian Jira password could be set via an environment variable (see below).
* Add compiled executable `jira` file to the `$PATH`

You can also use the executable from `bin` directory. Only needed to prepare config file.

### Environment variables:
You can specify the Atlassian Jira pass in the `config.json` or via `JIRA_PASS` environment variable.

### Usage

Run the command below for getting all available commands.
```
$ jira help
```
Currently implemented these commands:
* branch
* open


#### branch: create a branch from master based on the Jira issue summary

```
$ jira branch <issue-key>
```
Where `<issue-key>` is the key to the ticket in the Jira:
> `https://your-company.atlassian.net/browse/<issue-key>`

This will create a branch by this command:

```
git checkout -b feature/ISSUE-KEY_your_issue_summary_here master
```

You should implement a mapping for matching an issue key with a branch prefix in the `prefix` section of the `config.json` file.

Options:
* **-p, --prefix** - custom branch prefix
* **-m, --master** - master branch

#### open: open Atlassian Jira issue in a browser

```
$ jira open <issue-key>
```
This command is opening Jira issue in your default browser.