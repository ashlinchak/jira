#!/usr/bin/env bash

# windows
GOOS=windows GOARCH=386 go build -o tmp/windows/jira.exe
zip -j bin/jira.zip tmp/windows/jira.exe config.json

# linux
GOOS=linux GOARCH=386 go build -o tmp/linux/jira
zip -j bin/linux.zip tmp/linux/jira config.json

# macos
GOOS=darwin GOARCH=386 go build -o tmp/darwin/jira
zip -j bin/macos.zip tmp/darwin/jira config.json

echo "Done"
