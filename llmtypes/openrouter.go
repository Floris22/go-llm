package llmtypes

type OpenRouterRequest struct {
	Model          string           `json:"model"`
	Messages       []MessageForLLM  `json:"messages"`
	Temperature    *float64         `json:"temperature,omitempty"`
	MaxTokens      *int             `json:"max_tokens,omitempty"`
	ResponseFormat *map[string]any  `json:"response_format,omitempty"`
	Tools          []map[string]any `json:"tools,omitempty"`
	ToolChoice     *ToolChoiceEnum  `json:"tool_choice,omitempty"`
	Reasoning      *ReasoningConfig `json:"reasoning,omitempty"`
	Provider       *ProviderConfig  `json:"provider,omitempty"`
}

type OpenRouterRequestWithParts struct {
	Model          string              `json:"model"`
	Messages       []PartMessageForLLM `json:"messages"`
	Temperature    *float64            `json:"temperature,omitempty"`
	MaxTokens      *int                `json:"max_tokens,omitempty"`
	ResponseFormat *map[string]any     `json:"response_format,omitempty"`
	Tools          []map[string]any    `json:"tools,omitempty"`
	ToolChoice     *ToolChoiceEnum     `json:"tool_choice,omitempty"`
	Reasoning      *ReasoningConfig    `json:"reasoning,omitempty"`
	Provider       *ProviderConfig     `json:"provider,omitempty"`
}

type OpenRouterResponse struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Created  int64  `json:"created"`
	Choices  []struct {
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role      RoleEnum                 `json:"role"`
			Content   string                   `json:"content"`
			Refusal   *string                  `json:"refusal"`
			Reasoning *string                  `json:"reasoning"`
			ToolCalls []MessageForLLMToolCalls `json:"tool_calls,omitempty"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
