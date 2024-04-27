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

// V1APIGetAPIService contains methods and other services that help with
// interacting with the bannerify API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV1APIGetAPIService] method
// instead.
type V1APIGetAPIService struct {
	Options []option.RequestOption
}

// NewV1APIGetAPIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1APIGetAPIService(opts ...option.RequestOption) (r *V1APIGetAPIService) {
	r = &V1APIGetAPIService{}
	r.Options = opts
	return
}

func (r *V1APIGetAPIService) Get(ctx context.Context, query V1APIGetAPIGetParams, opts ...option.RequestOption) (res *V1APIGetAPIGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/apis.getApi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V1APIGetAPIGetResponse struct {
	// The id of the key
	ID string `json:"id,required"`
	// The id of the workspace that owns the api
	WorkspaceID string `json:"workspaceId,required"`
	// The name of the api. This is internal and your users will not see this.
	Name string                     `json:"name"`
	JSON v1APIGetAPIGetResponseJSON `json:"-"`
}

// v1APIGetAPIGetResponseJSON contains the JSON metadata for the struct
// [V1APIGetAPIGetResponse]
type v1APIGetAPIGetResponseJSON struct {
	ID          apijson.Field
	WorkspaceID apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1APIGetAPIGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1APIGetAPIGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1APIGetAPIGetParams struct {
	// The id of the api to fetch
	APIID param.Field[string] `query:"apiId,required"`
}

// URLQuery serializes [V1APIGetAPIGetParams]'s query parameters as `url.Values`.
func (r V1APIGetAPIGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
