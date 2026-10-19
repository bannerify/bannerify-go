package bannerify

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"testing"
)

type signedURLVectors struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseUrl"`
	Encoded []struct {
		Input   string `json:"input"`
		Encoded string `json:"encoded"`
	} `json:"encoded"`
	Golden []struct {
		Name          string `json:"name"`
		TemplateID    string `json:"templateId"`
		Format        string `json:"format"`
		Modifications []struct {
			Name string `json:"name"`
			Text string `json:"text"`
		} `json:"modifications"`
		URL string `json:"url"`
	} `json:"golden"`
}

// The vectors come from the JavaScript URLSearchParams, which the server uses to
// verify a signed URL.
func loadVectors(t *testing.T) signedURLVectors {
	t.Helper()

	raw, err := os.ReadFile("testdata/signed-url-vectors.json")
	if err != nil {
		t.Fatal(err)
	}

	var vectors signedURLVectors
	if err := json.Unmarshal(raw, &vectors); err != nil {
		t.Fatal(err)
	}

	return vectors
}

func queryPair(rawURL, key string) string {
	query := rawURL[strings.Index(rawURL, "?")+1:]
	for _, pair := range strings.Split(query, "&") {
		if strings.HasPrefix(pair, key+"=") {
			return pair
		}
	}

	return ""
}

func TestEncodeFormMatchesURLSearchParams(t *testing.T) {
	for _, vector := range loadVectors(t).Encoded {
		if got := encodeForm(vector.Input); got != vector.Encoded {
			t.Errorf("encodeForm(%q) = %q, want %q", vector.Input, got, vector.Encoded)
		}
	}
}

func TestGenerateImageSignedURLEncodesValuesLikeURLSearchParams(t *testing.T) {
	vectors := loadVectors(t)
	client := NewClient(vectors.APIKey, WithBaseURL(vectors.BaseURL))

	for _, vector := range vectors.Encoded {
		signedURL, err := client.GenerateImageSignedURL(vector.Input, nil)
		if err != nil {
			t.Fatal(err)
		}

		if got, want := queryPair(signedURL, "templateId"), "templateId="+vector.Encoded; got != want {
			t.Errorf("templateId for %q = %q, want %q", vector.Input, got, want)
		}
	}
}

func TestGenerateImageSignedURLMatchesGoldenURLs(t *testing.T) {
	vectors := loadVectors(t)
	client := NewClient(vectors.APIKey, WithBaseURL(vectors.BaseURL))

	for _, golden := range vectors.Golden {
		opts := &CreateImageOptions{Format: golden.Format}
		for _, modification := range golden.Modifications {
			text := modification.Text
			opts.Modifications = append(opts.Modifications, Modification{Name: modification.Name, Text: &text})
		}

		got, err := client.GenerateImageSignedURL(golden.TemplateID, opts)
		if err != nil {
			t.Fatal(err)
		}

		if got != golden.URL {
			t.Errorf("%s:\n got  %s\n want %s", golden.Name, got, golden.URL)
		}
	}
}

func TestGenerateImageSignedURLSignsTheSortedQuery(t *testing.T) {
	client := NewClient("test-key", WithBaseURL("https://api.example.test/v1"))
	text := "A*B~C"

	signedURL, err := client.GenerateImageSignedURL("tpl_x", &CreateImageOptions{
		Format:        "jpeg",
		Modifications: []Modification{{Name: "title", Text: &text}},
	})
	if err != nil {
		t.Fatal(err)
	}

	query := strings.SplitN(signedURL, "?", 2)[1]
	pairs := strings.Split(query, "&")
	last := len(pairs) - 1
	if !strings.HasPrefix(pairs[last], "sign=") {
		t.Fatalf("sign is not the last parameter: %s", query)
	}

	signed := pairs[:last]
	if !sort.StringsAreSorted(signed) {
		t.Errorf("parameters are not sorted by name: %v", signed)
	}

	mac := hmac.New(sha256.New, []byte("test-key"))
	mac.Write([]byte(strings.Join(signed, "&")))
	if want := hex.EncodeToString(mac.Sum(nil)); pairs[last] != "sign="+want {
		t.Errorf("sign = %s, want sign=%s", pairs[last], want)
	}
}

func TestGenerateImageSignedURLRejectsAnUnsupportedFormat(t *testing.T) {
	client := NewClient("test-key")

	if _, err := client.GenerateImageSignedURL("tpl_x", &CreateImageOptions{Format: "gif"}); err == nil {
		t.Fatal("expected an error for the gif format")
	}
}

func newTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	return NewClient("test-key", WithBaseURL(server.URL))
}

func TestCreateImageSendsNoThumbnail(t *testing.T) {
	var body map[string]interface{}
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/templates/createImage" {
			t.Errorf("path = %s", r.URL.Path)
		}
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &body)
		_, _ = w.Write([]byte("image"))
	})

	response := client.CreateImage(context.Background(), "tpl_x", &CreateImageOptions{Format: "webp"})

	if response.Error != nil {
		t.Fatalf("unexpected error: %v", response.Error)
	}
	if string(response.Result) != "image" {
		t.Errorf("result = %q", response.Result)
	}

	keys := make([]string, 0, len(body))
	for key := range body {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if got := strings.Join(keys, ","); got != "apiKey,format,modifications,templateId" {
		t.Errorf("request fields = %s", got)
	}
}

