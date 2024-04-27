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

func TestV1KeyGetVerificationListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.GetVerifications.List(context.TODO(), bannerify.V1KeyGetVerificationListParams{
		End:         bannerify.F(int64(1620000000000)),
		Granularity: bannerify.F(bannerify.V1KeyGetVerificationListParamsGranularityDay),
		KeyID:       bannerify.F("key_1234"),
		OwnerID:     bannerify.F("chronark"),
		Start:       bannerify.F(int64(1620000000000)),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
