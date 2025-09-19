package helpers

import (
	"encoding/json/v2"
	"log"

	t "github.com/Floris22/go-llm/llmtypes"
)

func CreateRequestBody(
	messages []t.MessageForLLM,
	model string,
	temperature *float64,
	maxTokens *int,
	schema *t.StructuredOutputSchema,
	tools *[]t.ToolSchema,
	reasoning *t.ReasoningConfig,
	provider *t.ProviderConfig,
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
		Temperature: &temperatureValue,
		MaxTokens:   &maxTokensValue,
	}

	if schema != nil && tools != nil {
		log.Fatalf("Cannot do structured response and tool call repsponse at the same time")
	}

	if schema != nil {
		fullSchema := map[string]any{
			"type": "json_schema",
			"json_schema": map[string]any{
				"name":   schema.Name,
				"strict": schema.Strict,
				"schema": map[string]any{
					"type":                 schema.Schema.Type,
					"properties":           schema.Schema.Properties,
					"required":             schema.Schema.Required,
					"additionalProperties": schema.Schema.AdditionalProperties,
				},
			},
		}
		reqBody.ResponseFormat = &fullSchema
	}

	if tools != nil {
		var toolsFull []map[string]any
		for _, tool := range *tools {
			toolDef := map[string]any{
				"type": "function",
				"function": map[string]any{
					"name":        tool.Name,
					"description": tool.Description,
					"parameters":  tool.Parameters,
				},
			}
			toolsFull = append(toolsFull, toolDef)
		}

		reqBody.Tools = &toolsFull

		tc := t.ToolChoiceRequired
		reqBody.ToolChoice = &tc
	}

	if reasoning != nil {
		reqBody.Reasoning = reasoning
	}

	if provider != nil {
		reqBody.Provider = provider
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return body
}
