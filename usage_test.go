// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify_test

import (
	"context"
	"os"
	"testing"

	"github.com/bannerify/bannerify-go"
	"github.com/bannerify/bannerify-go/internal/testutil"
	"github.com/bannerify/bannerify-go/option"
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
		option.WithAPIKey("My API Key"),
	)
	infoGetResponse, err := client.Info.Get(context.TODO(), bannerify.InfoGetParams{
		APIKey: bannerify.F("REPLACE_ME"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", infoGetResponse.ID)
}
