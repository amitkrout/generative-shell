package pkg

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func MyPromptFormat() {
	c := color.New(color.FgRed, color.Bold)
	c.Print("Ask me> ")
}

func GenerateOpFormat(command string) {
	runOutput := color.New(color.BgHiMagenta, color.Bold).Sprintf(">>> Run:")
	commandOutput := color.New(color.FgGreen, color.Bold).Sprintf(" %s", strings.TrimSpace(command))
	confirmOutput := color.New(color.FgMagenta, color.Bold).Sprintf(" (y/n) ")
	fmt.Print(runOutput, commandOutput, confirmOutput)
}
