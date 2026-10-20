// Package bannerify provides a simple client for the Bannerify API
package bannerify

import (
	"bufio"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://api.bannerify.co/v1"
	userAgent      = "bannerify-go/0.1.0"
	docsURL        = "https://bannerify.co/docs"
)

var supportedFormats = map[string]struct{}{
	"png":  {},
	"jpeg": {},
	"webp": {},
}

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

// ErrorResponse is the error object that the API returns. The SDK also uses it
// for the errors it creates itself, such as INVALID_FORMAT.
type ErrorResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Docs      string `json:"docs"`
	RequestID string `json:"requestId"`
}

// Error implements the error interface.
func (e *ErrorResponse) Error() string {
	if e.RequestID == "" {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("%s: %s (request %s)", e.Code, e.Message, e.RequestID)
}

// CreateImageOptions contains options for creating an image
type CreateImageOptions struct {
	Modifications []Modification
	Format        string // "png", "jpeg", or "webp"
}

// CreateImage generates an image from a template
func (c *Client) CreateImage(ctx context.Context, templateID string, opts *CreateImageOptions) *Response {
	if opts == nil {
		opts = &CreateImageOptions{}
	}

	format, err := normalizeFormat(opts.Format)
	if err != nil {
		return sdkError("INVALID_FORMAT", err.Error())
	}

	return c.post(ctx, "/templates/createImage", map[string]interface{}{
		"apiKey":        c.apiKey,
		"templateId":    templateID,
		"modifications": opts.Modifications,
		"format":        format,
	})
}

// CreatePDF generates a PDF from a template
func (c *Client) CreatePDF(ctx context.Context, templateID string, modifications []Modification) *Response {
	return c.post(ctx, "/templates/createPdf", map[string]interface{}{
		"apiKey":        c.apiKey,
		"templateId":    templateID,
		"modifications": modifications,
	})
}

// GenerateImageSignedURL generates a signed URL for on-demand image generation
func (c *Client) GenerateImageSignedURL(templateID string, opts *CreateImageOptions) (string, error) {
	if opts == nil {
		opts = &CreateImageOptions{}
	}

	params := map[string]string{
		"apiKeyHashed": fmt.Sprintf("%x", sha256.Sum256([]byte(c.apiKey))),
		"templateId":   templateID,
	}

	if opts.Format != "" {
		format, err := normalizeFormat(opts.Format)
		if err != nil {
			return "", err
		}
		params["format"] = format
	}

	if len(opts.Modifications) > 0 {
		modsJSON, err := json.Marshal(opts.Modifications)
		if err != nil {
			return "", err
		}
		params["modifications"] = string(modsJSON)
	}

	// The signature is an HMAC keyed with the API key, so the URL can be shared
	// without exposing anything that can sign a different one. The server signs
	// the parameters sorted by name and encoded like URLSearchParams.
	query := encodeQuery(params)
	mac := hmac.New(sha256.New, []byte(c.apiKey))
	mac.Write([]byte(query))

	return fmt.Sprintf("%s/templates/signedurl?%s&sign=%s", c.baseURL, query, hex.EncodeToString(mac.Sum(nil))), nil
}

// RenderVideoEvent is one event of a video render. Type is "queued" (Ahead
// renders wait before it), "progress" (Frame of TotalFrames rendered, Progress
// from 0 to 1), "done" (URL of the MP4) or "error" (Error).
type RenderVideoEvent struct {
	Type        string         `json:"type"`
	Ahead       int            `json:"ahead,omitempty"`
	Frame       int            `json:"frame,omitempty"`
	TotalFrames int            `json:"totalFrames,omitempty"`
	Progress    float64        `json:"progress,omitempty"`
	URL         string         `json:"url,omitempty"`
	Error       *ErrorResponse `json:"error,omitempty"`
}

// RenderVideo renders a video template as an MP4 and streams the events of the
// render. The last event is "done" or "error". Each second of the video counts
// as one generation, and as two at 60 fps. Cancel ctx to stop following the
// render.
func (c *Client) RenderVideo(ctx context.Context, templateID string, modifications []Modification) iter.Seq[RenderVideoEvent] {
	return func(yield func(RenderVideoEvent) bool) {
		failed := func(response *Response) {
			yield(RenderVideoEvent{Type: "error", Error: response.Error})
		}

		if modifications == nil {
			modifications = []Modification{}
		}
		body, err := json.Marshal(map[string]interface{}{
			"apiKey":        c.apiKey,
			"templateId":    templateID,
			"modifications": modifications,
		})
		if err != nil {
			failed(sdkError("JSON_ERROR", err.Error()))
			return
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/templates/renderVideo?stream=true", bytes.NewReader(body))
		if err != nil {
			failed(sdkError("REQUEST_ERROR", err.Error()))
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
		req.Header.Set("User-Agent", userAgent)

		// A render can wait in line for minutes, so the client timeout does not
		// apply to the stream; ctx still ends it.
		streaming := *c.httpClient
		streaming.Timeout = 0
		resp, err := streaming.Do(req)
		if err != nil {
			failed(sdkError("HTTP_ERROR", err.Error()))
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			result, _ := io.ReadAll(resp.Body)
			yield(RenderVideoEvent{Type: "error", Error: parseAPIError(resp.StatusCode, result)})
			return
		}

		scanner := bufio.NewScanner(resp.Body)
		var data []string
		for scanner.Scan() {
			line := scanner.Text()
			if value, ok := strings.CutPrefix(line, "data:"); ok {
				data = append(data, strings.TrimSpace(value))
				continue
			}
			if line != "" || len(data) == 0 {
				continue
			}
			var event RenderVideoEvent
			if err := json.Unmarshal([]byte(strings.Join(data, "\n")), &event); err != nil {
				failed(sdkError("JSON_ERROR", err.Error()))
				return
			}
			data = nil
			if !yield(event) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			failed(sdkError("READ_ERROR", err.Error()))
		}
	}
}

// RenderAndWaitVideo renders a video template as an MP4 and returns its URL
// when it is stored. The error is an *ErrorResponse.
func (c *Client) RenderAndWaitVideo(ctx context.Context, templateID string, modifications []Modification) (string, error) {
	for event := range c.RenderVideo(ctx, templateID, modifications) {
		switch event.Type {
		case "done":
			return event.URL, nil
		case "error":
			return "", event.Error
		}
	}
	return "", sdkError("STREAM_ENDED", "the render ended before the video was done").Error
}

// post sends a JSON request. A 200 response gives its body in Result. Any other
// outcome gives an Error.
func (c *Client) post(ctx context.Context, path string, payload map[string]interface{}) *Response {
	body, err := json.Marshal(payload)
	if err != nil {
		return sdkError("JSON_ERROR", err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return sdkError("REQUEST_ERROR", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return sdkError("HTTP_ERROR", err.Error())
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return sdkError("READ_ERROR", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return &Response{Error: parseAPIError(resp.StatusCode, result)}
	}

	return &Response{Result: result}
}

// parseAPIError reads the error object of an API response. A body that holds no
// error object becomes an HTTP_ERROR with the status and the start of the body.
func parseAPIError(status int, body []byte) *ErrorResponse {
	var envelope struct {
		Error *ErrorResponse `json:"error"`
	}
	if err := json.Unmarshal(body, &envelope); err == nil && envelope.Error != nil && envelope.Error.Code != "" {
		return envelope.Error
	}

	snippet := string(body)
	if len(snippet) > 200 {
		snippet = strings.ToValidUTF8(snippet[:200], "")
	}

	return &ErrorResponse{
		Code:    "HTTP_ERROR",
		Message: fmt.Sprintf("HTTP %d: %s", status, snippet),
		Docs:    docsURL,
	}
}

func sdkError(code, message string) *Response {
	return &Response{Error: &ErrorResponse{Code: code, Message: message, Docs: docsURL}}
}

func normalizeFormat(format string) (string, error) {
	if format == "" {
		return "png", nil
	}
	if _, ok := supportedFormats[format]; ok {
		return format, nil
	}
	return "", fmt.Errorf("unsupported format %q. Valid formats: png, jpeg, webp", format)
}

// encodeQuery joins the parameters sorted by name, each key and value encoded
// like URLSearchParams.
func encodeQuery(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	pairs := make([]string, len(keys))
	for i, key := range keys {
		pairs[i] = encodeForm(key) + "=" + encodeForm(params[key])
	}

	return strings.Join(pairs, "&")
}

// encodeForm encodes a string like URLSearchParams (application/x-www-form-urlencoded):
// letters, digits, and * - . _ stay, a space becomes +, and every other byte becomes %XX.
// url.QueryEscape differs for * and ~, so it cannot sign a URL that the server verifies.
func encodeForm(value string) string {
	const hexDigits = "0123456789ABCDEF"

	var encoded strings.Builder
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9',
			c == '*', c == '-', c == '.', c == '_':
			encoded.WriteByte(c)
		case c == ' ':
			encoded.WriteByte('+')
		default:
			encoded.WriteByte('%')
			encoded.WriteByte(hexDigits[c>>4])
			encoded.WriteByte(hexDigits[c&15])
		}
	}

	return encoded.String()
}
