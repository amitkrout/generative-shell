package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

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
		time.Sleep(5 * time.Second)

		pkg.GenerateOpFormat(command)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		program, flag := pkg.ProgramAndSubcommand()

		if input == "y" || input == "yes" {
			cmd := exec.Command(program, flag, command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
		} else if input == "n" || input == "no" {
			fmt.Println("Ok, I won't run that command.")
		} else {
			fmt.Println("I didn't understand your input. Please enter y or n. Defaulting to no.")
		}
	}
}
