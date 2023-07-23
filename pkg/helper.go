package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/generative-shell/trainingdata"
)

func KeyboardInput() string {
	reader := bufio.NewReader(os.Stdin)
	request, _ := reader.ReadString('\n')
	return strings.TrimSpace(request)
}

func DetectTerminalType() string {
	var shellType string = os.Getenv("SHELL")

	if strings.Contains(strings.TrimSpace(shellType), "bash") {
		return "bash"
	} else if strings.Contains(strings.TrimSpace(shellType), "zsh") {
		return "zsh"
	} else if strings.Contains(strings.TrimSpace(shellType), "sh") {
		return "sh"
	} else if strings.Contains(strings.TrimSpace(shellType), "CMD") {
		return "cmd"
	} else if strings.Contains(strings.TrimSpace(shellType), "PS") {
		return "powershell"
	} else {
		return ""
	}
}

func PromptRunsOnShellType() (string, string) {
	terminalName := DetectTerminalType()
	switch {
	case terminalName == "bash":
		return trainingdata.PromptShell, trainingdata.TemplateShell
	case terminalName == "zsh":
		return trainingdata.PromptShell, trainingdata.TemplateShell
	case terminalName == "sh":
		return trainingdata.PromptShell, trainingdata.TemplateShell
	case terminalName == "cmd":
		return trainingdata.PromptCMD, trainingdata.TemplateCMD
	case terminalName == "powershell":
		return trainingdata.PromptPS, trainingdata.TemplatePS
	default:
		return "", ""
	}
}

func Prompt() string {
	prompt, _ := PromptRunsOnShellType()
	if prompt == "" {
		err := fmt.Errorf("error: unsupported shell type")
		panic(err)
	}
	return prompt + fmt.Sprintf(Template(), KeyboardInput())
}

func Template() string {
	_, template := PromptRunsOnShellType()
	if template == "" {
		err := fmt.Errorf("error: unsupported shell type")
		panic(err)
	}
	return template
}

func ProgramAndSubcommand() (string, string) {
	terminalName := DetectTerminalType()
	switch {
	case terminalName == "bash":
		return "bash", "-c"
	case terminalName == "zsh":
		return "zsh", "-c"
	case terminalName == "sh":
		return "sh", "-c"
	case terminalName == "cmd":
		return "cmd", "/C"
	case terminalName == "powershell":
		return "powershell", "-Command"
	default:
		return "", ""
	}
}
