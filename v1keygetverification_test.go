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

func TestV1KeyGetVerificationListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.GetVerifications.List(context.TODO(), tempbannerify.V1KeyGetVerificationListParams{
		End:         tempbannerify.F(int64(1620000000000)),
		Granularity: tempbannerify.F(tempbannerify.V1KeyGetVerificationListParamsGranularityDay),
		KeyID:       tempbannerify.F("key_1234"),
		OwnerID:     tempbannerify.F("chronark"),
		Start:       tempbannerify.F(int64(1620000000000)),
	})
	if err != nil {
		var apierr *tempbannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
