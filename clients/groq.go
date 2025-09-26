package clients

import (
	"errors"
	"os"
	"strings"

	h "github.com/Floris22/go-llm/v2/internal/helpers"
	t "github.com/Floris22/go-llm/v2/llmtypes"
)

type GroqClient interface {
	Transcribe(
		model string,
		language string,
		audioURL *string,
		audioBytes *[]byte,
		timeOut *int,
	) (t.GroqTranscriptionResponse, error)
}

type groqClient struct {
	apiKey string
}

func NewGroqClient(apiKey string) GroqClient {
	return &groqClient{
		apiKey: apiKey,
	}
}

func (c *groqClient) Transcribe(
	model string,
	language string,
	audioURL *string,
	audioBytes *[]byte,
	timeOut *int,
) (t.GroqTranscriptionResponse, error) {
	if audioURL != nil && audioBytes != nil {
		return t.GroqTranscriptionResponse{}, errors.New("Either use audioURL or audioBytes but not both.")
	}

	if audioURL == nil && audioBytes == nil {
		return t.GroqTranscriptionResponse{}, errors.New("Either audioURL or audioBytes must be provided.")
	}

	if audioURL != nil {
		resp, err := h.TranscribeGroq(model, language, c.apiKey, audioURL, nil, timeOut)
		return resp, err
	} else {
		// create temp file from bytes
		tempFile, err := os.CreateTemp("", "audio.wav")
		if err != nil {
			return t.GroqTranscriptionResponse{}, err
		}
		defer os.Remove(tempFile.Name())
		_, err = tempFile.Write(*audioBytes)
		if err != nil {
			return t.GroqTranscriptionResponse{}, err
		}
		chunkPaths, tempDir, err := h.SplitAudio(tempFile.Name(), 20*1024*1024)
		if err != nil {
			return t.GroqTranscriptionResponse{}, err
		}
		defer os.RemoveAll(tempDir)

		var transcriptions []string
		var totalDuration float64
		for _, chunkPath := range chunkPaths {
			audioBytes, err := os.ReadFile(chunkPath)
			resp, err := h.TranscribeGroq(model, language, c.apiKey, nil, &audioBytes, timeOut)
			if err != nil {
				return t.GroqTranscriptionResponse{}, err
			}
			transcriptions = append(transcriptions, resp.Text)
			totalDuration += resp.Duration
		}

		return t.GroqTranscriptionResponse{
			Text:     strings.Join(transcriptions, ""),
			Duration: totalDuration,
		}, nil

	}

}
