package pkg

import (
	"fmt"
	"os"
)

func PreChecksAPIkey() bool {
	fmt.Println("Checking for API key...")
	return os.Getenv("OPENAI_API_KEY") != ""
}
