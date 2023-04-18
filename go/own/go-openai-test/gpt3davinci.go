package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient("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")

	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:  openai.GPT3Davinci,
			Prompt: "To be or not to be.",
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Text)
}
