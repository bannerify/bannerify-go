// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/bannerify-go/internal/apijson"
	"github.com/stainless-sdks/bannerify-go/internal/apiquery"
	"github.com/stainless-sdks/bannerify-go/internal/param"
	"github.com/stainless-sdks/bannerify-go/internal/requestconfig"
	"github.com/stainless-sdks/bannerify-go/option"
)

// APIKeyService contains methods and other services that help with interacting
// with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r *APIKeyService) {
	r = &APIKeyService{}
	r.Options = opts
	return
}

func (r *APIKeyService) List(ctx context.Context, apiID string, query APIKeyListParams, opts ...option.RequestOption) (res *APIKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v1/apis/%s/keys", apiID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIKeyListResponse struct {
	Keys []APIKeyListResponseKey `json:"keys,required"`
	// The total number of keys for this api
	Total int64                  `json:"total,required"`
	JSON  apiKeyListResponseJSON `json:"-"`
}

// apiKeyListResponseJSON contains the JSON metadata for the struct
// [APIKeyListResponse]
type apiKeyListResponseJSON struct {
	Keys        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type APIKeyListResponseKey struct {
	// The id of the key
	ID string `json:"id,required"`
	// The first few characters of the key to visually identify it
	Start string `json:"start,required"`
	// The id of the workspace that owns the key
	WorkspaceID string `json:"workspaceId,required"`
	// The id of the api that this key is for
	APIID string `json:"apiId"`
	// The unix timestamp in milliseconds when the key was created
	CreatedAt float64 `json:"createdAt"`
	// The unix timestamp in milliseconds when the key was deleted. We don't delete the
	// key outright, you can restore it later.
	DeletedAt float64 `json:"deletedAt"`
	// Sets if key is enabled or disabled. Disabled keys are not valid.
	Enabled bool `json:"enabled"`
	// The unix timestamp in milliseconds when the key will expire. If this field is
	// null or undefined, the key is not expiring.
	Expires float64 `json:"expires"`
	// Any additional metadata you want to store with the key
	Meta map[string]interface{} `json:"meta"`
	// The name of the key, give keys a name to easily identify their purpose
	Name string `json:"name"`
	// The id of the tenant associated with this key. Use whatever reference you have
	// in your system to identify the tenant. When verifying the key, we will send this
	// field back to you, so you know who is accessing your API.
	OwnerID string `json:"ownerId"`
	// All permissions this key has
	Permissions []string `json:"permissions"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit APIKeyListResponseKeysRatelimit `json:"ratelimit"`
	// Unkey allows you to refill remaining verifications on a key on a regular
	// interval.
	Refill APIKeyListResponseKeysRefill `json:"refill"`
	// The number of requests that can be made with this key before it becomes invalid.
	// If this field is null or undefined, the key has no request limit.
	Remaining float64 `json:"remaining"`
	// All roles this key belongs to
	Roles []string                  `json:"roles"`
	JSON  apiKeyListResponseKeyJSON `json:"-"`
}

// apiKeyListResponseKeyJSON contains the JSON metadata for the struct
// [APIKeyListResponseKey]
type apiKeyListResponseKeyJSON struct {
	ID          apijson.Field
	Start       apijson.Field
	WorkspaceID apijson.Field
	APIID       apijson.Field
	CreatedAt   apijson.Field
	DeletedAt   apijson.Field
	Enabled     apijson.Field
	Expires     apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	OwnerID     apijson.Field
	Permissions apijson.Field
	Ratelimit   apijson.Field
	Refill      apijson.Field
	Remaining   apijson.Field
	Roles       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeyListResponseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseKeyJSON) RawJSON() string {
	return r.raw
}

// Unkey comes with per-key ratelimiting out of the box.
type APIKeyListResponseKeysRatelimit struct {
	// The total amount of burstable requests.
	Limit int64 `json:"limit,required"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval,required"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate,required"`
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
	// accurate.
	Type APIKeyListResponseKeysRatelimitType `json:"type"`
	JSON apiKeyListResponseKeysRatelimitJSON `json:"-"`
}

// apiKeyListResponseKeysRatelimitJSON contains the JSON metadata for the struct
// [APIKeyListResponseKeysRatelimit]
type apiKeyListResponseKeysRatelimitJSON struct {
	Limit          apijson.Field
	RefillInterval apijson.Field
	RefillRate     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *APIKeyListResponseKeysRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseKeysRatelimitJSON) RawJSON() string {
	return r.raw
}

// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
// accurate.
type APIKeyListResponseKeysRatelimitType string

const (
	APIKeyListResponseKeysRatelimitTypeFast       APIKeyListResponseKeysRatelimitType = "fast"
	APIKeyListResponseKeysRatelimitTypeConsistent APIKeyListResponseKeysRatelimitType = "consistent"
)

func (r APIKeyListResponseKeysRatelimitType) IsKnown() bool {
	switch r {
	case APIKeyListResponseKeysRatelimitTypeFast, APIKeyListResponseKeysRatelimitTypeConsistent:
		return true
	}
	return false
}

// Unkey allows you to refill remaining verifications on a key on a regular
// interval.
type APIKeyListResponseKeysRefill struct {
	// Resets `remaining` to this value every interval.
	Amount int64 `json:"amount,required"`
	// Determines the rate at which verifications will be refilled.
	Interval APIKeyListResponseKeysRefillInterval `json:"interval,required"`
	// The unix timestamp in miliseconds when the key was last refilled.
	LastRefillAt float64                          `json:"lastRefillAt"`
	JSON         apiKeyListResponseKeysRefillJSON `json:"-"`
}

// apiKeyListResponseKeysRefillJSON contains the JSON metadata for the struct
// [APIKeyListResponseKeysRefill]
type apiKeyListResponseKeysRefillJSON struct {
	Amount       apijson.Field
	Interval     apijson.Field
	LastRefillAt apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *APIKeyListResponseKeysRefill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyListResponseKeysRefillJSON) RawJSON() string {
	return r.raw
}

// Determines the rate at which verifications will be refilled.
type APIKeyListResponseKeysRefillInterval string

const (
	APIKeyListResponseKeysRefillIntervalDaily   APIKeyListResponseKeysRefillInterval = "daily"
	APIKeyListResponseKeysRefillIntervalMonthly APIKeyListResponseKeysRefillInterval = "monthly"
)

func (r APIKeyListResponseKeysRefillInterval) IsKnown() bool {
	switch r {
	case APIKeyListResponseKeysRefillIntervalDaily, APIKeyListResponseKeysRefillIntervalMonthly:
		return true
	}
	return false
}

type APIKeyListParams struct {
	// The maximum number of keys to return
	Limit param.Field[int64] `query:"limit"`
	// Use this to fetch the next page of results. A new cursor will be returned in the
	// response if there are more results.
	Offset param.Field[float64] `query:"offset"`
	// If provided, this will only return keys where the `ownerId` matches.
	OwnerID param.Field[string] `query:"ownerId"`
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
