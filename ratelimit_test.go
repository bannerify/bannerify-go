// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempbannerify_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/TEMP_bannerify-go"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/shared"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/testutil"
	"github.com/stainless-sdks/TEMP_bannerify-go/option"
)

func TestRatelimitLimitWithOptionalParams(t *testing.T) {
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
	_, err := client.Ratelimits.Limit(context.TODO(), tempbannerify.RatelimitLimitParams{
		Duration:   tempbannerify.F(int64(60000)),
		Identifier: tempbannerify.F("user_123"),
		Limit:      tempbannerify.F(int64(0)),
		Async:      tempbannerify.F(true),
		Cost:       tempbannerify.F(int64(2)),
		Meta: tempbannerify.F(map[string]tempbannerify.RatelimitLimitParamsMetaUnion{
			"foo": shared.UnionString("string"),
		}),
		Namespace: tempbannerify.F("email.outbound"),
		Resources: tempbannerify.F([]tempbannerify.RatelimitLimitParamsResource{{
			Type: tempbannerify.F("project"),
			ID:   tempbannerify.F("p_123"),
			Name: tempbannerify.F("dub"),
			Meta: tempbannerify.F(map[string]tempbannerify.RatelimitLimitParamsResourcesMetaUnion{
				"foo": shared.UnionString("string"),
			}),
		}}),
	})
	if err != nil {
		var apierr *tempbannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
