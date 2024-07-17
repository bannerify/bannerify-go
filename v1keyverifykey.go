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

// V1KeyVerifyKeyService contains methods and other services that help with
// interacting with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1KeyVerifyKeyService] method instead.
type V1KeyVerifyKeyService struct {
	Options []option.RequestOption
}

// NewV1KeyVerifyKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1KeyVerifyKeyService(opts ...option.RequestOption) (r *V1KeyVerifyKeyService) {
	r = &V1KeyVerifyKeyService{}
	r.Options = opts
	return
}

func (r *V1KeyVerifyKeyService) New(ctx context.Context, body V1KeyVerifyKeyNewParams, opts ...option.RequestOption) (res *V1KeysVerifyKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.verifyKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1KeyVerifyKeyNewParams struct {
	// The key to verify
	Key param.Field[string] `json:"key,required"`
	// The id of the api where the key belongs to. This is optional for now but will be
	// required soon. The key will be verified against the api's configuration. If the
	// key does not belong to the api, the verification will fail.
	APIID param.Field[string] `json:"apiId"`
	// Perform RBAC checks
	Authorization param.Field[V1KeyVerifyKeyNewParamsAuthorization] `json:"authorization"`
	// Use 'ratelimits' with `[{ name: "default", cost: 2}]`
	Ratelimit param.Field[V1KeyVerifyKeyNewParamsRatelimit] `json:"ratelimit"`
	// You can check against multiple ratelimits when verifying a key. Let's say you
	// are building an app that uses AI under the hood and you want to limit your
	// customers to 500 requests per hour, but also ensure they use up less than 20k
	// tokens per day.
	Ratelimits param.Field[[]V1KeyVerifyKeyNewParamsRatelimit] `json:"ratelimits"`
}

func (r V1KeyVerifyKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Perform RBAC checks
type V1KeyVerifyKeyNewParamsAuthorization struct {
	// A query for which permissions you require
	Permissions param.Field[V1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion] `json:"permissions"`
}

func (r V1KeyVerifyKeyNewParamsAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A query for which permissions you require
type V1KeyVerifyKeyNewParamsAuthorizationPermissions struct {
	And param.Field[interface{}] `json:"and,required"`
	Or  param.Field[interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissions) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion() {
}

// A query for which permissions you require
//
// Satisfied by [shared.UnionString],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAnd],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOr],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissions].
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion interface {
	ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion()
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAnd struct {
	And param.Field[[]V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion] `json:"and,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion() {
}

// A query for which permissions you require
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAnd struct {
	And param.Field[interface{}] `json:"and,required"`
	Or  param.Field[interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion() {
}

// A query for which permissions you require
//
// Satisfied by [shared.UnionString],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndAnd],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOr],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAnd].
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion interface {
	ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion()
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndAnd struct {
	And param.Field[[]interface{}] `json:"and,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion() {
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOr struct {
	Or param.Field[[]V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion() {
}

// A query for which permissions you require
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOr struct {
	And param.Field[interface{}] `json:"and,required"`
	Or  param.Field[interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion() {
}

// A query for which permissions you require
//
// Satisfied by [shared.UnionString],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrAnd],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrOr],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOr].
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion interface {
	ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion()
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrAnd struct {
	And param.Field[[]interface{}] `json:"and,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion() {
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrOr struct {
	Or param.Field[[]interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion() {
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOr struct {
	Or param.Field[[]V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion() {
}

// A query for which permissions you require
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOr struct {
	And param.Field[interface{}] `json:"and,required"`
	Or  param.Field[interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion() {
}

// A query for which permissions you require
//
// Satisfied by [shared.UnionString],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAnd],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrOr],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOr].
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion interface {
	ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion()
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAnd struct {
	And param.Field[[]V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion] `json:"and,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion() {
}

// A query for which permissions you require
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAnd struct {
	And param.Field[interface{}] `json:"and,required"`
	Or  param.Field[interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion() {
}

// A query for which permissions you require
//
// Satisfied by [shared.UnionString],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndAnd],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndOr],
// [V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAnd].
type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion interface {
	ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion()
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndAnd struct {
	And param.Field[[]interface{}] `json:"and,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndAnd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndAnd) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion() {
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndOr struct {
	Or param.Field[[]interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion() {
}

type V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrOr struct {
	Or param.Field[[]interface{}] `json:"or,required"`
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrOr) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r V1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrOr) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion() {
}

// Use 'ratelimits' with `[{ name: "default", cost: 2}]`
type V1KeyVerifyKeyNewParamsRatelimit struct {
	// Override how many tokens are deducted during the ratelimit operation.
	Cost param.Field[int64] `json:"cost"`
}

func (r V1KeyVerifyKeyNewParamsRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
