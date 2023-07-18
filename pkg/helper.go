package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		return PromptShell, TemplateShell
	case terminalName == "zsh":
		return PromptShell, TemplateShell
	case terminalName == "sh":
		return PromptShell, TemplateShell
	case terminalName == "cmd":
		return PromptCMD, TemplateCMD
	case terminalName == "powershell":
		return PromptPS, TemplatePS
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
