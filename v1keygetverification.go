// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/bannerify-go/internal/apijson"
	"github.com/stainless-sdks/bannerify-go/internal/apiquery"
	"github.com/stainless-sdks/bannerify-go/internal/param"
	"github.com/stainless-sdks/bannerify-go/internal/requestconfig"
	"github.com/stainless-sdks/bannerify-go/option"
)

// V1KeyGetVerificationService contains methods and other services that help with
// interacting with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1KeyGetVerificationService] method instead.
type V1KeyGetVerificationService struct {
	Options []option.RequestOption
}

// NewV1KeyGetVerificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1KeyGetVerificationService(opts ...option.RequestOption) (r *V1KeyGetVerificationService) {
	r = &V1KeyGetVerificationService{}
	r.Options = opts
	return
}

func (r *V1KeyGetVerificationService) List(ctx context.Context, query V1KeyGetVerificationListParams, opts ...option.RequestOption) (res *V1KeyGetVerificationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.getVerifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V1KeyGetVerificationListResponse struct {
	Verifications []V1KeyGetVerificationListResponseVerification `json:"verifications,required"`
	JSON          v1KeyGetVerificationListResponseJSON           `json:"-"`
}

// v1KeyGetVerificationListResponseJSON contains the JSON metadata for the struct
// [V1KeyGetVerificationListResponse]
type v1KeyGetVerificationListResponseJSON struct {
	Verifications apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1KeyGetVerificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyGetVerificationListResponseJSON) RawJSON() string {
	return r.raw
}

type V1KeyGetVerificationListResponseVerification struct {
	// The number of requests that were rate limited
	RateLimited int64 `json:"rateLimited,required"`
	// The number of successful requests
	Success int64 `json:"success,required"`
	// The timestamp of the usage data
	Time int64 `json:"time,required"`
	// The number of requests that exceeded the usage limit
	UsageExceeded int64                                            `json:"usageExceeded,required"`
	JSON          v1KeyGetVerificationListResponseVerificationJSON `json:"-"`
}

// v1KeyGetVerificationListResponseVerificationJSON contains the JSON metadata for
// the struct [V1KeyGetVerificationListResponseVerification]
type v1KeyGetVerificationListResponseVerificationJSON struct {
	RateLimited   apijson.Field
	Success       apijson.Field
	Time          apijson.Field
	UsageExceeded apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V1KeyGetVerificationListResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyGetVerificationListResponseVerificationJSON) RawJSON() string {
	return r.raw
}

type V1KeyGetVerificationListParams struct {
	// The end of the period to fetch usage for as unix milliseconds timestamp
	End param.Field[int64] `query:"end"`
	// The granularity of the usage data to fetch, currently only `day` is supported
	Granularity param.Field[V1KeyGetVerificationListParamsGranularity] `query:"granularity"`
	// The id of the key to fetch, either `keyId` or `ownerId` must be provided
	KeyID param.Field[string] `query:"keyId"`
	// The owner id to fetch keys for, either `keyId` or `ownerId` must be provided
	OwnerID param.Field[string] `query:"ownerId"`
	// The start of the period to fetch usage for as unix milliseconds timestamp
	Start param.Field[int64] `query:"start"`
}

// URLQuery serializes [V1KeyGetVerificationListParams]'s query parameters as
// `url.Values`.
func (r V1KeyGetVerificationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The granularity of the usage data to fetch, currently only `day` is supported
type V1KeyGetVerificationListParamsGranularity string

const (
	V1KeyGetVerificationListParamsGranularityDay V1KeyGetVerificationListParamsGranularity = "day"
)

func (r V1KeyGetVerificationListParamsGranularity) IsKnown() bool {
	switch r {
	case V1KeyGetVerificationListParamsGranularityDay:
		return true
	}
	return false
}
