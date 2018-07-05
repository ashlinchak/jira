#!/usr/bin/env bash

# windows
GOOS=windows GOARCH=386 go build -o tmp/windows/jira.exe
cp config.json.example tmp/windows/config.json
zip -j bin/windows.zip tmp/windows/jira.exe tmp/windows/config.json

# linux
GOOS=linux GOARCH=386 go build -o tmp/linux/jira
cp config.json.example tmp/linux/config.json
zip -j bin/linux.zip tmp/linux/jira tmp/linux/config.json

# macos
GOOS=darwin GOARCH=386 go build -o tmp/darwin/jira
cp config.json.example tmp/darwin/config.json
zip -j bin/macos.zip tmp/darwin/jira tmp/darwin/config.json

echo "Done"
