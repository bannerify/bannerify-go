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

func TestV1KeyVerifyKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.VerifyKey.New(context.TODO(), bannerify.V1KeyVerifyKeyNewParams{
		Key:   bannerify.F("sk_1234"),
		APIID: bannerify.F("api_1234"),
		Authorization: bannerify.F(bannerify.V1KeyVerifyKeyNewParamsAuthorization{
			Permissions: bannerify.F[any](map[string]interface{}{
				"or": map[string]interface{}{
					"0": map[string]interface{}{
						"and": map[string]interface{}{
							"0": "dns.record.read",
							"1": "dns.record.update",
						},
					},
					"1": "admin",
				},
			}),
		}),
		Ratelimit: bannerify.F(bannerify.V1KeyVerifyKeyNewParamsRatelimit{
			Cost: bannerify.F(int64(0)),
		}),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
