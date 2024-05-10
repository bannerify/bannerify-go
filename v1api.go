// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bannerify

import (
	"github.com/stainless-sdks/bannerify-go/option"
)

// V1APIService contains methods and other services that help with interacting with
// the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1APIService] method instead.
type V1APIService struct {
	Options   []option.RequestOption
	GetAPI    *V1APIGetAPIService
	CreateAPI *V1APICreateAPIService
}

// NewV1APIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1APIService(opts ...option.RequestOption) (r *V1APIService) {
	r = &V1APIService{}
	r.Options = opts
	r.GetAPI = NewV1APIGetAPIService(opts...)
	r.CreateAPI = NewV1APICreateAPIService(opts...)
	return
}
