package llmtypes

type RoleEnum string

const (
	RoleUser      RoleEnum = "user"
	RoleAssistant RoleEnum = "assistant"
	RoleTool      RoleEnum = "tool"
	RoleSystem    RoleEnum = "system"
)

type ProviderSortEnum string

const (
	// Prioritize lowest price
	SortPrice ProviderSortEnum = "price"

	// Prioritize highest throughput
	SortThroughput ProviderSortEnum = "throughput"

	// Prioritize lowest latency
	SortLatency ProviderSortEnum = "latency"
)

type ToolChoiceEnum string

const (
	// Model can only use tools
	ToolChoiceRequired ToolChoiceEnum = "required"

	// Let model decide which tools to use
	ToolChoiceAuto ToolChoiceEnum = "auto"
)
