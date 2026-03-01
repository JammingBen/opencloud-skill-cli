package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

type Client struct {
	baseURL     string
	insecure    bool
	tokenSource oauth2.TokenSource
}

type Response struct {
	StatusCode int
	Body       string
}

func NewClient(baseURL string, insecure bool, ts oauth2.TokenSource) *Client {
	return &Client{
		baseURL:     baseURL,
		insecure:    insecure,
		tokenSource: ts,
	}
}

// MakeRequest makes an HTTP request to the specified URL with the given method and body and returns the response
func (c *Client) MakeRequest(path string, method string, body string) (*Response, error) {
	fullURL, err := url.JoinPath(c.baseURL, "graph", path)
	if err != nil {
		return nil, fmt.Errorf("failed to join path: %w", err)
	}

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set up http transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.insecure},
	}

	// Get token source for authentication
	if c.tokenSource != nil {
		token, err := c.tokenSource.Token()
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}

	// Add body if provided
	if body != "" {
		req.Body = io.NopCloser(bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
	}

	slog.Debug("Making request", "method", method, "url", fullURL)

	// Do request
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	slog.Debug("Received response", "status", resp.StatusCode)

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return readBody(resp)
}

// readBody reads the response and returns a Response struct containing the status code and body as a string
func readBody(resp *http.Response) (*Response, error) {
	defer resp.Body.Close()

	r := &Response{
		StatusCode: resp.StatusCode,
	}

	if resp.Body == nil {
		return r, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, fmt.Errorf("failed to read response body: %w", err)
	}

	if len(bodyBytes) == 0 {
		return r, nil
	}

	r.Body = string(bodyBytes)
	return r, nil
}
