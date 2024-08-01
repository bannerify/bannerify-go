// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/bannerify/bannerify-go"
	"github.com/bannerify/bannerify-go/internal/testutil"
	"github.com/bannerify/bannerify-go/option"
	"github.com/bannerify/bannerify-go/shared"
)

func TestRatelimitLimitWithOptionalParams(t *testing.T) {
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
	_, err := client.Ratelimits.Limit(context.TODO(), bannerify.RatelimitLimitParams{
		Duration:   bannerify.F(int64(60000)),
		Identifier: bannerify.F("user_123"),
		Limit:      bannerify.F(int64(0)),
		Async:      bannerify.F(true),
		Cost:       bannerify.F(int64(2)),
		Meta: bannerify.F(map[string]bannerify.RatelimitLimitParamsMetaUnion{
			"foo": shared.UnionString("string"),
		}),
		Namespace: bannerify.F("email.outbound"),
		Resources: bannerify.F([]bannerify.RatelimitLimitParamsResource{{
			Type: bannerify.F("project"),
			ID:   bannerify.F("p_123"),
			Name: bannerify.F("dub"),
			Meta: bannerify.F(map[string]bannerify.RatelimitLimitParamsResourcesMetaUnion{
				"foo": shared.UnionString("string"),
			}),
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
