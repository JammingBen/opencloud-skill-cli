package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
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
func (c *Client) MakeRequest(path string, method string) (string, error) {
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

	if c.tokenSource != nil {
		token, err := c.tokenSource.Token()
		if err != nil {
			return "", fmt.Errorf("failed to get token: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}

	var respBody string
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %w", err)
		}
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, bodyBytes, "", "  "); err != nil {
			return "", err
		}
		respBody = dst.String()
	} else {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return respBody, nil
}
