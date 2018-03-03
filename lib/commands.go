package lib

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func ExecuteCommand(command string, args []string) {

	cmd := exec.Command(command, args...)

	c := color.New(color.FgGreen)
	c.Println("run: ", strings.Join(cmd.Args, " "))

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
