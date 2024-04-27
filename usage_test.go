// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/bannerify-go"
	"github.com/stainless-sdks/bannerify-go/internal/testutil"
	"github.com/stainless-sdks/bannerify-go/option"
)

func TestUsage(t *testing.T) {
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
	v1LivenessResponse, err := client.V1.Liveness(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", v1LivenessResponse.Services)
}
