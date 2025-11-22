// Package bannerify provides a simple client for the Bannerify API
package bannerify

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"
)

const (
	defaultBaseURL = "https://api.bannerify.co/v1"
	userAgent      = "bannerify-go/0.1.0"
)

// Client is the Bannerify API client
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Bannerify client
func NewClient(apiKey string, opts ...Option) *Client {
	client := &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Option is a functional option for configuring the client
type Option func(*Client)

// WithBaseURL sets a custom base URL
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithTimeout sets a custom timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// Response represents an API response
type Response struct {
	Result []byte
	Error  *ErrorResponse
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Docs    string `json:"docs"`
}

// CreateImageOptions contains options for creating an image
type CreateImageOptions struct {
	Modifications []Modification
	Format        string // "png" or "svg"
	Thumbnail     bool
}

// CreateImage generates an image from a template
func (c *Client) CreateImage(ctx context.Context, templateID string, opts *CreateImageOptions) *Response {
	if opts == nil {
		opts = &CreateImageOptions{}
	}

	payload := map[string]interface{}{
		"apiKey":        c.apiKey,
		"templateId":    templateID,
		"modifications": opts.Modifications,
		"format":        opts.Format,
		"thumbnail":     opts.Thumbnail,
	}

	if opts.Format == "" {
		payload["format"] = "png"
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "JSON_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/templates/createImage", bytes.NewReader(body))
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "REQUEST_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "HTTP_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "READ_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	if resp.StatusCode != http.StatusOK {
		return &Response{Error: &ErrorResponse{
			Code:    "HTTP_ERROR",
			Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(result)),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	return &Response{Result: result}
}

// CreatePDF generates a PDF from a template
func (c *Client) CreatePDF(ctx context.Context, templateID string, modifications []Modification) *Response {
	payload := map[string]interface{}{
		"apiKey":        c.apiKey,
		"templateId":    templateID,
		"modifications": modifications,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "JSON_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/templates/createPdf", bytes.NewReader(body))
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "REQUEST_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "HTTP_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{Error: &ErrorResponse{
			Code:    "READ_ERROR",
			Message: err.Error(),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	if resp.StatusCode != http.StatusOK {
		return &Response{Error: &ErrorResponse{
			Code:    "HTTP_ERROR",
			Message: fmt.Sprintf("HTTP %d", resp.StatusCode),
			Docs:    "https://bannerify.co/docs",
		}}
	}

	return &Response{Result: result}
}

// GenerateImageSignedURL generates a signed URL for on-demand image generation
func (c *Client) GenerateImageSignedURL(templateID string, opts *CreateImageOptions) (string, error) {
	if opts == nil {
		opts = &CreateImageOptions{}
	}

	apiKeyHashed := fmt.Sprintf("%x", sha256.Sum256([]byte(c.apiKey)))

	params := url.Values{}
	params.Set("apiKeyHashed", apiKeyHashed)
	params.Set("templateId", templateID)

	if opts.Format == "svg" {
		params.Set("format", "svg")
	}

	if len(opts.Modifications) > 0 {
		modsJSON, err := json.Marshal(opts.Modifications)
		if err != nil {
			return "", err
		}
		params.Set("modifications", string(modsJSON))
	}

	if opts.Thumbnail {
		params.Set("thumbnail", "true")
	}

	// Sort parameters
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedParams := url.Values{}
	for _, k := range keys {
		sortedParams.Set(k, params.Get(k))
	}

	// Generate signature
	queryString := sortedParams.Encode()
	signInput := queryString + apiKeyHashed
	sign := fmt.Sprintf("%x", sha256.Sum256([]byte(signInput)))
	sortedParams.Set("sign", sign)

	return fmt.Sprintf("%s/templates/signedurl?%s", c.baseURL, sortedParams.Encode()), nil
}

