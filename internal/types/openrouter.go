// Package types forsees types for multiple API responses
package types

type OpenRouterRequest struct {
	Model          string           `json:"model"`
	Messages       []map[string]any `json:"messages"`
	Temperature    float64          `json:"temperature"`
	MaxTokens      int              `json:"max_tokens"`
	ResponseFormat map[string]any   `json:"response_format,omitempty"`
	Tools          []map[string]any `json:"tools,omitempty"`
	ToolChoice     string           `json:"tool_choice,omitempty"`
	Reasoning      map[string]any   `json:"reasoning,omitempty"`
}

type OpenRouterResponse struct {
	Provider string              `json:"provider"`
	Model    string              `json:"model"`
	Created  int64               `json:"created"`
	Choices  []OpenRouterChoices `json:"choices"`
	Usage    OpenRouterUsage     `json:"usage"`
}

type OpenRouterChoices struct {
	FinishReason string            `json:"finish_reason"`
	Message      OpenRouterMessage `json:"message"`
}

type OpenRouterMessage struct {
	Role      string                `json:"role"`
	Content   string                `json:"content"`
	Refusal   *string               `json:"refusal"`
	Reasoning *string               `json:"reasoning"`
	ToolCalls *[]OpenRouterToolCall `json:"tool_calls"`
}

type OpenRouterToolCall struct {
	ID       string             `json:"id"`
	Type     string             `json:"type"`
	Function OpenRouterFunction `json:"function"`
}

type OpenRouterFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type OpenRouterUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
