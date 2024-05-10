// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/bannerify-go/internal/apijson"
	"github.com/stainless-sdks/bannerify-go/internal/requestconfig"
	"github.com/stainless-sdks/bannerify-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options []option.RequestOption
	Keys    *V1KeyService
	APIs    *V1APIService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r *V1Service) {
	r = &V1Service{}
	r.Options = opts
	r.Keys = NewV1KeyService(opts...)
	r.APIs = NewV1APIService(opts...)
	return
}

func (r *V1Service) Liveness(ctx context.Context, opts ...option.RequestOption) (res *V1LivenessResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/liveness"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V1LivenessResponse struct {
	Services V1LivenessResponseServices `json:"services,required"`
	// The status of the server
	Status string                 `json:"status,required"`
	JSON   v1LivenessResponseJSON `json:"-"`
}

// v1LivenessResponseJSON contains the JSON metadata for the struct
// [V1LivenessResponse]
type v1LivenessResponseJSON struct {
	Services    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1LivenessResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1LivenessResponseJSON) RawJSON() string {
	return r.raw
}

type V1LivenessResponseServices struct {
	// The name of the connected analytics service
	Analytics string `json:"analytics,required"`
	// The name of the connected logger service
	Logger string `json:"logger,required"`
	// The name of the connected metrics service
	Metrics string `json:"metrics,required"`
	// The name of the connected ratelimit service
	Ratelimit string `json:"ratelimit,required"`
	// The name of the connected usagelimit service
	Usagelimit string                         `json:"usagelimit,required"`
	JSON       v1LivenessResponseServicesJSON `json:"-"`
}

// v1LivenessResponseServicesJSON contains the JSON metadata for the struct
// [V1LivenessResponseServices]
type v1LivenessResponseServicesJSON struct {
	Analytics   apijson.Field
	Logger      apijson.Field
	Metrics     apijson.Field
	Ratelimit   apijson.Field
	Usagelimit  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1LivenessResponseServices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1LivenessResponseServicesJSON) RawJSON() string {
	return r.raw
}
