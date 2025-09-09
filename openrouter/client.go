package openrouter

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"time"

	h "github.com/Floris22/go-llm/internal/helpers"
	t "github.com/Floris22/go-llm/internal/types"
	r "github.com/Floris22/go-llm/internal/utils/requests"
)

type OpenRouterClient interface {
	GenerateText(
		messages []map[string]any,
		model string,
		temperature *float64,
		maxTokens *int,
		timeOut *int,
		reasoning *map[string]any,
		provider *map[string]any,
	) (t.OpenRouterResponse, error)

	GenerateTools(
		messages []map[string]any,
		tools []map[string]any,
		model string,
		temperature *float64,
		maxTokens *int,
		timeOut *int,
		reasoning *map[string]any,
		provider *map[string]any,
	) (t.OpenRouterResponse, error)

	GenerateStuctured(
		messages []map[string]any,
		schema map[string]any,
		model string,
		temperature *float64,
		maxTokens *int,
		timeOut *int,
		reasoning *map[string]any,
		provider *map[string]any,
	) (t.OpenRouterResponse, error)
}

type client struct {
	apiKey string
}

func NewClient(apiKey string) OpenRouterClient {
	return &client{apiKey: apiKey}
}

func (c *client) GenerateText(
	messages []map[string]any,
	model string,
	temperature *float64,
	maxTokens *int,
	timeOut *int,
	reasoning *map[string]any,
	provider *map[string]any,
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

	body := h.CreateRequestBody(messages, model, temperature, maxTokens, nil, nil, reasoning, provider)
	respBody, statusCode, err := r.PostReq(
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

func (c *client) GenerateTools(
	messages []map[string]any,
	tools []map[string]any,
	model string,
	temperature *float64,
	maxTokens *int,
	timeOut *int,
	reasoning *map[string]any,
	provider *map[string]any,
) (t.OpenRouterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.apiKey,
	}

	body := h.CreateRequestBody(messages, model, temperature, maxTokens, nil, &tools, reasoning, provider)

	respBody, statusCode, err := r.PostReq(
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

func (c *client) GenerateStuctured(
	messages []map[string]any,
	schema map[string]any,
	model string,
	temperature *float64,
	maxTokens *int,
	timeOut *int,
	reasoning *map[string]any,
	provider *map[string]any,
) (t.OpenRouterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + c.apiKey,
	}

	body := h.CreateRequestBody(messages, model, temperature, maxTokens, &schema, nil, reasoning, provider)

	respBody, statusCode, err := r.PostReq(
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
