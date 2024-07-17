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
	"github.com/stainless-sdks/bannerify-go/shared"
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
			Permissions: bannerify.F[bannerify.V1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion](shared.UnionString("string")),
		}),
		Ratelimit: bannerify.F(bannerify.V1KeyVerifyKeyNewParamsRatelimit{
			Cost: bannerify.F(int64(0)),
		}),
		Ratelimits: bannerify.F([]bannerify.V1KeyVerifyKeyNewParamsRatelimit{{
			Name:     bannerify.F("requests"),
			Cost:     bannerify.F(int64(0)),
			Limit:    bannerify.F(500.000000),
			Duration: bannerify.F(3600000.000000),
		}, {
			Name:     bannerify.F("tokens"),
			Cost:     bannerify.F(int64(0)),
			Limit:    bannerify.F(20000.000000),
			Duration: bannerify.F(86400000.000000),
		}}),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
