package oidc

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

type Config struct {
	ServerURL     string        `json:"server_url"`
	Token         *oauth2.Token `json:"token"`
	Insecure      bool          `json:"insecure"`
	ClientID      string        `json:"client_id"`
	TokenEndpoint string        `json:"token_endpoint"`
}

// GetConfigPath returns the path to the configuration file
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "opencloud-cli", "config.json"), nil
}

// LoadConfig reads the configuration from the config file
func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Save writes the configuration to the config file
func (c *Config) Save() error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

// Clear removes the token and server URL from the config and saves it
func (c *Config) Clear() error {
	c.Token = nil
	c.ServerURL = ""
	c.Insecure = false
	c.ClientID = ""
	c.TokenEndpoint = ""
	return c.Save()
}

// GetAccessToken returns the access token from the config, or an empty string if not set
func (c *Config) GetAccessToken() string {
	if c.Token == nil {
		return ""
	}
	return c.Token.AccessToken
}

// GetTokenSource returns an oauth2.TokenSource that automatically refreshes the token and saves it back to the config file.
func (c *Config) GetTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	if c.Token == nil || c.ClientID == "" || c.TokenEndpoint == "" {
		return nil, nil
	}

	// If insecure is set, create a custom HTTP client and add it to the context
	if c.Insecure {
		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	}

	conf := &oauth2.Config{
		ClientID: c.ClientID,
		Endpoint: oauth2.Endpoint{
			TokenURL: c.TokenEndpoint,
		},
	}

	ts := conf.TokenSource(ctx, c.Token)
	return &savingTokenSource{
		ts:  ts,
		cfg: c,
	}, nil
}

// savingTokenSource is a custom oauth2.TokenSource that saves the token back to the config file whenever it is refreshed.
type savingTokenSource struct {
	ts  oauth2.TokenSource
	cfg *Config
}

// Token retrieves a token from the underlying TokenSource, saves it to the config if it has changed, and returns it.
func (s *savingTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.ts.Token()
	if err != nil {
		return nil, err
	}

	// If the token was refreshed, save it to the config
	if s.cfg.Token == nil || token.AccessToken != s.cfg.Token.AccessToken {
		s.cfg.Token = token
		if err := s.cfg.Save(); err != nil {
			return nil, err
		}
	}

	return token, nil
}
