package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/generative-shell/pkg"
)

func main() {

	if !pkg.PreChecksAPIkey() {
		fmt.Println("Please set your API key as an environment variable named OPENAI_API_KEY")
		return
	}

	for {

		pkg.MyPromptFormat()

		resp, err := pkg.Completion()
		if err != nil {
			fmt.Println(err)
			return
		}

		command := resp.Choices[0].Text

		confirm := false
		promptRun := fmt.Sprintf(">>> Run: \033[34m%s\033[0m", command)
		if err = survey.AskOne(&survey.Confirm{Message: promptRun, Default: confirm}, &confirm); err != nil {
			fmt.Println(err)
			return
		}

		if confirm {
			cmd := exec.Command("bash", "-c", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}
