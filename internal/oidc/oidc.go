package oidc

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/browser"
	"golang.design/x/clipboard"
	"golang.org/x/oauth2"
)

// Default to OpenCloud desktop client ID because it allows redirects to localhost
// also see https://docs.opencloud.eu/docs/dev/server/apis/http/authorization
const defaultClientID = "OpenCloudDesktop"

type OpenCloudConfig struct {
	OpenIdConnect struct {
		MetadataURL string `json:"metadata_url"`
		Authority   string `json:"authority"`
		ClientID    string `json:"client_id"`
	} `json:"openIdConnect"`
}

type OIDCMetadata struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
}

type OIDCClient struct {
	ServerURL    string
	ClientID     string
	Insecure     bool
	UseClipboard bool
}

// NewOIDCClient creates a new OIDCClient with the given parameters.
func NewOIDCClient(serverURL string, insecure bool, clientID string, useClipboard bool) *OIDCClient {
	return &OIDCClient{
		ServerURL:    serverURL,
		ClientID:     clientID,
		Insecure:     insecure,
		UseClipboard: useClipboard,
	}
}

// Login performs the OIDC login flow, saves the token to config, and handles the local server for the callback.
func (c *OIDCClient) Login() error {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: c.Insecure},
		},
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)

	// Discover OIDC endpoints
	endpoints, discoveredClientID, err := discoverOIDC(ctx, httpClient, c.ServerURL)
	if err != nil {
		return fmt.Errorf("failed to discover OIDC configuration: %w", err)
	}

	clientID := c.ClientID
	if clientID == "" {
		// Priority: Flag > discovered > default to desktop client id
		if discoveredClientID != "" {
			clientID = discoveredClientID
		} else {
			clientID = defaultClientID
		}
	}

	// Generate PKCE verifier and challenge
	verifier, challenge, err := generatePKCE()
	if err != nil {
		return fmt.Errorf("failed to generate PKCE: %w", err)
	}

	// Setup local server for callback
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("failed to start local server: %w", err)
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	redirectURL := fmt.Sprintf("http://127.0.0.1:%d/callback", port)

	// OAuth2 config
	conf := &oauth2.Config{
		ClientID: clientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  endpoints.AuthorizationEndpoint,
			TokenURL: endpoints.TokenEndpoint,
		},
		RedirectURL: redirectURL,
		Scopes:      []string{"openid", "profile", "email", "offline_access"},
	}

	// URL for authorization
	url := conf.AuthCodeURL("state",
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("code_challenge", challenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)

	fmt.Printf("Opening browser for authentication...\n")
	fmt.Printf("Client ID: %s\n", clientID)
	fmt.Printf("Redirect URI: %s\n", redirectURL)
	if err := browser.OpenURL(url); err != nil {
		msg := color.RedString("Failed to open browser: %v", err)
		fmt.Println(msg)
		fmt.Printf("Please open this URL manually: %s\n", url)
	}

	// Channel to receive the code
	codeChan := make(chan string)
	errChan := make(chan error)

	srv := &http.Server{}
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			errChan <- fmt.Errorf("no code in callback")
			fmt.Fprintf(w, "Authentication failed! No code received.")
			return
		}
		fmt.Fprintf(w, "Authentication successful! You can close this window.")
		codeChan <- code
	})

	go func() {
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Wait for code or error
	var code string
	select {
	case code = <-codeChan:
		// Exchange code for token
		token, err := conf.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", verifier))
		if err != nil {
			return fmt.Errorf("failed to exchange token: %w", err)
		}

		// Save to config
		cfg, err := LoadConfig()
		if err != nil {
			return err
		}
		cfg.ServerURL = c.ServerURL
		cfg.Insecure = c.Insecure
		cfg.ClientID = clientID
		cfg.TokenEndpoint = endpoints.TokenEndpoint

		if c.UseClipboard {
			if err := copyToClipboard(token.AccessToken); err != nil {
				return fmt.Errorf("failed to copy token to clipboard: %w", err)
			}
		} else {
			cfg.Token = token
		}

		if err := cfg.Save(); err != nil {
			return err
		}

		msg := color.GreenString("✓ Successfully logged in!")
		fmt.Println(msg)
		if c.UseClipboard {
			fmt.Println("The access token has been copied to your clipboard.")
		}
		srv.Shutdown(ctx)
		return nil
	case err := <-errChan:
		srv.Shutdown(ctx)
		return err
	case <-time.After(5 * time.Minute):
		srv.Shutdown(ctx)
		return fmt.Errorf("login timed out")
	}
}

// discoverOIDC tries to discover the OIDC metadata URL and client ID using multiple strategies:
// 1. Fetch config.json from the server and check for OIDC metadata URL or authority.
// 2. If not found, try the standard .well-known/openid-configuration endpoint.
// It returns the discovered OIDC metadata, client ID, and any error encountered.
func discoverOIDC(ctx context.Context, client *http.Client, serverURL string) (*OIDCMetadata, string, error) {
	serverURL = strings.TrimSuffix(serverURL, "/")

	// 1. Try config.json
	configURL := path.Join(serverURL, "config.json")
	req, _ := http.NewRequestWithContext(ctx, "GET", configURL, nil)
	resp, err := client.Do(req)

	var metadataURL string
	clientID := ""

	if err == nil && resp.StatusCode == http.StatusOK {
		var ocConfig OpenCloudConfig
		if err := json.NewDecoder(resp.Body).Decode(&ocConfig); err == nil {
			if ocConfig.OpenIdConnect.MetadataURL != "" {
				metadataURL = ocConfig.OpenIdConnect.MetadataURL
			} else if ocConfig.OpenIdConnect.Authority != "" {
				metadataURL = strings.TrimSuffix(ocConfig.OpenIdConnect.Authority, "/") + "/.well-known/openid-configuration"
			}
			if ocConfig.OpenIdConnect.ClientID != "" {
				clientID = ocConfig.OpenIdConnect.ClientID
			}
		}
		resp.Body.Close()
	}

	// 2. If no metadata URL from config.json, try standard .well-known
	if metadataURL == "" {
		metadataURL = serverURL + "/.well-known/openid-configuration"
	}

	// 3. Fetch OIDC metadata
	req, _ = http.NewRequestWithContext(ctx, "GET", metadataURL, nil)
	resp, err = client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch OIDC metadata: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("OIDC metadata request failed with status: %d", resp.StatusCode)
	}

	var metadata OIDCMetadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, "", fmt.Errorf("failed to decode OIDC metadata: %w", err)
	}

	return &metadata, clientID, nil
}

// generatePKCE creates a random PKCE verifier and its corresponding challenge.
func generatePKCE() (verifier string, challenge string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}
	verifier = base64.RawURLEncoding.EncodeToString(b)

	h := sha256.New()
	h.Write([]byte(verifier))
	challenge = base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	return verifier, challenge, nil
}

func copyToClipboard(text string) error {
	if err := clipboard.Init(); err != nil {
		return fmt.Errorf("failed to initialize clipboard: %w", err)
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}
