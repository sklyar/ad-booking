package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"
)

// HTTPResponse represents the data returned from an HTTP request.
type HTTPResponse struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
	Duration   time.Duration
}

// HTTPClient is a client for making HTTP requests.
type HTTPClient struct {
	baseClient *http.Client
	baseURL    string
}

// Do send an HTTP POST request to a specified endpoint using the HTTPClient.
// It returns a response containing details of the response received.
func (c HTTPClient) Do(t *testing.T, ctx context.Context, endpoint string, reqBody []byte) HTTPResponse {
	t.Helper()

	reqURL, err := url.JoinPath(c.baseURL, endpoint)
	require.NoError(t, err, "failed to join URL path")

	req, err := http.NewRequest("POST", reqURL, bytes.NewReader(reqBody))
	require.NoError(t, err, "failed to create request")

	h := make(http.Header)
	h.Set("Content-Type", "application/json")
	req.Header = h

	now := time.Now()

	resp, err := c.baseClient.Do(req.WithContext(ctx))
	require.NoError(t, err, "failed to send request")
	defer resp.Body.Close()

	elapsed := time.Since(now)

	dec := json.NewDecoder(resp.Body)
	rawRespBody := json.RawMessage{}
	err = dec.Decode(&rawRespBody)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		require.NoErrorf(t, err, "failed to decode response")
	}

	return HTTPResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       rawRespBody,
		Duration:   elapsed,
	}
}
