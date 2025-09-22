package flaro

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// SignUp creates a new user account
func (c *Client) SignUp(email, password string) (*AuthResponse, error) {
	// Generate PKCS code verifier and challenge
	verifier, err := c.generateCodeVerifier(64)
	if err != nil {
		return nil, fmt.Errorf("failed to generate code verifier: %w", err)
	}

	challenge := c.codeChallengeFromVerifier(verifier)

	// Create sign up request
	req := SignUpRequest{
		Email:    email,
		Password: password,
		Data:     nil,
		GotrueMetaSecurity: GotrueMetaSecurity{
			CaptchaToken: nil,
		},
		CodeChallenge:       challenge,
		CodeChallengeMethod: "s256",
	}

	// Make the request
	resp, err := c.makeRequest("POST", "/auth/v1/signup", req)
	if err != nil {
		return nil, fmt.Errorf("failed to make sign up request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("sign up failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response
	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return nil, fmt.Errorf("failed to parse auth response: %w", err)
	}

	return &authResp, nil
}

// SignIn authenticates an existing user
func (c *Client) SignIn(email, password string) (*AuthResponse, error) {
	// Create sign in request
	req := SignInRequest{
		Email:    email,
		Password: password,
		GotrueMetaSecurity: GotrueMetaSecurity{
			CaptchaToken: nil,
		},
	}

	// Make the request
	resp, err := c.makeRequest("POST", "/auth/v1/token?grant_type=password", req)
	if err != nil {
		return nil, fmt.Errorf("failed to make sign in request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("sign in failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiErr
	}

	// Parse successful response
	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return nil, fmt.Errorf("failed to parse auth response: %w", err)
	}

	return &authResp, nil
}

// RefreshToken refreshes an access token using a refresh token
func (c *Client) RefreshToken(refreshToken string) (*AuthResponse, error) {
	req := RefreshTokenRequest{
		RefreshToken: refreshToken,
	}

	endpoint := "/auth/v1/token?grant_type=refresh_token"
	resp, err := c.makeRequest("POST", endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("refresh token failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("refresh token failed: %s", apiErr.Message)
	}

	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return nil, fmt.Errorf("failed to parse refresh token response: %w", err)
	}

	return &authResp, nil
}

// ChangePassword changes a user's password
func (c *Client) ChangePassword(accessToken, newPassword string) (*ChangePasswordResponse, error) {
	req := ChangePasswordRequest{
		Password: newPassword,
	}

	endpoint := "/auth/v1/user"
	resp, err := c.makeAuthenticatedRequest("POST", endpoint, req, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to change password: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode != 200 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("change password failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("change password failed: %s", apiErr.Message)
	}

	var changeResp ChangePasswordResponse
	if err := json.Unmarshal(body, &changeResp); err != nil {
		return nil, fmt.Errorf("failed to parse change password response: %w", err)
	}

	return &changeResp, nil
}

// SignOff logs out the current user. scope can be "local" (this device) or other scopes if supported.
// API expects a POST to /auth/v1/logout?scope=<scope> with Authorization and apikey headers. Returns 204.
func (c *Client) SignOff(accessToken, scope string) error {
	if scope == "" {
		scope = "local"
	}
	v := url.Values{}
	v.Set("scope", scope)
	endpoint := "/auth/v1/logout?" + v.Encode()

	resp, err := c.makeAuthenticatedRequest("POST", endpoint, map[string]string{}, accessToken)
	if err != nil {
		return fmt.Errorf("failed to sign off: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 204 {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("sign off failed with status %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("sign off failed: %s", apiErr.Message)
	}
	return nil
}
