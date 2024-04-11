// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tempbannerify

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/TEMP_bannerify-go/internal/apijson"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/param"
	"github.com/stainless-sdks/TEMP_bannerify-go/internal/requestconfig"
	"github.com/stainless-sdks/TEMP_bannerify-go/option"
)

// V1KeyUpdateRemainingService contains methods and other services that help with
// interacting with the bannerify API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV1KeyUpdateRemainingService]
// method instead.
type V1KeyUpdateRemainingService struct {
	Options []option.RequestOption
}

// NewV1KeyUpdateRemainingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1KeyUpdateRemainingService(opts ...option.RequestOption) (r *V1KeyUpdateRemainingService) {
	r = &V1KeyUpdateRemainingService{}
	r.Options = opts
	return
}

func (r *V1KeyUpdateRemainingService) New(ctx context.Context, body V1KeyUpdateRemainingNewParams, opts ...option.RequestOption) (res *V1KeyUpdateRemainingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/keys.updateRemaining"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1KeyUpdateRemainingNewResponse struct {
	// The number of remaining requests for this key after updating it. `null` means
	// unlimited.
	Remaining int64                               `json:"remaining,required,nullable"`
	JSON      v1KeyUpdateRemainingNewResponseJSON `json:"-"`
}

// v1KeyUpdateRemainingNewResponseJSON contains the JSON metadata for the struct
// [V1KeyUpdateRemainingNewResponse]
type v1KeyUpdateRemainingNewResponseJSON struct {
	Remaining   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1KeyUpdateRemainingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1KeyUpdateRemainingNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1KeyUpdateRemainingNewParams struct {
	// The id of the key you want to modify
	KeyID param.Field[string] `json:"keyId,required"`
	// The operation you want to perform on the remaining count
	Op param.Field[V1KeyUpdateRemainingNewParamsOp] `json:"op,required"`
	// The value you want to set, add or subtract the remaining count by
	Value param.Field[int64] `json:"value,required"`
}

func (r V1KeyUpdateRemainingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operation you want to perform on the remaining count
type V1KeyUpdateRemainingNewParamsOp string

const (
	V1KeyUpdateRemainingNewParamsOpIncrement V1KeyUpdateRemainingNewParamsOp = "increment"
	V1KeyUpdateRemainingNewParamsOpDecrement V1KeyUpdateRemainingNewParamsOp = "decrement"
	V1KeyUpdateRemainingNewParamsOpSet       V1KeyUpdateRemainingNewParamsOp = "set"
)

func (r V1KeyUpdateRemainingNewParamsOp) IsKnown() bool {
	switch r {
	case V1KeyUpdateRemainingNewParamsOpIncrement, V1KeyUpdateRemainingNewParamsOpDecrement, V1KeyUpdateRemainingNewParamsOpSet:
		return true
	}
	return false
}
