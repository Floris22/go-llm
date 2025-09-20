package llmtypes

// This will only work for certain models. This might break with the wrong models.
// Also, don't add both Effort and MaxTokens or it will break.
type ReasoningConfig struct {
	// OpenAI style reasoning effort settings
	Effort *string `json:"effort,omitempty"`

	// Non-OpenAI style reasoning effort settings
	// Gemini uses this for example
	MaxTokens *int `json:"max_tokens,omitempty"`

	// Some models (e.g. Grok 4 fast), have a reasoning and non-reasoning mode.
	// This is enabled by default, but you can disable it.
	Enabled *bool `json:"enabled,omitempty"`
}
