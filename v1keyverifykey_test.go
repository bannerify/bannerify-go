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

func TestV1KeyVerifyKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.VerifyKey.New(context.TODO(), tempbannerify.V1KeyVerifyKeyNewParams{
		Key:   tempbannerify.F("sk_1234"),
		APIID: tempbannerify.F("api_1234"),
		Authorization: tempbannerify.F(tempbannerify.V1KeyVerifyKeyNewParamsAuthorization{
			Permissions: tempbannerify.F[any](map[string]interface{}{
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
	})
	if err != nil {
		var apierr *tempbannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
