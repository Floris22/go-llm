package helpers

import (
	"encoding/json/v2"
	"fmt"
	"log"

	t "github.com/Floris22/go-llm/v2/llmtypes"
)

func CreateRequestBody(
	messages []t.MessageForLLM,
	messageParts []t.PartMessageForLLM,
	model string,
	temperature *float64,
	maxTokens *int,
	schema *t.StructuredOutputSchema,
	tools *[]t.ToolSchema,
	reasoning *t.ReasoningConfig,
	provider *t.ProviderConfig,
) ([]byte, error) {
	maxTokensValue := 32000
	temperatureValue := 0.7

	if messageParts != nil && messages != nil {
		return nil, fmt.Errorf("Cannot send both message and message parts")
	}
	if messageParts == nil && messages == nil {
		return nil, fmt.Errorf("Must send either message or message parts")
	}

	if maxTokens != nil {
		maxTokensValue = *maxTokens
	}

	if temperature != nil {
		temperatureValue = *temperature
	}

	reqBody := map[string]any{
		"model":       model,
		"temperature": temperatureValue,
		"max_tokens":  maxTokensValue,
	}

	if messages != nil {
		reqBody["messages"] = messages
	} else {
		reqBody["messages"] = messageParts
	}

	if schema != nil && tools != nil {
		return nil, fmt.Errorf("Cannot do structured response and tool call repsponse at the same time")
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
		reqBody["response_format"] = &fullSchema
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

		reqBody["tools"] = &toolsFull

		tc := t.ToolChoiceRequired
		reqBody["tool_choice"] = &tc
	}

	if reasoning != nil {
		reqBody["reasoning"] = reasoning
	}

	if provider != nil {
		reqBody["provider"] = provider
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return body, nil
}
