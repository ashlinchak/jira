package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func FilePath(file string) string {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	return exPath + "/" + file
}

func ExecuteCommand(command string, args []string) {

	cmd := exec.Command(command, args...)

	c := color.New(color.FgGreen)
	c.Println("run: ", strings.Join(cmd.Args, " "))

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
