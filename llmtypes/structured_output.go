package llmtypes

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
