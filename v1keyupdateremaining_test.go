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
)

func TestV1KeyUpdateRemainingNew(t *testing.T) {
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
	_, err := client.V1.Keys.UpdateRemaining.New(context.TODO(), bannerify.V1KeyUpdateRemainingNewParams{
		KeyID: bannerify.F("key_123"),
		Op:    bannerify.F(bannerify.V1KeyUpdateRemainingNewParamsOpIncrement),
		Value: bannerify.F(int64(1)),
	})
	if err != nil {
		var apierr *bannerify.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
