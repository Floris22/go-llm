package llmtypes

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
