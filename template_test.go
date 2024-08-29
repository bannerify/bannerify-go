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

	"github.com/bannerify/bannerify-go"
	"github.com/bannerify/bannerify-go/option"
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
		APIKey:     bannerify.F("apiKey"),
		TemplateID: bannerify.F("tpl_xxxxxxxxx"),
		Debug:      bannerify.F("_debug"),
		Format:     bannerify.F(bannerify.TemplateNewImageParamsFormatPng),
		Modifications: bannerify.F([]bannerify.TemplateNewImageParamsModification{{
			Name:    bannerify.F("Text 1"),
			Barcode: bannerify.F("1234567890"),
			Chart: bannerify.F(map[string]interface{}{
				"foo": "bar",
			}),
			Color:   bannerify.F("#FF0000"),
			Qrcode:  bannerify.F("Some text"),
			Src:     bannerify.F("https://example.com/image.jpg"),
			Star:    bannerify.F(5.000000),
			Text:    bannerify.F("Hello World"),
			Visible: bannerify.F(true),
		}, {
			Name:    bannerify.F("Text 1"),
			Barcode: bannerify.F("1234567890"),
			Chart: bannerify.F(map[string]interface{}{
				"foo": "bar",
			}),
			Color:   bannerify.F("#FF0000"),
			Qrcode:  bannerify.F("Some text"),
			Src:     bannerify.F("https://example.com/image.jpg"),
			Star:    bannerify.F(5.000000),
			Text:    bannerify.F("Hello World"),
			Visible: bannerify.F(true),
		}, {
			Name:    bannerify.F("Text 1"),
			Barcode: bannerify.F("1234567890"),
			Chart: bannerify.F(map[string]interface{}{
				"foo": "bar",
			}),
			Color:   bannerify.F("#FF0000"),
			Qrcode:  bannerify.F("Some text"),
			Src:     bannerify.F("https://example.com/image.jpg"),
			Star:    bannerify.F(5.000000),
			Text:    bannerify.F("Hello World"),
			Visible: bannerify.F(true),
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
		Sign:          bannerify.F("sign"),
		TemplateID:    bannerify.F("tpl_xxxxxxxxx"),
		Debug:         bannerify.F("_debug"),
		APIKeyHashed:  bannerify.F("apiKeyHashed"),
		APIKeyMd5:     bannerify.F("apiKeyMd5"),
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
