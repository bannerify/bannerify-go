// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/bannerify-go"
	"github.com/stainless-sdks/bannerify-go/internal/testutil"
	"github.com/stainless-sdks/bannerify-go/option"
)

func TestKeyNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bannerify.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Keys.New(context.TODO(), bannerify.KeyNewParams{
		APIID:      bannerify.F("api_123"),
		ByteLength: bannerify.F(int64(16)),
		Expires:    bannerify.F(int64(1623869797161)),
		Meta: bannerify.F(map[string]interface{}{
			"billingTier": "PRO",
			"trialEnds":   "2023-06-16T17:16:37.161Z",
		}),
		Name:    bannerify.F("my key"),
		OwnerID: bannerify.F("team_123"),
		Prefix:  bannerify.F("string"),
		Ratelimit: bannerify.F(bannerify.KeyNewParamsRatelimit{
			Type:           bannerify.F(bannerify.KeyNewParamsRatelimitTypeFast),
			Limit:          bannerify.F(int64(10)),
			RefillRate:     bannerify.F(int64(1)),
			RefillInterval: bannerify.F(int64(60)),
		}),
		Remaining: bannerify.F(int64(1000)),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyVerifyWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bannerify.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Keys.Verify(context.TODO(), bannerify.KeyVerifyParams{
		Key:   bannerify.F("sk_1234"),
		APIID: bannerify.F("api_1234"),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
