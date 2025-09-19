package llmtypes

type ParameterProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Items       *struct {
		Type string `json:"type"`
	} `json:"items,omitempty"`
}

type ToolParameters struct {
	Type       string                       `json:"type"`
	Properties map[string]ParameterProperty `json:"properties"`
	Required   []string                     `json:"required"`
}

// ToolSchema defines a json schema for a tool call.
// Notice how it only needs the "function" part and doesn't
// require you to specify the "type": "function" part
type ToolSchema struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  ToolParameters `json:"parameters"`
}

// -----------------------------
// MessageForLLM
// -----------------------------

type RoleEnum string

const (
	RoleUser      RoleEnum = "user"
	RoleAssistant RoleEnum = "assistant"
	RoleTool      RoleEnum = "tool"
	RoleSystem    RoleEnum = "system"
)

// MessageForLLM defines a typical message sent to an LLM.
// Check RoleEnum for the available roles.
type MessageForLLM struct {
	Role    RoleEnum `json:"role"`
	Content string   `json:"content"`
}

// Used when you want to send multi input content.
// For example, when sending an image, you use a list of dicts with the type "image_url" and the url.
type PartMessageForLLM struct {
	Role    RoleEnum      `json:"role"`
	Content []ContentPart `json:"content"`
}

type ContentPart struct {
	Type     string       `json:"type"`
	Text     string       `json:"text,omitempty"`
	ImageURL *ImageStruct `json:"image_url,omitempty"`
}

type ImageStruct struct {
	URL string `json:"url"`
}

// -----------------------------
// StructuredOutput
// -----------------------------

type StructuredOutputProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Items       *struct {
		Type string `json:"type"`
	} `json:"items,omitempty"`
	Enum *[]string `json:"enum,omitempty"`
}

type StructuredOutputSchemaDefinition struct {
	Type                 string                              `json:"type"`
	Properties           map[string]StructuredOutputProperty `json:"properties"`
	Required             []string                            `json:"required"`
	AdditionalProperties bool                                `json:"additionalProperties"`
}

// StructuredOutput defines a structured output from an LLM.
// Notice how it only needs the "json_schema" part and doesn't
// require you to specify the "type": "json_schema" part
type StructuredOutputSchema struct {
	Name   string                           `json:"name"`
	Strict bool                             `json:"strict"`
	Schema StructuredOutputSchemaDefinition `json:"schema"`
}

// -----------------------------
// ReasoningConfig
// -----------------------------

// This will only work for certain models. This might break with the wrong models.
// Also, don't add both Effort and MaxTokens or it will break.
type ReasoningConfig struct {
	// OpenAI style reasoning effort settings
	Effort *string `json:"effort,omitempty"`

	// Non-OpenAI style reasoning effort settings
	// Gemini uses this for example
	MaxTokens *int `json:"max_tokens,omitempty"`
}

// -----------------------------
// ProviderConfig
// -----------------------------

type ProviderSortEnum string

const (
	// Prioritize lowest price
	SortPrice ProviderSortEnum = "price"

	// Prioritize highest throughput
	SortThroughput ProviderSortEnum = "throughput"

	// Prioritize lowest latency
	SortLatency ProviderSortEnum = "latency"
)

// ProviderConfig defines the configuration for a provider for OpenRouter.
type ProviderConfig struct {
	// List of providers to use, in order of priority
	Order *[]string `json:"order,omitempty"`

	// To allow use of other providers if none of your chosen providers are available
	AllowFallbacks bool `json:"allow_fallbacks"`

	// Only allow providers that require your set parameters (structured output, reasoning, etc)
	RequireParameters bool `json:"require_parameters"`

	// Allow providers to collect data or not
	DataCollection bool `json:"data_collection"`

	// Zero Data Retention providers only.
	ZDR bool `json:"zdr"`

	// List of providers to use, no others will be used if this is set
	Only *[]string `json:"only,omitempty"`

	// List of providers to ignore
	Ignore *[]string `json:"ignore,omitempty"`

	// Sort providers by price, throughput, or latency
	Sort *ProviderSortEnum `json:"sort,omitempty"`
}

// -----------------------------------------
// Groq types and schemas
// -----------------------------------------

type GroqTranscriptionResponse struct {
	Text     string  `json:"text"`
	Duration float64 `json:"duration"`
}

// -----------------------------------------
// OpenRouter types and schemas
// -----------------------------------------

type ToolChoiceEnum string

const (
	// Model can only use tools
	ToolChoiceRequired ToolChoiceEnum = "required"

	// Let model decide which tools to use
	ToolChoiceAuto ToolChoiceEnum = "auto"
)

type OpenRouterRequest struct {
	Model          string            `json:"model"`
	Messages       []MessageForLLM   `json:"messages"`
	Temperature    *float64          `json:"temperature,omitempty"`
	MaxTokens      *int              `json:"max_tokens,omitempty"`
	ResponseFormat *map[string]any   `json:"response_format,omitempty"`
	Tools          *[]map[string]any `json:"tools,omitempty"`
	ToolChoice     *ToolChoiceEnum   `json:"tool_choice,omitempty"`
	Reasoning      *ReasoningConfig  `json:"reasoning,omitempty"`
	Provider       *ProviderConfig   `json:"provider,omitempty"`
}

type OpenRouterResponse struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Created  int64  `json:"created"`
	Choices  []struct {
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role      RoleEnum `json:"role"`
			Content   string   `json:"content"`
			Refusal   *string  `json:"refusal"`
			Reasoning *string  `json:"reasoning"`
			ToolCalls *[]struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Function struct {
					Name      string `json:"name"`
					Arguments string `json:"arguments"`
				} `json:"function"`
			} `json:"tool_calls,omitempty"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
