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
	apiKey                    string
	retryModel                string
	retryModelReasoningConfig *t.ReasoningConfig
	enableRetry               bool
}

// Creates a new open router OpenRouterClient.
// If retry is enabled and the model name is "" we
// set openai/gpt-oss-120b:nitro as default and reasoningConfig to nil.
// Otherwise, it is your responsibility to provide a valid model name and reasoningConfig for that model.
func NewOpenRouterClient(
	apiKey string,
	enableRetry bool,
	retryModel string,
	retryModelReasoningConfig *t.ReasoningConfig,
) OpenRouterClient {
	if enableRetry && retryModel == "" {
		retryModel = "openai/gpt-oss-120b:nitro"
		retryModelReasoningConfig = nil
	}
	return &openRouterClient{
		apiKey:                    apiKey,
		retryModel:                retryModel,
		enableRetry:               enableRetry,
		retryModelReasoningConfig: retryModelReasoningConfig,
	}
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
	if err != nil {
		return t.OpenRouterResponse{}, err
	}
	if statusCode != 200 {
		// 408 == request timed out, 429 == rate limited, 502 model down or invalid response
		if (statusCode == 408 || statusCode == 429 || statusCode == 502) && c.enableRetry {
			body, err := h.CreateRequestBody(messages, messageParts, c.retryModel, temperature, maxTokens, nil, nil, c.retryModelReasoningConfig, nil)
			if err != nil {
				return t.OpenRouterResponse{}, err
			}
			respBody, err = h.DoReqWithRetries(ctx, headers, body)
			if err != nil {
				return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API failed retry after 3 attempts: %s", string(respBody))
			}
		} else {
			return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API returned status code %d with error: %s", statusCode, string(respBody))
		}
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
	if err != nil {
		return t.OpenRouterResponse{}, err
	}
	if statusCode != 200 {
		if (statusCode == 408 || statusCode == 429 || statusCode == 502) && c.enableRetry {
			body, err := h.CreateRequestBody(messages, messageParts, c.retryModel, temperature, maxTokens, nil, &tools, c.retryModelReasoningConfig, nil)
			respBody, err = h.DoReqWithRetries(ctx, headers, body)
			if err != nil {
				return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API failed retry after 3 attempts: %s", string(respBody))
			}
		} else {
			return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API returned status code %d with error: %s", statusCode, string(respBody))
		}
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
	if err != nil {
		return t.OpenRouterResponse{}, err
	}
	if statusCode != 200 {
		if (statusCode == 408 || statusCode == 429 || statusCode == 502) && c.enableRetry {
			body, err := h.CreateRequestBody(messages, messageParts, c.retryModel, temperature, maxTokens, &schema, nil, c.retryModelReasoningConfig, nil)
			respBody, err = h.DoReqWithRetries(ctx, headers, body)
			if err != nil {
				return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API failed retry after 3 attempts: %s", string(respBody))
			}
		} else {
			return t.OpenRouterResponse{}, fmt.Errorf("OpenRouter API returned status code %d with error: %s", statusCode, string(respBody))
		}
	}

	var response t.OpenRouterResponse
	err = json.Unmarshal(respBody, &response)
	return response, err
}
