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
