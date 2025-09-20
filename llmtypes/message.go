package llmtypes

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
