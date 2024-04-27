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

// V1APICreateAPIService contains methods and other services that help with
// interacting with the bannerify API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV1APICreateAPIService] method
// instead.
type V1APICreateAPIService struct {
	Options []option.RequestOption
}

// NewV1APICreateAPIService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1APICreateAPIService(opts ...option.RequestOption) (r *V1APICreateAPIService) {
	r = &V1APICreateAPIService{}
	r.Options = opts
	return
}

func (r *V1APICreateAPIService) New(ctx context.Context, body V1APICreateAPINewParams, opts ...option.RequestOption) (res *V1APICreateAPINewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/apis.createApi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1APICreateAPINewResponse struct {
	// The id of the api
	APIID string                        `json:"apiId,required"`
	JSON  v1APICreateAPINewResponseJSON `json:"-"`
}

// v1APICreateAPINewResponseJSON contains the JSON metadata for the struct
// [V1APICreateAPINewResponse]
type v1APICreateAPINewResponseJSON struct {
	APIID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1APICreateAPINewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1APICreateAPINewResponseJSON) RawJSON() string {
	return r.raw
}

type V1APICreateAPINewParams struct {
	// The name for your API. This is not customer facing.
	Name param.Field[string] `json:"name,required"`
}

func (r V1APICreateAPINewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
