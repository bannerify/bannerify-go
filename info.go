// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bannerify/bannerify-go/internal/apijson"
	"github.com/bannerify/bannerify-go/internal/apiquery"
	"github.com/bannerify/bannerify-go/internal/param"
	"github.com/bannerify/bannerify-go/internal/requestconfig"
	"github.com/bannerify/bannerify-go/option"
)

// InfoService contains methods and other services that help with interacting with
// the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInfoService] method instead.
type InfoService struct {
	Options []option.RequestOption
}

// NewInfoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInfoService(opts ...option.RequestOption) (r *InfoService) {
	r = &InfoService{}
	r.Options = opts
	return
}

// Get project info
func (r *InfoService) Get(ctx context.Context, query InfoGetParams, opts ...option.RequestOption) (res *InfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InfoGetResponse struct {
	ID        string              `json:"id,required"`
	CreatedAt string              `json:"createdAt,required"`
	Name      string              `json:"name,required"`
	JSON      infoGetResponseJSON `json:"-"`
}

// infoGetResponseJSON contains the JSON metadata for the struct [InfoGetResponse]
type infoGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r infoGetResponseJSON) RawJSON() string {
	return r.raw
}

type InfoGetParams struct {
	// The api key to use for this request
	APIKey param.Field[string] `query:"apiKey,required"`
}

// URLQuery serializes [InfoGetParams]'s query parameters as `url.Values`.
func (r InfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
