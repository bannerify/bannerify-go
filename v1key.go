// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"reflect"

	"github.com/stainless-sdks/bannerify-go/internal/apijson"
	"github.com/stainless-sdks/bannerify-go/option"
	"github.com/tidwall/gjson"
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
	// The id of the key
	KeyID string `json:"keyId"`
	// Whether the key is valid or not. A key could be invalid for a number of reasons,
	// for example if it has expired, has no more verifications left or if it has been
	// deleted.
	Valid bool `json:"valid,required"`
	// The name of the key, give keys a name to easily identifiy their purpose
	Name string `json:"name"`
	// The id of the tenant associated with this key. Use whatever reference you have
	// in your system to identify the tenant. When verifying the key, we will send this
	// field back to you, so you know who is accessing your API.
	OwnerID string      `json:"ownerId"`
	Meta    interface{} `json:"meta,required"`
	// The unix timestamp in milliseconds when the key will expire. If this field is
	// null or undefined, the key is not expiring.
	Expires   float64     `json:"expires"`
	Ratelimit interface{} `json:"ratelimit,required"`
	// The number of requests that can be made with this key before it becomes invalid.
	// If this field is null or undefined, the key has no request limit.
	Remaining float64 `json:"remaining"`
	// A machine readable response code.
	Code V1KeysVerifyKeyResponseCode `json:"code,required"`
	// Sets the key to be enabled or disabled. Disabled keys will not verify.
	Enabled     bool        `json:"enabled"`
	Permissions interface{} `json:"permissions,required"`
	// The environment of the key, this is what what you set when you crated the key
	Environment string                      `json:"environment"`
	JSON        v1KeysVerifyKeyResponseJSON `json:"-"`
	union       V1KeysVerifyKeyResponseUnion
}

// v1KeysVerifyKeyResponseJSON contains the JSON metadata for the struct
// [V1KeysVerifyKeyResponse]
type v1KeysVerifyKeyResponseJSON struct {
	KeyID       apijson.Field
	Valid       apijson.Field
	Name        apijson.Field
	OwnerID     apijson.Field
	Meta        apijson.Field
	Expires     apijson.Field
	Ratelimit   apijson.Field
	Remaining   apijson.Field
	Code        apijson.Field
	Enabled     apijson.Field
	Permissions apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v1KeysVerifyKeyResponseJSON) RawJSON() string {
	return r.raw
}

func (r *V1KeysVerifyKeyResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r V1KeysVerifyKeyResponse) AsUnion() V1KeysVerifyKeyResponseUnion {
	return r.union
}

// Union satisfied by [V1KeysVerifyKeyResponseObject] or
// [V1KeysVerifyKeyResponseObject].
type V1KeysVerifyKeyResponseUnion interface {
	implementsV1KeysVerifyKeyResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1KeysVerifyKeyResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1KeysVerifyKeyResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V1KeysVerifyKeyResponseObject{}),
		},
	)
}

type V1KeysVerifyKeyResponseObject struct {
	// A machine readable response code.
	Code V1KeysVerifyKeyResponseObjectCode `json:"code,required"`
	// The id of the key
	KeyID string `json:"keyId,required"`
	// Whether the key is valid or not. A key could be invalid for a number of reasons,
	// for example if it has expired, has no more verifications left or if it has been
	// deleted.
	Valid bool `json:"valid,required"`
	// Sets the key to be enabled or disabled. Disabled keys will not verify.
	Enabled bool `json:"enabled"`
	// The environment of the key, this is what what you set when you crated the key
	Environment string `json:"environment"`
	// The unix timestamp in milliseconds when the key will expire. If this field is
	// null or undefined, the key is not expiring.
	Expires float64 `json:"expires"`
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
	Ratelimit V1KeysVerifyKeyResponseObjectRatelimit `json:"ratelimit"`
	// The number of requests that can be made with this key before it becomes invalid.
	// If this field is null or undefined, the key has no request limit.
	Remaining float64                           `json:"remaining"`
	JSON      v1KeysVerifyKeyResponseObjectJSON `json:"-"`
}

// v1KeysVerifyKeyResponseObjectJSON contains the JSON metadata for the struct
// [V1KeysVerifyKeyResponseObject]
type v1KeysVerifyKeyResponseObjectJSON struct {
	Code        apijson.Field
	KeyID       apijson.Field
	Valid       apijson.Field
	Enabled     apijson.Field
	Environment apijson.Field
	Expires     apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	OwnerID     apijson.Field
	Permissions apijson.Field
	Ratelimit   apijson.Field
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeysVerifyKeyResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeysVerifyKeyResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r V1KeysVerifyKeyResponseObject) implementsV1KeysVerifyKeyResponse() {}

// A machine readable response code.
type V1KeysVerifyKeyResponseObjectCode string

const (
	V1KeysVerifyKeyResponseObjectCodeValid V1KeysVerifyKeyResponseObjectCode = "VALID"
)

func (r V1KeysVerifyKeyResponseObjectCode) IsKnown() bool {
	switch r {
	case V1KeysVerifyKeyResponseObjectCodeValid:
		return true
	}
	return false
}

// The ratelimit configuration for this key. If this field is null or undefined,
// the key has no ratelimit.
type V1KeysVerifyKeyResponseObjectRatelimit struct {
	// Maximum number of requests that can be made inside a window
	Limit float64 `json:"limit,required"`
	// Remaining requests after this verification
	Remaining float64 `json:"remaining,required"`
	// Unix timestamp in milliseconds when the ratelimit will reset
	Reset float64                                    `json:"reset,required"`
	JSON  v1KeysVerifyKeyResponseObjectRatelimitJSON `json:"-"`
}

// v1KeysVerifyKeyResponseObjectRatelimitJSON contains the JSON metadata for the
// struct [V1KeysVerifyKeyResponseObjectRatelimit]
type v1KeysVerifyKeyResponseObjectRatelimitJSON struct {
	Limit       apijson.Field
	Remaining   apijson.Field
	Reset       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeysVerifyKeyResponseObjectRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeysVerifyKeyResponseObjectRatelimitJSON) RawJSON() string {
	return r.raw
}

// A machine readable response code.
type V1KeysVerifyKeyResponseCode string

const (
	V1KeysVerifyKeyResponseCodeValid                   V1KeysVerifyKeyResponseCode = "VALID"
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
	case V1KeysVerifyKeyResponseCodeValid, V1KeysVerifyKeyResponseCodeNotFound, V1KeysVerifyKeyResponseCodeForbidden, V1KeysVerifyKeyResponseCodeUsageExceeded, V1KeysVerifyKeyResponseCodeRateLimited, V1KeysVerifyKeyResponseCodeUnauthorized, V1KeysVerifyKeyResponseCodeDisabled, V1KeysVerifyKeyResponseCodeInsufficientPermissions:
		return true
	}
	return false
}
