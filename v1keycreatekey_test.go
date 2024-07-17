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

func TestV1KeyCreateKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.CreateKey.New(context.TODO(), bannerify.V1KeyCreateKeyNewParams{
		APIID:       bannerify.F("api_123"),
		ByteLength:  bannerify.F(int64(16)),
		Enabled:     bannerify.F(false),
		Environment: bannerify.F("environment"),
		Expires:     bannerify.F(int64(1623869797161)),
		ExternalID:  bannerify.F("team_123"),
		Meta: bannerify.F(map[string]interface{}{
			"billingTier": "PRO",
			"trialEnds":   "2023-06-16T17:16:37.161Z",
		}),
		Name:        bannerify.F("my key"),
		OwnerID:     bannerify.F("team_123"),
		Permissions: bannerify.F([]string{"domains.create_record", "say_hello"}),
		Prefix:      bannerify.F("prefix"),
		Ratelimit: bannerify.F(bannerify.V1KeyCreateKeyNewParamsRatelimit{
			Async:          bannerify.F(true),
			Type:           bannerify.F(bannerify.V1KeyCreateKeyNewParamsRatelimitTypeFast),
			Limit:          bannerify.F(int64(10)),
			Duration:       bannerify.F(int64(60000)),
			RefillRate:     bannerify.F(int64(1)),
			RefillInterval: bannerify.F(int64(1)),
		}),
		Refill: bannerify.F(bannerify.V1KeyCreateKeyNewParamsRefill{
			Interval: bannerify.F(bannerify.V1KeyCreateKeyNewParamsRefillIntervalDaily),
			Amount:   bannerify.F(int64(0)),
		}),
		Remaining: bannerify.F(int64(1000)),
		Roles:     bannerify.F([]string{"admin", "finance"}),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
