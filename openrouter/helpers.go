package openrouter

import (
	"encoding/json/v2"
	"log"

	t "github.com/Floris22/go-llm/internal/types"
)

func CreateRequestBody(
	messages []map[string]string,
	model string,
	temperature *float64,
	maxTokens *int,
	schema *map[string]any,
	tools *[]map[string]any,
	reasoning *map[string]any,
) []byte {
	maxTokensValue := 32000
	temperatureValue := 0.7

	if maxTokens != nil {
		maxTokensValue = *maxTokens
	}
	if temperature != nil {
		temperatureValue = *temperature
	}

	reqBody := t.OpenRouterRequest{
		Model:       model,
		Messages:    messages,
		Temperature: temperatureValue,
		MaxTokens:   maxTokensValue,
	}

	if schema != nil && tools != nil {
		log.Fatalf("Cannot do structured response and tool call repsponse at the same time")
	}
	if schema != nil {
		reqBody.ResponseFormat = *schema
	}
	if tools != nil {
		reqBody.Tools = *tools
		reqBody.ToolChoice = "required"
	}
	if reasoning != nil {
		reqBody.Reasoning = *reasoning
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return body
}
