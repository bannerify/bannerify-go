// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempbannerify

import (
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/apijson"
	"github.com/stainless-sdks/TEMP_bannerify-go/option"
)

// V1KeyService contains methods and other services that help with interacting with
// the bannerify API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewV1KeyService] method instead.
type V1KeyService struct {
	Options          []option.RequestOption
	GetKey           *V1KeyGetKeyService
	DeleteKey        *V1KeyDeleteKeyService
	CreateKey        *V1KeyCreateKeyService
	VerifyKey        *V1KeyVerifyKeyService
	UpdateKey        *V1KeyUpdateKeyService
	UpdateRemaining  *V1KeyUpdateRemainingService
	GetVerifications *V1KeyGetVerificationService
}

// NewV1KeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1KeyService(opts ...option.RequestOption) (r *V1KeyService) {
	r = &V1KeyService{}
	r.Options = opts
	r.GetKey = NewV1KeyGetKeyService(opts...)
	r.DeleteKey = NewV1KeyDeleteKeyService(opts...)
	r.CreateKey = NewV1KeyCreateKeyService(opts...)
	r.VerifyKey = NewV1KeyVerifyKeyService(opts...)
	r.UpdateKey = NewV1KeyUpdateKeyService(opts...)
	r.UpdateRemaining = NewV1KeyUpdateRemainingService(opts...)
	r.GetVerifications = NewV1KeyGetVerificationService(opts...)
	return
}

type V1KeysVerifyKeyResponse struct {
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
	//   - RATE_LIMITED: the key has been ratelimited
	//   - UNAUTHORIZED: the key is not authorized
	//   - DISABLED: the key is disabled
	//   - INSUFFICIENT_PERMISSIONS: you do not have the required permissions to perform
	//     this action
	Code V1KeysVerifyKeyResponseCode `json:"code"`
	// Sets the key to be enabled or disabled. Disabled keys will not verify.
	Enabled bool `json:"enabled"`
	// The environment of the key, this is what what you set when you crated the key
	Environment string `json:"environment"`
	// The unix timestamp in milliseconds when the key will expire. If this field is
	// null or undefined, the key is not expiring.
	Expires float64 `json:"expires"`
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
	// A list of all the permissions this key is connected to.
	Permissions []string `json:"permissions"`
	// The ratelimit configuration for this key. If this field is null or undefined,
	// the key has no ratelimit.
	Ratelimit V1KeysVerifyKeyResponseRatelimit `json:"ratelimit"`
	// The number of requests that can be made with this key before it becomes invalid.
	// If this field is null or undefined, the key has no request limit.
	Remaining float64                     `json:"remaining"`
	JSON      v1KeysVerifyKeyResponseJSON `json:"-"`
}

// v1KeysVerifyKeyResponseJSON contains the JSON metadata for the struct
// [V1KeysVerifyKeyResponse]
type v1KeysVerifyKeyResponseJSON struct {
	Valid       apijson.Field
	Code        apijson.Field
	Enabled     apijson.Field
	Environment apijson.Field
	Expires     apijson.Field
	KeyID       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	OwnerID     apijson.Field
	Permissions apijson.Field
	Ratelimit   apijson.Field
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeysVerifyKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeysVerifyKeyResponseJSON) RawJSON() string {
	return r.raw
}

// If the key is invalid this field will be set to the reason why it is invalid.
// Possible values are:
//
//   - NOT_FOUND: the key does not exist or has expired
//   - FORBIDDEN: the key is not allowed to access the api
//   - USAGE_EXCEEDED: the key has exceeded its request limit
//   - RATE_LIMITED: the key has been ratelimited
//   - UNAUTHORIZED: the key is not authorized
//   - DISABLED: the key is disabled
//   - INSUFFICIENT_PERMISSIONS: you do not have the required permissions to perform
//     this action
type V1KeysVerifyKeyResponseCode string

const (
	V1KeysVerifyKeyResponseCodeNotFound                V1KeysVerifyKeyResponseCode = "NOT_FOUND"
	V1KeysVerifyKeyResponseCodeForbidden               V1KeysVerifyKeyResponseCode = "FORBIDDEN"
	V1KeysVerifyKeyResponseCodeUsageExceeded           V1KeysVerifyKeyResponseCode = "USAGE_EXCEEDED"
	V1KeysVerifyKeyResponseCodeRateLimited             V1KeysVerifyKeyResponseCode = "RATE_LIMITED"
	V1KeysVerifyKeyResponseCodeUnauthorized            V1KeysVerifyKeyResponseCode = "UNAUTHORIZED"
	V1KeysVerifyKeyResponseCodeDisabled                V1KeysVerifyKeyResponseCode = "DISABLED"
	V1KeysVerifyKeyResponseCodeInsufficientPermissions V1KeysVerifyKeyResponseCode = "INSUFFICIENT_PERMISSIONS"
)

func (r V1KeysVerifyKeyResponseCode) IsKnown() bool {
	switch r {
	case V1KeysVerifyKeyResponseCodeNotFound, V1KeysVerifyKeyResponseCodeForbidden, V1KeysVerifyKeyResponseCodeUsageExceeded, V1KeysVerifyKeyResponseCodeRateLimited, V1KeysVerifyKeyResponseCodeUnauthorized, V1KeysVerifyKeyResponseCodeDisabled, V1KeysVerifyKeyResponseCodeInsufficientPermissions:
		return true
	}
	return false
}

// The ratelimit configuration for this key. If this field is null or undefined,
// the key has no ratelimit.
type V1KeysVerifyKeyResponseRatelimit struct {
	// Maximum number of requests that can be made inside a window
	Limit float64 `json:"limit,required"`
	// Remaining requests after this verification
	Remaining float64 `json:"remaining,required"`
	// Unix timestamp in milliseconds when the ratelimit will reset
	Reset float64                              `json:"reset,required"`
	JSON  v1KeysVerifyKeyResponseRatelimitJSON `json:"-"`
}

// v1KeysVerifyKeyResponseRatelimitJSON contains the JSON metadata for the struct
// [V1KeysVerifyKeyResponseRatelimit]
type v1KeysVerifyKeyResponseRatelimitJSON struct {
	Limit       apijson.Field
	Remaining   apijson.Field
	Reset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeysVerifyKeyResponseRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeysVerifyKeyResponseRatelimitJSON) RawJSON() string {
	return r.raw
}
