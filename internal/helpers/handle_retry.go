package helpers

import (
	"context"
	"math"
	"time"
)

func DoGroqWithRetries(
	ctx context.Context,
	headers map[string]string,
	body []byte,
) ([]byte, error) {
	var lastErr error
	for attempt := range 5 {
		respBody, statusCode, err := PostReq(
			ctx, "https://api.groq.com/openai/v1/audio/transcriptions", headers, body, nil,
		)
		if err == nil && statusCode == 200 {
			return respBody, nil
		}
		lastErr = err

		if attempt+1 < 5 {
			time.Sleep(time.Duration(100*math.Pow(2, float64(attempt))) * time.Millisecond)
		}
	}
	return nil, lastErr
}

func DoReqWithRetries(
	ctx context.Context,
	headers map[string]string,
	body []byte,
) ([]byte, error) {
	var lastErr error
	for attempt := range 3 {
		respBody, statusCode, err := PostReq(
			ctx, "https://openrouter.ai/api/v1/chat/completions", headers, body, nil,
		)
		if err == nil && statusCode == 200 {
			return respBody, nil
		}
		lastErr = err

		if attempt+1 < 3 {
			time.Sleep(time.Duration(100*math.Pow(2, float64(attempt))) * time.Millisecond)
		}
	}
	return nil, lastErr
}
