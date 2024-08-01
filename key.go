// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"context"
	"net/http"

	"github.com/bannerify/bannerify-go/internal/apijson"
	"github.com/bannerify/bannerify-go/internal/param"
	"github.com/bannerify/bannerify-go/internal/requestconfig"
	"github.com/bannerify/bannerify-go/option"
)

// KeyService contains methods and other services that help with interacting with
// the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	Options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r *KeyService) {
	r = &KeyService{}
	r.Options = opts
	return
}

func (r *KeyService) New(ctx context.Context, body KeyNewParams, opts ...option.RequestOption) (res *KeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *KeyService) Verify(ctx context.Context, body KeyVerifyParams, opts ...option.RequestOption) (res *KeyVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type KeyNewResponse struct {
	// The newly created api key, do not store this on your own system but pass it
	// along to your user.
	Key string `json:"key,required"`
	// The id of the key. This is not a secret and can be stored as a reference if you
	// wish. You need the keyId to update or delete a key later.
	KeyID string             `json:"keyId,required"`
	JSON  keyNewResponseJSON `json:"-"`
}

// keyNewResponseJSON contains the JSON metadata for the struct [KeyNewResponse]
type keyNewResponseJSON struct {
	Key         apijson.Field
	KeyID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseJSON) RawJSON() string {
	return r.raw
}

type KeyVerifyResponse struct {
	// Whether the key is valid or not. A key could be invalid for a number of reasons,
	// for example if it has expired, has no more verifications left or if it has been
	// deleted.
	Valid bool `json:"valid,required"`
	// If the key is invalid this field will be set to the reason why it is invalid.
	// Possible values are:
	//
	//   - NOT_FOUND: the key does not exist or has expired
	//   - FORBIDDEN: the key is not allowed to access the api
	//   - USAGE_EXCEEDED: the key has exceeded its request limit
	//   - RATE_LIMITED: the key has been ratelimited,
	//   - INSUFFICIENT_PERMISSIONS: you do not have the required permissions to perform
	//     this action
	Code KeyVerifyResponseCode `json:"code"`
	// The unix timestamp in milliseconds when the key was created
	CreatedAt int64 `json:"createdAt"`
	// The unix timestamp in milliseconds when the key was deleted. We don't delete the
	// key outright, you can restore it later.
	DeletedAt int64 `json:"deletedAt"`
	// The unix timestamp in milliseconds when the key will expire. If this field is
	// null or undefined, the key is not expiring.
	Expires int64 `json:"expires"`
	// The id of the key
	KeyID string `json:"keyId"`
	// Any additional metadata you want to store with the key
	Meta map[string]interface{} `json:"meta"`
	// The name of the key, give keys a name to easily identifiy their purpose
	Name string `json:"name"`
	// The id of the tenant associated with this key. Use whatever reference you have
	// in your system to identify the tenant. When verifying the key, we will send this
	// field back to you, so you know who is accessing your API.
	OwnerID string `json:"ownerId"`
	// The ratelimit configuration for this key. If this field is null or undefined,
	// the key has no ratelimit.
	Ratelimit KeyVerifyResponseRatelimit `json:"ratelimit"`
	// The number of requests that can be made with this key before it becomes invalid.
	// If this field is null or undefined, the key has no request limit.
	Remaining int64                 `json:"remaining"`
	JSON      keyVerifyResponseJSON `json:"-"`
}

// keyVerifyResponseJSON contains the JSON metadata for the struct
// [KeyVerifyResponse]
type keyVerifyResponseJSON struct {
	Valid       apijson.Field
	Code        apijson.Field
	CreatedAt   apijson.Field
	DeletedAt   apijson.Field
	Expires     apijson.Field
	KeyID       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	OwnerID     apijson.Field
	Ratelimit   apijson.Field
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyVerifyResponseJSON) RawJSON() string {
	return r.raw
}

// If the key is invalid this field will be set to the reason why it is invalid.
// Possible values are:
//
//   - NOT_FOUND: the key does not exist or has expired
//   - FORBIDDEN: the key is not allowed to access the api
//   - USAGE_EXCEEDED: the key has exceeded its request limit
//   - RATE_LIMITED: the key has been ratelimited,
//   - INSUFFICIENT_PERMISSIONS: you do not have the required permissions to perform
//     this action
type KeyVerifyResponseCode string

const (
	KeyVerifyResponseCodeNotFound                KeyVerifyResponseCode = "NOT_FOUND"
	KeyVerifyResponseCodeForbidden               KeyVerifyResponseCode = "FORBIDDEN"
	KeyVerifyResponseCodeUsageExceeded           KeyVerifyResponseCode = "USAGE_EXCEEDED"
	KeyVerifyResponseCodeRateLimited             KeyVerifyResponseCode = "RATE_LIMITED"
	KeyVerifyResponseCodeUnauthorized            KeyVerifyResponseCode = "UNAUTHORIZED"
	KeyVerifyResponseCodeDisabled                KeyVerifyResponseCode = "DISABLED"
	KeyVerifyResponseCodeInsufficientPermissions KeyVerifyResponseCode = "INSUFFICIENT_PERMISSIONS"
	KeyVerifyResponseCodeExpired                 KeyVerifyResponseCode = "EXPIRED"
)

func (r KeyVerifyResponseCode) IsKnown() bool {
	switch r {
	case KeyVerifyResponseCodeNotFound, KeyVerifyResponseCodeForbidden, KeyVerifyResponseCodeUsageExceeded, KeyVerifyResponseCodeRateLimited, KeyVerifyResponseCodeUnauthorized, KeyVerifyResponseCodeDisabled, KeyVerifyResponseCodeInsufficientPermissions, KeyVerifyResponseCodeExpired:
		return true
	}
	return false
}

// The ratelimit configuration for this key. If this field is null or undefined,
// the key has no ratelimit.
type KeyVerifyResponseRatelimit struct {
	// Maximum number of requests that can be made inside a window
	Limit int64 `json:"limit,required"`
	// Remaining requests after this verification
	Remaining int64 `json:"remaining,required"`
	// Unix timestamp in milliseconds when the ratelimit will reset
	Reset int64                          `json:"reset,required"`
	JSON  keyVerifyResponseRatelimitJSON `json:"-"`
}

// keyVerifyResponseRatelimitJSON contains the JSON metadata for the struct
// [KeyVerifyResponseRatelimit]
type keyVerifyResponseRatelimitJSON struct {
	Limit       apijson.Field
	Remaining   apijson.Field
	Reset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyVerifyResponseRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyVerifyResponseRatelimitJSON) RawJSON() string {
	return r.raw
}

type KeyNewParams struct {
	// Choose an `API` where this key should be created.
	APIID param.Field[string] `json:"apiId,required"`
	// The byte length used to generate your key determines its entropy as well as its
	// length. Higher is better, but keys become longer and more annoying to handle.
	// The default is 16 bytes, or 2^^128 possible combinations.
	ByteLength param.Field[int64] `json:"byteLength"`
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
	Ratelimit param.Field[KeyNewParamsRatelimit] `json:"ratelimit"`
	// You can limit the number of requests a key can make. Once a key reaches 0
	// remaining requests, it will automatically be disabled and is no longer valid
	// unless you update it.
	Remaining param.Field[int64] `json:"remaining"`
}

func (r KeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Unkey comes with per-key ratelimiting out of the box.
type KeyNewParamsRatelimit struct {
	// The total amount of burstable requests.
	Limit param.Field[int64] `json:"limit,required"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	RefillInterval param.Field[int64] `json:"refillInterval,required"`
	// How many tokens to refill during each refillInterval.
	RefillRate param.Field[int64] `json:"refillRate,required"`
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
	// accurate.
	Type param.Field[KeyNewParamsRatelimitType] `json:"type"`
}

func (r KeyNewParamsRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more
// accurate.
type KeyNewParamsRatelimitType string

const (
	KeyNewParamsRatelimitTypeFast       KeyNewParamsRatelimitType = "fast"
	KeyNewParamsRatelimitTypeConsistent KeyNewParamsRatelimitType = "consistent"
)

func (r KeyNewParamsRatelimitType) IsKnown() bool {
	switch r {
	case KeyNewParamsRatelimitTypeFast, KeyNewParamsRatelimitTypeConsistent:
		return true
	}
	return false
}

type KeyVerifyParams struct {
	// The key to verify
	Key param.Field[string] `json:"key,required"`
	// The id of the api where the key belongs to. This is optional for now but will be
	// required soon. The key will be verified against the api's configuration. If the
	// key does not belong to the api, the verification will fail.
	APIID param.Field[string] `json:"apiId"`
}

func (r KeyVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
