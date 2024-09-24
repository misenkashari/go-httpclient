package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// Receiver function that performs the underline http request.
// This will ensure a high level of abstraction on the http-client,
// making it more easy to use
func (c *HttpClient) sendAsyncRequest(method, path string, body []byte, config *AdditionalConfig) (Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.client.Timeout)
	defer cancel()

	url := path
	if c.baseUrl != "" {
		url = fmt.Sprintf("%s%s", c.baseUrl, path)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return Response{}, err
	}

	if config != nil {
		for key, value := range config.Headers {
			req.Header.Set(key, value)
		}
	}

	responseChan := make(chan Response)
	go func() {
		resp, err := c.client.Do(req)
		if err != nil {
			responseChan <- Response{Err: err}
			return
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			responseChan <- Response{Err: err}
			return
		}

		responseChan <- Response{
			Body:       string(bodyBytes),
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
		}
	}()

	select {
	case resp := <-responseChan:
		return resp, nil
	case <-ctx.Done():
		return Response{}, ctx.Err()
	}
}
