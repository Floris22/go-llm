package llmtypes

type GroqTranscriptionResponse struct {
	Text     string  `json:"text"`
	Duration float64 `json:"duration"`
}
