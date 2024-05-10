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

// V1KeyCreateKeyService contains methods and other services that help with
// interacting with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1KeyCreateKeyService] method instead.
type V1KeyCreateKeyService struct {
	Options []option.RequestOption
}

// NewV1KeyCreateKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1KeyCreateKeyService(opts ...option.RequestOption) (r *V1KeyCreateKeyService) {
	r = &V1KeyCreateKeyService{}
	r.Options = opts
	return
}

func (r *V1KeyCreateKeyService) New(ctx context.Context, body V1KeyCreateKeyNewParams, opts ...option.RequestOption) (res *V1KeyCreateKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.createKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1KeyCreateKeyNewResponse struct {
	// The newly created api key, do not store this on your own system but pass it
	// along to your user.
	Key string `json:"key,required"`
	// The id of the key. This is not a secret and can be stored as a reference if you
	// wish. You need the keyId to update or delete a key later.
	KeyID string                        `json:"keyId,required"`
	JSON  v1KeyCreateKeyNewResponseJSON `json:"-"`
}

// v1KeyCreateKeyNewResponseJSON contains the JSON metadata for the struct
// [V1KeyCreateKeyNewResponse]
type v1KeyCreateKeyNewResponseJSON struct {
	Key         apijson.Field
	KeyID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyCreateKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyCreateKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1KeyCreateKeyNewParams struct {
	// Choose an `API` where this key should be created.
	APIID param.Field[string] `json:"apiId,required"`
	// The byte length used to generate your key determines its entropy as well as its
	// length. Higher is better, but keys become longer and more annoying to handle.
	// The default is 16 bytes, or 2^^128 possible combinations.
	ByteLength param.Field[int64] `json:"byteLength"`
	// Sets if key is enabled or disabled. Disabled keys are not valid.
	Enabled param.Field[bool] `json:"enabled"`
	// Environments allow you to divide your keyspace.
	//
	// Some applications like Stripe, Clerk, WorkOS and others have a concept of "live"
	// and "test" keys to give the developer a way to develop their own application
	// without the risk of modifying real world resources.
	//
	// When you set an environment, we will return it back to you when validating the
	// key, so you can handle it correctly.
	Environment param.Field[string] `json:"environment"`
	// You can auto expire keys by providing a unix timestamp in milliseconds. Once
	// Keys expire they will automatically be disabled and are no longer valid unless
	// you enable them again.
	Expires param.Field[int64] `json:"expires"`
	// This is a place for dynamic meta data, anything that feels useful for you should
	// go here
	Meta param.Field[map[string]interface{}] `json:"meta"`
	// The name for your Key. This is not customer facing.
	Name param.Field[string] `json:"name"`
	// Your user’s Id. This will provide a link between Unkey and your customer record.
	// When validating a key, we will return this back to you, so you can clearly
	// identify your user from their api key.
	OwnerID param.Field[string] `json:"ownerId"`
	// To make it easier for your users to understand which product an api key belongs
	// to, you can add prefix them.
	//
	// For example Stripe famously prefixes their customer ids with cus* or their api
	// keys with sk_live*.
	//
	// The underscore is automatically added if you are defining a prefix, for example:
	// "prefix": "abc" will result in a key like abc_xxxxxxxxx
	Prefix param.Field[string] `json:"prefix"`
	// Unkey comes with per-key ratelimiting out of the box.
	Ratelimit param.Field[V1KeyCreateKeyNewParamsRatelimit] `json:"ratelimit"`
	// Unkey enables you to refill verifications for each key at regular intervals.
	Refill param.Field[V1KeyCreateKeyNewParamsRefill] `json:"refill"`
	// You can limit the number of requests a key can make. Once a key reaches 0
	// remaining requests, it will automatically be disabled and is no longer valid
	// unless you update it.
	Remaining param.Field[int64] `json:"remaining"`
	// A list of roles that this key should have. If the role does not exist, an error
	// is thrown
	Roles param.Field[[]string] `json:"roles"`
}

func (r V1KeyCreateKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Unkey comes with per-key ratelimiting out of the box.
type V1KeyCreateKeyNewParamsRatelimit struct {
	// The total amount of burstable requests.
	Limit param.Field[int64] `json:"limit,required"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval param.Field[int64] `json:"refillInterval,required"`
	// How many tokens to refill during each refillInterval.
	RefillRate param.Field[int64] `json:"refillRate,required"`
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
	// accurate.
	Type param.Field[V1KeyCreateKeyNewParamsRatelimitType] `json:"type"`
}

func (r V1KeyCreateKeyNewParamsRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
// accurate.
type V1KeyCreateKeyNewParamsRatelimitType string

const (
	V1KeyCreateKeyNewParamsRatelimitTypeFast       V1KeyCreateKeyNewParamsRatelimitType = "fast"
	V1KeyCreateKeyNewParamsRatelimitTypeConsistent V1KeyCreateKeyNewParamsRatelimitType = "consistent"
)

func (r V1KeyCreateKeyNewParamsRatelimitType) IsKnown() bool {
	switch r {
	case V1KeyCreateKeyNewParamsRatelimitTypeFast, V1KeyCreateKeyNewParamsRatelimitTypeConsistent:
		return true
	}
	return false
}

// Unkey enables you to refill verifications for each key at regular intervals.
type V1KeyCreateKeyNewParamsRefill struct {
	// The number of verifications to refill for each occurrence is determined
	// individually for each key.
	Amount param.Field[int64] `json:"amount,required"`
	// Unkey will automatically refill verifications at the set interval.
	Interval param.Field[V1KeyCreateKeyNewParamsRefillInterval] `json:"interval,required"`
}

func (r V1KeyCreateKeyNewParamsRefill) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Unkey will automatically refill verifications at the set interval.
type V1KeyCreateKeyNewParamsRefillInterval string

const (
	V1KeyCreateKeyNewParamsRefillIntervalDaily   V1KeyCreateKeyNewParamsRefillInterval = "daily"
	V1KeyCreateKeyNewParamsRefillIntervalMonthly V1KeyCreateKeyNewParamsRefillInterval = "monthly"
)

func (r V1KeyCreateKeyNewParamsRefillInterval) IsKnown() bool {
	switch r {
	case V1KeyCreateKeyNewParamsRefillIntervalDaily, V1KeyCreateKeyNewParamsRefillIntervalMonthly:
		return true
	}
	return false
}
