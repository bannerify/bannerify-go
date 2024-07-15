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

func TestV1KeyUpdateKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Keys.UpdateKey.New(context.TODO(), bannerify.V1KeyUpdateKeyNewParams{
		KeyID:   bannerify.F("key_123"),
		Enabled: bannerify.F(true),
		Expires: bannerify.F(0.000000),
		Meta: bannerify.F(map[string]interface{}{
			"roles": map[string]interface{}{
				"0": "admin",
				"1": "user",
			},
			"stripeCustomerId": "cus_1234",
		}),
		Name:    bannerify.F("Customer X"),
		OwnerID: bannerify.F("user_123"),
		Permissions: bannerify.F([]bannerify.V1KeyUpdateKeyNewParamsPermissionUnion{bannerify.V1KeyUpdateKeyNewParamsPermissionsObject{
			ID: bannerify.F("xxx"),
		}}),
		Ratelimit: bannerify.F(bannerify.V1KeyUpdateKeyNewParamsRatelimit{
			Type:           bannerify.F(bannerify.V1KeyUpdateKeyNewParamsRatelimitTypeFast),
			Async:          bannerify.F(true),
			Limit:          bannerify.F(int64(10)),
			RefillRate:     bannerify.F(int64(1)),
			RefillInterval: bannerify.F(int64(60)),
			Duration:       bannerify.F(int64(1)),
		}),
		Refill: bannerify.F(bannerify.V1KeyUpdateKeyNewParamsRefill{
			Interval: bannerify.F(bannerify.V1KeyUpdateKeyNewParamsRefillIntervalDaily),
			Amount:   bannerify.F(int64(100)),
		}),
		Remaining: bannerify.F(1000.000000),
		Roles: bannerify.F([]bannerify.V1KeyUpdateKeyNewParamsRoleUnion{bannerify.V1KeyUpdateKeyNewParamsRolesObject{
			ID: bannerify.F("xxx"),
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
