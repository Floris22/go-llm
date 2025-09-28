package main

import (
	"encoding/json/v2"
	"fmt"
	"os"

	"github.com/Floris22/go-llm/v2/clients"
	"github.com/Floris22/go-llm/v2/llmtypes"
)

// Helper functions to create pointers to primitive types
func pts(s string) *string   { return &s }
func pti(i int) *int         { return &i }
func ptf(f float64) *float64 { return &f }

// ------------------------------
// Function calling example
// ------------------------------
func provideWeatherDetails(city string, country string) string {
	return fmt.Sprintf("The weather in %s, %s is sunny. The temperature is 25 degrees Celsius.", city, country)
}

var provideWeatherDetailsSchema = llmtypes.ToolSchema{
	Name:        "get_weather_details",
	Description: "Retrieves the weather details of the city and country",
	Parameters: llmtypes.ToolParameters{
		Type: "object",
		Properties: map[string]llmtypes.ParameterProperty{
			"city": {
				Type:        "string",
				Description: "The city e.g. Brussels",
			},
			"country": {
				Type:        "string",
				Description: "The country e.g. Belgium",
			},
		},
		Required: []string{"city", "country"},
	},
}

func main() {
	client := clients.NewOpenRouterClient(os.Getenv("OPENROUTER_API_KEY"))

	var messages []llmtypes.MessageForLLM

	systemMessage := llmtypes.MessageForLLM{
		Role:    llmtypes.RoleSystem,
		Content: pts("You are a helpful assistant."),
	}
	messages = append(messages, systemMessage)

	userMessage := llmtypes.MessageForLLM{
		Role:    llmtypes.RoleUser,
		Content: pts("What is the weather like today in Belgium?"),
	}
	messages = append(messages, userMessage)

	resp, err := client.GenerateTools(
		messages,
		nil,
		[]llmtypes.ToolSchema{provideWeatherDetailsSchema},
		"google/gemini-2.5-flash-lite-preview-09-2025",
		ptf(0.1),
		pti(100),
		pti(3),
		&llmtypes.ReasoningConfig{
			MaxTokens: pti(0),
		},
		nil,
	)
	if err != nil {
		panic(err)
	}

	tc := resp.Choices[0].Message.ToolCalls[0]
	fmt.Printf("id: %s\ntype: %s\nname: %s\nargs: %s\n", tc.ID, tc.Type, tc.Function.Name, tc.Function.Arguments)

	var toolArgsLoaded map[string]any
	err = json.Unmarshal([]byte(tc.Function.Arguments), &toolArgsLoaded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("toolArgs: %s\n", toolArgsLoaded)
	fmt.Printf("toolArgs.city: %s\n", toolArgsLoaded["city"])
}

// ------------------------------
// Structured output example
// ------------------------------
var structuredOutputExample = llmtypes.StructuredOutputSchema{
	Name:   "weather_details",
	Strict: true,
	Schema: llmtypes.StructuredOutputSchemaDefinition{
		Type: "object",
		Properties: map[string]llmtypes.StructuredOutputProperty{
			"city": {
				Type:        "string",
				Description: "The city e.g. Brussels",
			},
			"country": {
				Type:        "string",
				Description: "The country e.g. Belgium",
			},
			"weather": {
				Type:        "string",
				Description: "The weather e.g. sunny",
			},
			"temperature": {
				Type:        "number",
				Description: "The temperature in Celsius",
			},
		},
		Required: []string{"city", "country", "weather"},
	},
}

// func main() {
// 	client := clients.NewOpenRouterClient(os.Getenv("OPENROUTER_API_KEY"))

// 	var messages []llmtypes.MessageForLLM

// 	systemMessage := llmtypes.MessageForLLM{
// 		Role:    llmtypes.RoleSystem,
// 		Content: "You are a helpful assistant. The weather in Belgium is always rainy, and 12 degrees Celsius.",
// 	}
// 	messages = append(messages, systemMessage)

// 	userMessage := llmtypes.MessageForLLM{
// 		Role:    llmtypes.RoleUser,
// 		Content: "What is the weather like today in Belgium?",
// 	}
// 	messages = append(messages, userMessage)

// 	resp, err := client.GenerateStructured(
// 		messages,
// 		structuredOutputExample,
// 		"google/gemini-2.5-flash-lite",
// 		ptf(0.1),
// 		pti(100),
// 		pti(3),
// 		&llmtypes.ReasoningConfig{
// 			MaxTokens: pti(0),
// 		},
// 		nil,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp.Choices[0].Message.Content)
// 	// Will propably print something like:
// 	// {
// 	//   "city": "Brussels",
// 	//   "country": "Belgium",
// 	//   "weather": "sunny",
// 	//   "temperature": 12
// 	// }
// }
