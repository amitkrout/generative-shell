package pkg

import (
	"fmt"

	"github.com/fatih/color"
)

func MyPromptFormat() {
	c := color.New(color.FgRed, color.Bold)
	c.Print("Ask me> ")
}

func GenerateOpFormat(command string) {
	output := color.New(color.FgGreen, color.Bold).Sprintf(">>> Run: %s", command)
	fmt.Print(output + " (y/n) ")
}
