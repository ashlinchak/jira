package lib

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func ExecuteCommand(command string, args []string) {
	cmd := exec.Command(command, args...)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := color.New(color.FgGreen)
	c.Print("=> create branch ")
	fmt.Print("\"" + args[2] + "\"")
	fmt.Println()
	c.Println("=> checkout to created branch")
}
