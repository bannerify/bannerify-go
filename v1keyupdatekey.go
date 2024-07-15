// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/bannerify-go/internal/apijson"
	"github.com/stainless-sdks/bannerify-go/internal/param"
	"github.com/stainless-sdks/bannerify-go/internal/requestconfig"
	"github.com/stainless-sdks/bannerify-go/option"
)

// V1KeyUpdateKeyService contains methods and other services that help with
// interacting with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1KeyUpdateKeyService] method instead.
type V1KeyUpdateKeyService struct {
	Options []option.RequestOption
}

// NewV1KeyUpdateKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1KeyUpdateKeyService(opts ...option.RequestOption) (r *V1KeyUpdateKeyService) {
	r = &V1KeyUpdateKeyService{}
	r.Options = opts
	return
}

func (r *V1KeyUpdateKeyService) New(ctx context.Context, body V1KeyUpdateKeyNewParams, opts ...option.RequestOption) (res *V1KeyUpdateKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.updateKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1KeyUpdateKeyNewResponse = interface{}

type V1KeyUpdateKeyNewParams struct {
	// The id of the key you want to modify
	KeyID param.Field[string] `json:"keyId,required"`
	// Set if key is enabled or disabled. If disabled, the key cannot be used to
	// verify.
	Enabled param.Field[bool] `json:"enabled"`
	// The unix timestamp in milliseconds when the key will expire. If this field is
	// null or undefined, the key is not expiring.
	Expires param.Field[float64] `json:"expires"`
	// Any additional metadata you want to store with the key
	Meta param.Field[map[string]interface{}] `json:"meta"`
	// The name of the key
	Name param.Field[string] `json:"name"`
	// The id of the tenant associated with this key. Use whatever reference you have
	// in your system to identify the tenant. When verifying the key, we will send this
	// field back to you, so you know who is accessing your API.
	OwnerID param.Field[string] `json:"ownerId"`
	// The permissions you want to set for this key. This overwrites all existing
	// permissions. Setting permissions requires the `rbac.*.add_permission_to_key`
	// permission.
	Permissions param.Field[[]V1KeyUpdateKeyNewParamsPermission] `json:"permissions"`
	// Unkey comes with per-key ratelimiting out of the box. Set `null` to disable.
	Ratelimit param.Field[V1KeyUpdateKeyNewParamsRatelimit] `json:"ratelimit"`
	// Unkey enables you to refill verifications for each key at regular intervals.
	Refill param.Field[V1KeyUpdateKeyNewParamsRefill] `json:"refill"`
	// The number of requests that can be made with this key before it becomes invalid.
	// Set `null` to disable.
	Remaining param.Field[float64] `json:"remaining"`
	// The roles you want to set for this key. This overwrites all existing roles.
	// Setting roles requires the `rbac.*.add_role_to_key` permission.
	Roles param.Field[[]V1KeyUpdateKeyNewParamsRole] `json:"roles"`
}

func (r V1KeyUpdateKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1KeyUpdateKeyNewParamsPermission struct {
	// The id of the permission. Provide either `id` or `name`. If both are provided
	// `id` is used.
	ID param.Field[string] `json:"id"`
	// Set to true to automatically create the permissions they do not exist yet. Only
	// works when specifying `name`. Autocreating permissions requires your root key to
	// have the `rbac.*.create_permission` permission, otherwise the request will get
	// rejected
	Create param.Field[bool] `json:"create"`
	// Identify the permission via its name. Provide either `id` or `name`. If both are
	// provided `id` is used.
	Name param.Field[string] `json:"name"`
}

func (r V1KeyUpdateKeyNewParamsPermission) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Unkey comes with per-key ratelimiting out of the box. Set `null` to disable.
type V1KeyUpdateKeyNewParamsRatelimit struct {
	// The total amount of requests allowed in a single window.
	Limit param.Field[int64] `json:"limit,required"`
	// Asnyc ratelimiting doesn't add latency, while sync ratelimiting is slightly more
	// accurate.
	Async param.Field[bool] `json:"async"`
	// The duration of each ratelimit window, in milliseconds. This field will become
	// required in a future version.
	Duration param.Field[int64] `json:"duration"`
	// Determines the speed at which tokens are refilled, in milliseconds. Deprecated,
	// use 'duration'
	RefillInterval param.Field[int64] `json:"refillInterval"`
	// How many tokens to refill during each refillInterval. Deprecated, use 'limit'
	// instead.
	RefillRate param.Field[int64] `json:"refillRate"`
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
	// accurate. Deprecated, use 'async' instead
	Type param.Field[V1KeyUpdateKeyNewParamsRatelimitType] `json:"type"`
}

func (r V1KeyUpdateKeyNewParamsRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
// accurate. Deprecated, use 'async' instead
type V1KeyUpdateKeyNewParamsRatelimitType string

const (
	V1KeyUpdateKeyNewParamsRatelimitTypeFast       V1KeyUpdateKeyNewParamsRatelimitType = "fast"
	V1KeyUpdateKeyNewParamsRatelimitTypeConsistent V1KeyUpdateKeyNewParamsRatelimitType = "consistent"
)

func (r V1KeyUpdateKeyNewParamsRatelimitType) IsKnown() bool {
	switch r {
	case V1KeyUpdateKeyNewParamsRatelimitTypeFast, V1KeyUpdateKeyNewParamsRatelimitTypeConsistent:
		return true
	}
	return false
}

// Unkey enables you to refill verifications for each key at regular intervals.
type V1KeyUpdateKeyNewParamsRefill struct {
	// The amount of verifications to refill for each occurrence is determined
	// individually for each key.
	Amount param.Field[int64] `json:"amount,required"`
	// Unkey will automatically refill verifications at the set interval. If null is
	// used the refill functionality will be removed from the key.
	Interval param.Field[V1KeyUpdateKeyNewParamsRefillInterval] `json:"interval,required"`
}

func (r V1KeyUpdateKeyNewParamsRefill) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Unkey will automatically refill verifications at the set interval. If null is
// used the refill functionality will be removed from the key.
type V1KeyUpdateKeyNewParamsRefillInterval string

const (
	V1KeyUpdateKeyNewParamsRefillIntervalDaily   V1KeyUpdateKeyNewParamsRefillInterval = "daily"
	V1KeyUpdateKeyNewParamsRefillIntervalMonthly V1KeyUpdateKeyNewParamsRefillInterval = "monthly"
)

func (r V1KeyUpdateKeyNewParamsRefillInterval) IsKnown() bool {
	switch r {
	case V1KeyUpdateKeyNewParamsRefillIntervalDaily, V1KeyUpdateKeyNewParamsRefillIntervalMonthly:
		return true
	}
	return false
}

type V1KeyUpdateKeyNewParamsRole struct {
	// The id of the role. Provide either `id` or `name`. If both are provided `id` is
	// used.
	ID param.Field[string] `json:"id"`
	// Set to true to automatically create the permissions they do not exist yet. Only
	// works when specifying `name`. Autocreating roles requires your root key to have
	// the `rbac.*.create_role` permission, otherwise the request will get rejected
	Create param.Field[bool] `json:"create"`
	// Identify the role via its name. Provide either `id` or `name`. If both are
	// provided `id` is used.
	Name param.Field[string] `json:"name"`
}

func (r V1KeyUpdateKeyNewParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
