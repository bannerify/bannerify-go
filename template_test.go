// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stainless-sdks/bannerify-go"
	"github.com/stainless-sdks/bannerify-go/option"
)

func TestTemplateNewImageWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := bannerify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Templates.NewImage(context.TODO(), bannerify.TemplateNewImageParams{
		APIKey:     bannerify.F("key_xxxxxxxxx"),
		TemplateID: bannerify.F("tpl_xxxxxxxxx"),
		Debug:      bannerify.F("_debug"),
		Format:     bannerify.F(bannerify.TemplateNewImageParamsFormatPng),
		Modifications: bannerify.F([]bannerify.TemplateNewImageParamsModification{{
			Name:    bannerify.F("Text 1"),
			Color:   bannerify.F("#FF0000"),
			Src:     bannerify.F("https://example.com/image.jpg"),
			Text:    bannerify.F("Hello World"),
			Barcode: bannerify.F("1234567890"),
			Qrcode:  bannerify.F("Some text"),
			Chart: bannerify.F(map[string]interface{}{
				"foo": map[string]interface{}{},
			}),
			Visible: bannerify.F(true),
			Star:    bannerify.F(5.000000),
		}, {
			Name:    bannerify.F("Text 1"),
			Color:   bannerify.F("#FF0000"),
			Src:     bannerify.F("https://example.com/image.jpg"),
			Text:    bannerify.F("Hello World"),
			Barcode: bannerify.F("1234567890"),
			Qrcode:  bannerify.F("Some text"),
			Chart: bannerify.F(map[string]interface{}{
				"foo": map[string]interface{}{},
			}),
			Visible: bannerify.F(true),
			Star:    bannerify.F(5.000000),
		}, {
			Name:    bannerify.F("Text 1"),
			Color:   bannerify.F("#FF0000"),
			Src:     bannerify.F("https://example.com/image.jpg"),
			Text:    bannerify.F("Hello World"),
			Barcode: bannerify.F("1234567890"),
			Qrcode:  bannerify.F("Some text"),
			Chart: bannerify.F(map[string]interface{}{
				"foo": map[string]interface{}{},
			}),
			Visible: bannerify.F(true),
			Star:    bannerify.F(5.000000),
		}}),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestTemplateSignedurlWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := bannerify.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Templates.Signedurl(context.TODO(), bannerify.TemplateSignedurlParams{
		APIKeyMd5:     bannerify.F("apiKeyMd5"),
		Sign:          bannerify.F("sign"),
		TemplateID:    bannerify.F("tpl_xxxxxxxxx"),
		Debug:         bannerify.F("_debug"),
		Format:        bannerify.F(bannerify.TemplateSignedurlParamsFormatPng),
		Modifications: bannerify.F("modifications"),
		Nocache:       bannerify.F("true"),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
