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

func TestAPIKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIs.Keys.List(
		context.TODO(),
		"api_1234",
		tempbannerify.APIKeyListParams{
			Limit:   tempbannerify.F(int64(100)),
			Offset:  tempbannerify.F(0.000000),
			OwnerID: tempbannerify.F("x"),
		},
	)
	if err != nil {
		var apierr *tempbannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
