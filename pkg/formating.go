package pkg

import "fmt"

func MyPromptFormat() {
	fmt.Print("\033[31m" + "Ask me> " + "\033[0m")
}
