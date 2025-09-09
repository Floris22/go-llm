package main

import (
	"fmt"
	"os"

	"github.com/Floris22/go-llm/groq"
)

func main() {
	client := groq.NewClient(os.Getenv("GROQ_API_KEY"))
	audioURL := "https://pub-b6dd1cbdc818468198392598b6f15769.r2.dev/test.wav"
	resp, err := client.Transcribe("whisper-large-v3-turbo", "nl", &audioURL, nil, nil)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", resp.Text)
	fmt.Printf("Duration: %f\n", resp.Duration)
}

// func main() {
// 	client := openrouter.NewClient(os.Getenv("OPENROUTER_API_KEY"))
// 	image_url := ""
// 	messages := []map[string]any{
// 		{"role": "system", "content": "You are a helpful assistant."},
// 		// {"role": "user", "content": "Who won the world series in 2020?"},
// 		{
// 			"role": "user",
// 			"content": []map[string]any{
// 				{"type": "text", "text": "Describe the image in detail."},
// 				{"type": "image_url", "image_url": map[string]string{"url": image_url}},
// 			},
// 		},
// 	}
// 	reasoning := map[string]any{
// 		// "effort":     "low",
// 		"max_tokens": 0,
// 	}

// 	// inputTools := []map[string]any{
// 	// 	{
// 	// 		"type": "function",
// 	// 		"function": map[string]any{
// 	// 			"name":        "answer",
// 	// 			"description": "Answer the user's question",
// 	// 			"parameters": map[string]any{
// 	// 				"type": "object",
// 	// 				"properties": map[string]any{
// 	// 					"answer": map[string]any{"type": "string"},
// 	// 				},
// 	// 				"required": []string{"answer"},
// 	// 			},
// 	// 		},
// 	// 	},
// 	// }

// 	inputSchema := map[string]any{
// 		"type": "json_schema",
// 		"json_schema": map[string]any{
// 			"name":   "answer",
// 			"strict": true,
// 			"schema": map[string]any{
// 				"type": "object",
// 				"properties": map[string]any{
// 					"answer": map[string]any{
// 						"type":        "string",
// 						"description": "The answer to the user's question",
// 					},
// 				},
// 				"required":             []string{"answer"},
// 				"additionalProperties": false,
// 			},
// 		},
// 	}

// 	resp, err := client.GenerateStuctured(messages, inputSchema, "google/gemini-2.5-flash", nil, nil, nil, &reasoning)
// 	// resp, err := client.GenerateTools(messages, inputTools, "google/gemini-2.5-flash-lite", nil, nil, nil, &reasoning)
// 	// resp, err := client.GenerateText(messages, "google/gemini-2.5-flash-lite", nil, nil, nil, &reasoning)

// 	if err != nil {
// 		fmt.Printf("An error occured: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// fmt.Printf("Full response: %v\n", resp)
// 	fmt.Printf("Response text: %s\n", resp.Choices[0].Message.Content)
// 	// fmt.Printf("Response: %v\n", (*resp.Choices[0].Message.ToolCalls)[0].Function.Arguments)

// 	fmt.Printf("Usage:\n- CompletionTokens: %d\n- PromptTokens: %d\n", resp.Usage.CompletionTokens, resp.Usage.PromptTokens)
// }
