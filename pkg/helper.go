package pkg

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func KeyboardInput() string {
	reader := bufio.NewReader(os.Stdin)
	request, _ := reader.ReadString('\n')
	return strings.TrimSpace(request)
}

func PlatformAndTerm() (string, string) {
	var osType string = runtime.GOOS

	if osType == "linux" {
		return "linux", "bash"
	} else if osType == "darwin" {
		return "mac", "bash"
	} else if osType == "windows" {
		var term string = os.Getenv("TERM")
		if term == "cygwin" {
			return "windows", "bash"
		} else if term == "cmd" {
			return "windows", "cmd"
		} else if term == "powershell" {
			return "windows", "powershell"
		} else {
			return "", ""
		}
	} else {
		return "", ""
	}
}

func PromptRunsOnShellType() (string, string) {
	platform, terminal := PlatformAndTerm()
	switch {
	case platform == "linux" && terminal == "bash":
		return PromptShell, TemplateShell
	case platform == "mac" && terminal == "bash":
		return PromptShell, TemplateShell
	case platform == "windows" && terminal == "bash":
		return PromptShell, TemplateShell
	case platform == "windows" && terminal == "cmd":
		return PromptCMD, TemplateCMD
	case platform == "windows" && terminal == "powershell":
		return PromptPS, TemplatePS
	default:
		return "", ""
	}
}

func Prompt() string {
	prompt, _ := PromptRunsOnShellType()
	if prompt == "" {
		err := fmt.Errorf("error: could not determine shell type")
		panic(err)
	}
	return prompt
}

func Template() string {
	_, template := PromptRunsOnShellType()
	if template == "" {
		err := fmt.Errorf("error: could not determine shell type")
		panic(err)
	}
	return template
}
