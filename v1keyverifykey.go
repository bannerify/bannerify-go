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
	Ratelimit     param.Field[V1KeyVerifyKeyNewParamsRatelimit]     `json:"ratelimit"`
}

func (r V1KeyVerifyKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Perform RBAC checks
type V1KeyVerifyKeyNewParamsAuthorization struct {
	// A query for which permissions you require
	Permissions param.Field[interface{}] `json:"permissions"`
}

func (r V1KeyVerifyKeyNewParamsAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1KeyVerifyKeyNewParamsRatelimit struct {
	// Override how many tokens are deducted during the ratelimit operation.
	Cost param.Field[int64] `json:"cost"`
}

func (r V1KeyVerifyKeyNewParamsRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