const badRequestBody = `{"error":{"code":"BAD_REQUEST","message":"Invalid api key","docs":"https://bannerify.co/docs/api-reference/errors/code/BAD_REQUEST","requestId":"req:abc123"}}`

func TestRequestsReturnTheAPIError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(badRequestBody))
	})

	responses := map[string]*Response{
		"CreateImage": client.CreateImage(context.Background(), "tpl_x", nil),
		"CreatePDF":   client.CreatePDF(context.Background(), "tpl_x", nil),
	}

	for name, response := range responses {
		if response.Error == nil {
			t.Fatalf("%s: expected an error", name)
		}

		want := ErrorResponse{
			Code:      "BAD_REQUEST",
			Message:   "Invalid api key",
			Docs:      "https://bannerify.co/docs/api-reference/errors/code/BAD_REQUEST",
			RequestID: "req:abc123",
		}
		if *response.Error != want {
			t.Errorf("%s: error = %+v, want %+v", name, *response.Error, want)
		}

		var asError error = response.Error
		var apiError *ErrorResponse
		if !errors.As(asError, &apiError) || apiError.Code != "BAD_REQUEST" {
			t.Errorf("%s: errors.As did not find the API error", name)
		}
		if got := asError.Error(); got != "BAD_REQUEST: Invalid api key (request req:abc123)" {
			t.Errorf("%s: Error() = %q", name, got)
		}
	}
}

func TestRequestsReturnHTTPErrorWhenTheBodyIsNotAnAPIError(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("Bad Gateway"))
	})

	response := client.CreatePDF(context.Background(), "tpl_x", nil)

	if response.Error == nil {
		t.Fatal("expected an error")
	}
	if response.Error.Code != "HTTP_ERROR" || response.Error.Message != "HTTP 502: Bad Gateway" {
		t.Errorf("error = %+v", *response.Error)
	}
}

func TestCreateImageRejectsAnUnsupportedFormatWithoutARequest(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("the client sent a request")
	})

	response := client.CreateImage(context.Background(), "tpl_x", &CreateImageOptions{Format: "gif"})

	if response.Error == nil || response.Error.Code != "INVALID_FORMAT" {
		t.Errorf("error = %+v", response.Error)
	}
}

func sseBody(events ...string) string {
	var body strings.Builder
	for _, event := range events {
		body.WriteString("data: " + event + "\n\n")
	}
	return body.String()
}

const doneEvent = `{"type":"done","url":"https://images.bannerify.co/files/a.mp4"}`

func TestRenderVideoStreamsTheEventsOfTheRender(t *testing.T) {
	var body map[string]interface{}
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/templates/renderVideo" || r.URL.Query().Get("stream") != "true" {
			t.Errorf("url = %s", r.URL)
		}
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &body)
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = w.Write([]byte(sseBody(
			`{"type":"queued","ahead":1}`,
			`{"type":"progress","frame":30,"totalFrames":60,"progress":0.5}`,
			doneEvent,
		)))
	})

	title := "Hi"
	var events []RenderVideoEvent
	for event := range client.RenderVideo(context.Background(), "tpl_x", []Modification{{Name: "Title", Text: &title}}) {
		events = append(events, event)
	}

	want := []RenderVideoEvent{
		{Type: "queued", Ahead: 1},
		{Type: "progress", Frame: 30, TotalFrames: 60, Progress: 0.5},
		{Type: "done", URL: "https://images.bannerify.co/files/a.mp4"},
	}
	if len(events) != len(want) {
		t.Fatalf("events = %+v", events)
	}
	for i := range want {
		if events[i] != want[i] {
			t.Errorf("event %d = %+v, want %+v", i, events[i], want[i])
		}
	}
	modifications, _ := json.Marshal(body["modifications"])
	if body["templateId"] != "tpl_x" || string(modifications) != `[{"name":"Title","text":"Hi"}]` {
		t.Errorf("body = %v", body)
	}
}

func TestRenderVideoTurnsARefusedRequestIntoOneErrorEvent(t *testing.T) {
	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPreconditionFailed)
		_, _ = w.Write([]byte(`{"error":{"code":"PRECONDITION_FAILED","message":"Video rendering is not available yet.","docs":"","requestId":"req_1"}}`))
	})

	var events []RenderVideoEvent
	for event := range client.RenderVideo(context.Background(), "tpl_x", nil) {
		events = append(events, event)
	}

	if len(events) != 1 || events[0].Type != "error" || events[0].Error.Code != "PRECONDITION_FAILED" || events[0].Error.RequestID != "req_1" {
		t.Errorf("events = %+v", events)
	}
}

func TestRenderAndWaitVideoReturnsTheURLOrTheError(t *testing.T) {
	waited := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(sseBody(`{"type":"queued","ahead":0}`, doneEvent)))
	})
	url, err := waited.RenderAndWaitVideo(context.Background(), "tpl_x", nil)
	if err != nil || url != "https://images.bannerify.co/files/a.mp4" {
		t.Errorf("url = %q, err = %v", url, err)
	}

	failed := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(sseBody(`{"type":"error","error":{"code":"INTERNAL_SERVER_ERROR","message":"Could not download a clip.","requestId":"req_2"}}`)))
	})
	_, err = failed.RenderAndWaitVideo(context.Background(), "tpl_x", nil)
	var apiError *ErrorResponse
	if !errors.As(err, &apiError) || apiError.Message != "Could not download a clip." {
		t.Errorf("err = %v", err)
	}
}
