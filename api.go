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

// APIService contains methods and other services that help with interacting with
// the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIService] method instead.
type APIService struct {
	Options []option.RequestOption
	Keys    *APIKeyService
}

// NewAPIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIService(opts ...option.RequestOption) (r *APIService) {
	r = &APIService{}
	r.Options = opts
	r.Keys = NewAPIKeyService(opts...)
	return
}

func (r *APIService) DeleteAPI(ctx context.Context, body APIDeleteAPIParams, opts ...option.RequestOption) (res *APIDeleteAPIResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/apis.deleteApi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIDeleteAPIResponse = interface{}

type APIDeleteAPIParams struct {
	// The id of the api to delete
	APIID param.Field[string] `json:"apiId,required"`
}

func (r APIDeleteAPIParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
