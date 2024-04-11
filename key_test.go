// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempbannerify_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/TEMP_bannerify-go"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/testutil"
	"github.com/stainless-sdks/TEMP_bannerify-go/option"
)

func TestKeyNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := tempbannerify.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Keys.New(context.TODO(), tempbannerify.KeyNewParams{
		APIID:      tempbannerify.F("api_123"),
		ByteLength: tempbannerify.F(int64(16)),
		Expires:    tempbannerify.F(int64(1623869797161)),
		Meta: tempbannerify.F(map[string]interface{}{
			"billingTier": "PRO",
			"trialEnds":   "2023-06-16T17:16:37.161Z",
		}),
		Name:    tempbannerify.F("my key"),
		OwnerID: tempbannerify.F("team_123"),
		Prefix:  tempbannerify.F("string"),
		Ratelimit: tempbannerify.F(tempbannerify.KeyNewParamsRatelimit{
			Type:           tempbannerify.F(tempbannerify.KeyNewParamsRatelimitTypeFast),
			Limit:          tempbannerify.F(int64(10)),
			RefillRate:     tempbannerify.F(int64(1)),
			RefillInterval: tempbannerify.F(int64(60)),
		}),
		Remaining: tempbannerify.F(int64(1000)),
	})
	if err != nil {
		var apierr *tempbannerify.Error
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
	client := tempbannerify.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Keys.Verify(context.TODO(), tempbannerify.KeyVerifyParams{
		Key:   tempbannerify.F("sk_1234"),
		APIID: tempbannerify.F("api_1234"),
	})
	if err != nil {
		var apierr *tempbannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
