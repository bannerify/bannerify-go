// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempbannerify

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/TEMP_bannerify-go/internal/apijson"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/apiquery"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/param"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/requestconfig"
	"github.com/stainless-sdks/TEMP_bannerify-go/option"
)

// V1KeyGetKeyService contains methods and other services that help with
// interacting with the bannerify API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV1KeyGetKeyService] method
// instead.
type V1KeyGetKeyService struct {
	Options []option.RequestOption
}

// NewV1KeyGetKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1KeyGetKeyService(opts ...option.RequestOption) (r *V1KeyGetKeyService) {
	r = &V1KeyGetKeyService{}
	r.Options = opts
	return
}

func (r *V1KeyGetKeyService) Get(ctx context.Context, query V1KeyGetKeyGetParams, opts ...option.RequestOption) (res *V1KeyGetKeyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.getKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V1KeyGetKeyGetResponse struct {
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
	Ratelimit V1KeyGetKeyGetResponseRatelimit `json:"ratelimit"`
	// Unkey allows you to refill remaining verifications on a key on a regular
	// interval.
	Refill V1KeyGetKeyGetResponseRefill `json:"refill"`
	// The number of requests that can be made with this key before it becomes invalid.
	// If this field is null or undefined, the key has no request limit.
	Remaining float64 `json:"remaining"`
	// All roles this key belongs to
	Roles []string                   `json:"roles"`
	JSON  v1KeyGetKeyGetResponseJSON `json:"-"`
}

// v1KeyGetKeyGetResponseJSON contains the JSON metadata for the struct
// [V1KeyGetKeyGetResponse]
type v1KeyGetKeyGetResponseJSON struct {
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

func (r *V1KeyGetKeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyGetKeyGetResponseJSON) RawJSON() string {
	return r.raw
}

// Unkey comes with per-key ratelimiting out of the box.
type V1KeyGetKeyGetResponseRatelimit struct {
	// The total amount of burstable requests.
	Limit int64 `json:"limit,required"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval int64 `json:"refillInterval,required"`
	// How many tokens to refill during each refillInterval.
	RefillRate int64 `json:"refillRate,required"`
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
	// accurate.
	Type V1KeyGetKeyGetResponseRatelimitType `json:"type,required"`
	JSON v1KeyGetKeyGetResponseRatelimitJSON `json:"-"`
}

// v1KeyGetKeyGetResponseRatelimitJSON contains the JSON metadata for the struct
// [V1KeyGetKeyGetResponseRatelimit]
type v1KeyGetKeyGetResponseRatelimitJSON struct {
	Limit          apijson.Field
	RefillInterval apijson.Field
	RefillRate     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V1KeyGetKeyGetResponseRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyGetKeyGetResponseRatelimitJSON) RawJSON() string {
	return r.raw
}

// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
// accurate.
type V1KeyGetKeyGetResponseRatelimitType string

const (
	V1KeyGetKeyGetResponseRatelimitTypeFast       V1KeyGetKeyGetResponseRatelimitType = "fast"
	V1KeyGetKeyGetResponseRatelimitTypeConsistent V1KeyGetKeyGetResponseRatelimitType = "consistent"
)

func (r V1KeyGetKeyGetResponseRatelimitType) IsKnown() bool {
	switch r {
	case V1KeyGetKeyGetResponseRatelimitTypeFast, V1KeyGetKeyGetResponseRatelimitTypeConsistent:
		return true
	}
	return false
}

// Unkey allows you to refill remaining verifications on a key on a regular
// interval.
type V1KeyGetKeyGetResponseRefill struct {
	// Resets `remaining` to this value every interval.
	Amount int64 `json:"amount,required"`
	// Determines the rate at which verifications will be refilled.
	Interval V1KeyGetKeyGetResponseRefillInterval `json:"interval,required"`
	// The unix timestamp in miliseconds when the key was last refilled.
	LastRefillAt float64                          `json:"lastRefillAt"`
	JSON         v1KeyGetKeyGetResponseRefillJSON `json:"-"`
}

// v1KeyGetKeyGetResponseRefillJSON contains the JSON metadata for the struct
// [V1KeyGetKeyGetResponseRefill]
type v1KeyGetKeyGetResponseRefillJSON struct {
	Amount       apijson.Field
	Interval     apijson.Field
	LastRefillAt apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1KeyGetKeyGetResponseRefill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyGetKeyGetResponseRefillJSON) RawJSON() string {
	return r.raw
}

// Determines the rate at which verifications will be refilled.
type V1KeyGetKeyGetResponseRefillInterval string

const (
	V1KeyGetKeyGetResponseRefillIntervalDaily   V1KeyGetKeyGetResponseRefillInterval = "daily"
	V1KeyGetKeyGetResponseRefillIntervalMonthly V1KeyGetKeyGetResponseRefillInterval = "monthly"
)

func (r V1KeyGetKeyGetResponseRefillInterval) IsKnown() bool {
	switch r {
	case V1KeyGetKeyGetResponseRefillIntervalDaily, V1KeyGetKeyGetResponseRefillIntervalMonthly:
		return true
	}
	return false
}

type V1KeyGetKeyGetParams struct {
	// The id of the key to fetch
	KeyID param.Field[string] `query:"keyId,required"`
}

// URLQuery serializes [V1KeyGetKeyGetParams]'s query parameters as `url.Values`.
func (r V1KeyGetKeyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
