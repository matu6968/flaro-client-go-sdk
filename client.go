package flaro

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	BaseURL = "https://sb.flaroapp.pl"
)

// Client represents the Flaro API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient creates a new Flaro API client with the provided API key
func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: BaseURL,
		apiKey:  apiKey,
	}
}

// NewClientFromEnv creates a new Flaro API client using API key from environment variable
func NewClientFromEnv() (*Client, error) {
	apiKey := os.Getenv("FLARO_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("FLARO_API_KEY environment variable is required but not set")
	}
	return NewClient(apiKey), nil
}

// NewClientWithOptions creates a new Flaro API client with custom options
func NewClientWithOptions(baseURL, apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

// generateCodeVerifier makes a high-entropy random string (43–128 chars)
func (c *Client) generateCodeVerifier(length int) (string, error) {
	if length < 43 || length > 128 {
		return "", fmt.Errorf("length must be between 43 and 128")
	}
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	// base64url encode without padding
	verifier := base64.RawURLEncoding.EncodeToString(buf)
	// trim/pad to exact length
	if len(verifier) > length {
		verifier = verifier[:length]
	}
	return verifier, nil
}

// codeChallengeFromVerifier computes the S256 code_challenge
func (c *Client) codeChallengeFromVerifier(verifier string) string {
	h := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h[:]) // no padding
}

// makeRequest makes an HTTP request to the Flaro API
func (c *Client) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	return c.makeAuthenticatedRequest(method, endpoint, body, "")
}

// makeAuthenticatedRequest makes an HTTP request to the Flaro API with optional authentication
func (c *Client) makeAuthenticatedRequest(method, endpoint string, body interface{}, accessToken string) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, c.baseURL+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.apiKey)
	req.Header.Set("x-client-info", "flaro-go-sdk/1.0.0")

	// Add authorization header if access token is provided
	if accessToken != "" {
		req.Header.Set("authorization", "Bearer "+accessToken)
	}

	return c.httpClient.Do(req)
}
