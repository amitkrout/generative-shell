package pkg

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func CompletionRequestBody() openai.CompletionRequest {
	return openai.CompletionRequest{
		Model:       Model,
		Prompt:      Prompt(),
		MaxTokens:   MaxTokens,
		Stop:        []string{"\n"},
		Temperature: 0.0,
	}
}

func Completion() (openai.CompletionResponse, error) {
	ctx := context.Background()
	req := CompletionRequestBody()
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return client.CreateCompletion(ctx, req)
}
