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

// V1KeyDeleteKeyService contains methods and other services that help with
// interacting with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1KeyDeleteKeyService] method instead.
type V1KeyDeleteKeyService struct {
	Options []option.RequestOption
}

// NewV1KeyDeleteKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1KeyDeleteKeyService(opts ...option.RequestOption) (r *V1KeyDeleteKeyService) {
	r = &V1KeyDeleteKeyService{}
	r.Options = opts
	return
}

func (r *V1KeyDeleteKeyService) New(ctx context.Context, body V1KeyDeleteKeyNewParams, opts ...option.RequestOption) (res *V1KeyDeleteKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.deleteKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1KeyDeleteKeyNewResponse = interface{}

type V1KeyDeleteKeyNewParams struct {
	// The id of the key to revoke
	KeyID param.Field[string] `json:"keyId,required"`
}

func (r V1KeyDeleteKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
