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

func TestV1KeyUpdateKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.UpdateKey.New(context.TODO(), tempbannerify.V1KeyUpdateKeyNewParams{
		KeyID:   tempbannerify.F("key_123"),
		Enabled: tempbannerify.F(true),
		Expires: tempbannerify.F(0.000000),
		Meta: tempbannerify.F(map[string]interface{}{
			"roles": map[string]interface{}{
				"0": "admin",
				"1": "user",
			},
			"stripeCustomerId": "cus_1234",
		}),
		Name:    tempbannerify.F("Customer X"),
		OwnerID: tempbannerify.F("user_123"),
		Ratelimit: tempbannerify.F(tempbannerify.V1KeyUpdateKeyNewParamsRatelimit{
			Type:           tempbannerify.F(tempbannerify.V1KeyUpdateKeyNewParamsRatelimitTypeFast),
			Limit:          tempbannerify.F(int64(10)),
			RefillRate:     tempbannerify.F(int64(1)),
			RefillInterval: tempbannerify.F(int64(60)),
		}),
		Refill: tempbannerify.F(tempbannerify.V1KeyUpdateKeyNewParamsRefill{
			Interval: tempbannerify.F(tempbannerify.V1KeyUpdateKeyNewParamsRefillIntervalDaily),
			Amount:   tempbannerify.F(int64(100)),
		}),
		Remaining: tempbannerify.F(1000.000000),
	})
	if err != nil {
		var apierr *tempbannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
