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
