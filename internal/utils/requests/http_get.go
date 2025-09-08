package requests

import (
	"context"
	"io"
	"net/http"
)

// GetReq requires a context and url, other parameters are optional.
// Returns the response body as bytes, the HTTP status code and any error.
// HTTP status code 0 means an error / cancellation occured before any request was sent.
func GetReq(
	ctx context.Context,
	url string,
	headers map[string]string,
	client *http.Client,
) ([]byte, int, error) {
	if client == nil {
		client = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return respBody, resp.StatusCode, err
}
