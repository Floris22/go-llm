package main

import (
	"os"

	"github.com/Floris22/go-llm/openrouter"
)

func main() {
	client := openrouter.NewClient(os.Getenv("OPENROUTER_API_KEY"))

	resp, err := client.GenerateText(
		[]map[string]any{
			{"role": "system", "content": "You are a helpful assistant. Answer very short and concise."},
			{"role": "user", "content": "Is a banana more yellow than a yellow kiwi?"},
		},
		"google/gemini-2.5-flash-lite",
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	println(resp.Choices[0].Message.Content)
}
