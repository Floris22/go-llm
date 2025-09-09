package helpers

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	t "github.com/Floris22/go-llm/internal/types"
	r "github.com/Floris22/go-llm/internal/utils/requests"
)

func TranscribeGroq(
	model string,
	language string,
	apiKey string,
	audioURL *string,
	audioBytes *[]byte,
	timeOut *int,
) (t.GroqTranscriptionResponse, error) {
	timeoutValue := 30
	if timeOut != nil {
		timeoutValue = *timeOut
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutValue)*time.Second)
	defer cancel()

	if language == "" {
		language = "en"
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	writer.WriteField("model", model)
	writer.WriteField("language", language)
	writer.WriteField("response_format", "verbose_json")

	if audioBytes != nil {
		var file io.Reader
		file = bytes.NewReader(*audioBytes)
		part, _ := writer.CreateFormFile("file", "file.wav")
		io.Copy(part, file)
	} else {
		writer.WriteField("url", *audioURL)
	}
	writer.Close()

	headers := map[string]string{
		"Content-Type":  writer.FormDataContentType(),
		"Authorization": "Bearer " + apiKey,
	}

	respBody, statusCode, err := r.PostReq(
		ctx, "https://api.groq.com/openai/v1/audio/transcriptions", headers, buf.Bytes(), nil,
	)
	if statusCode != 200 {
		return t.GroqTranscriptionResponse{}, fmt.Errorf("Groq API returned status code %d with error: %s", statusCode, string(respBody))
	}
	if err != nil {
		return t.GroqTranscriptionResponse{}, err
	}

	var response t.GroqTranscriptionResponse
	err = json.Unmarshal(respBody, &response)
	return response, err
}
