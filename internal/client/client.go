package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
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

func NewClient(baseURL string, insecure bool, ts oauth2.TokenSource) *Client {
	return &Client{
		baseURL:     baseURL,
		insecure:    insecure,
		tokenSource: ts,
	}
}

// MakeRequest makes an HTTP request to the specified URL with the given method and returns the response
func (c *Client) MakeRequest(path string, method string, body string) (string, error) {
	fullURL, err := url.JoinPath(c.baseURL, "graph", path)
	if err != nil {
		return "", fmt.Errorf("failed to join path: %w", err)
	}

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set up http transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.insecure},
	}

	// Get token source for authentication
	if c.tokenSource != nil {
		token, err := c.tokenSource.Token()
		if err != nil {
			return "", fmt.Errorf("failed to get token: %w", err)
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
		return "", fmt.Errorf("failed to execute request: %w", err)
	}

	slog.Debug("Received response", "status", resp.StatusCode)

	// Read and return response body
	return readBody(resp)
}

// readBody reads the response body and returns it as a string.
// It also checks for HTTP errors and returns an error if the status code is 400 or above.
func readBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return "", nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, bodyBytes, "", "  "); err != nil {
		return "", err
	}

	return dst.String(), nil
}
