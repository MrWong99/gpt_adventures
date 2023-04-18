package main

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	token, err := readToken()
	if err != nil {
		log.Fatal(err)
	}

	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "How did I make this request?",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, choice := range resp.Choices {
		log.Printf("Choice %d | FinishReason: %s | Message.Name: %s | Message.Role: %s | Message.Content: %s\n",
			i, choice.FinishReason, choice.Message.Name, choice.Message.Role, choice.Message.Content)
	}
}

func readToken() (token string, err error) {
	var ok bool
	if token, ok = os.LookupEnv("API_TOKEN"); ok {
		return
	}
	var tokenBytes []byte
	tokenBytes, err = os.ReadFile("API_TOKEN")
	if tokenBytes != nil {
		token = string(tokenBytes)
	}
	return
}
