package clients

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"time"

	h "github.com/Floris22/go-llm/v2/internal/helpers"
	t "github.com/Floris22/go-llm/v2/llmtypes"
)

type OpenRouterClient interface {
	GenerateText(
		messages []t.MessageForLLM,
		messageParts []t.PartMessageForLLM,
		model string,
		temperature *float64,
		maxTokens *int,
		timeOut *int,
		reasoning *t.ReasoningConfig,
		provider *t.ProviderConfig,
	) (t.OpenRouterResponse, error)

	GenerateTools(
		messages []t.MessageForLLM,
		messageParts []t.PartMessageForLLM,
		tools []t.ToolSchema,
		model string,
		temperature *float64,
		maxTokens *int,
		timeOut *int,
		reasoning *t.ReasoningConfig,
		provider *t.ProviderConfig,
	) (t.OpenRouterResponse, error)

	GenerateStructured(
		messages []t.MessageForLLM,
		messageParts []t.PartMessageForLLM,
		schema t.StructuredOutputSchema,
		model string,
		temperature *float64,
		maxTokens *int,
		timeOut *int,
		reasoning *t.ReasoningConfig,
		provider *t.ProviderConfig,
	) (t.OpenRouterResponse, error)
}

type openRouterClient struct {
	apiKey string
}

func NewOpenRouterClient(apiKey string) OpenRouterClient {
	return &openRouterClient{apiKey: apiKey}
}

func (c *openRouterClient) GenerateText(
	messages []t.MessageForLLM,
	messageParts []t.PartMessageForLLM,
	model string,
	temperature *float64,
	maxTokens *int,
	timeOut *int,
	reasoning *t.ReasoningConfig,
	provider *t.ProviderConfig,
) (t.OpenRouterResponse, error) {
	timeoutValue := 15
	if timeOut != nil {
		timeoutValue = *timeOut
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutValue)*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.apiKey,
	}

	body, err := h.CreateRequestBody(messages, messageParts, model, temperature, maxTokens, nil, nil, reasoning, provider)
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	respBody, statusCode, err := h.PostReq(
		ctx, "https://openrouter.ai/api/v1/chat/completions", headers, body, nil,
	)
	if statusCode != 200 {
		return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API returned status code %d with error: %s", statusCode, string(respBody))
	}
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	var response t.OpenRouterResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	return response, nil
}

func (c *openRouterClient) GenerateTools(
	messages []t.MessageForLLM,
	messageParts []t.PartMessageForLLM,
	tools []t.ToolSchema,
	model string,
	temperature *float64,
	maxTokens *int,
	timeOut *int,
	reasoning *t.ReasoningConfig,
	provider *t.ProviderConfig,
) (t.OpenRouterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.apiKey,
	}

	body, err := h.CreateRequestBody(messages, messageParts, model, temperature, maxTokens, nil, &tools, reasoning, provider)
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	respBody, statusCode, err := h.PostReq(
		ctx, "https://openrouter.ai/api/v1/chat/completions", headers, body, nil,
	)
	if statusCode != 200 {
		return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API returned status code %d with error: %s", statusCode, string(respBody))
	}
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	var response t.OpenRouterResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	return response, nil
}

func (c *openRouterClient) GenerateStructured(
	messages []t.MessageForLLM,
	messageParts []t.PartMessageForLLM,
	schema t.StructuredOutputSchema,
	model string,
	temperature *float64,
	maxTokens *int,
	timeOut *int,
	reasoning *t.ReasoningConfig,
	provider *t.ProviderConfig,
) (t.OpenRouterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.apiKey,
	}

	body, err := h.CreateRequestBody(messages, messageParts, model, temperature, maxTokens, &schema, nil, reasoning, provider)
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	respBody, statusCode, err := h.PostReq(
		ctx, "https://openrouter.ai/api/v1/chat/completions", headers, body, nil,
	)
	if statusCode != 200 {
		return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API returned status code %d with error: %s", statusCode, string(respBody))
	}
	if err != nil {
		return t.OpenRouterResponse{}, err
	}

	var response t.OpenRouterResponse
	err = json.Unmarshal(respBody, &response)
	return response, err
}
