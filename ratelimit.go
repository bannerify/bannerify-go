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

// RatelimitService contains methods and other services that help with interacting
// with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRatelimitService] method instead.
type RatelimitService struct {
	Options []option.RequestOption
}

// NewRatelimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatelimitService(opts ...option.RequestOption) (r *RatelimitService) {
	r = &RatelimitService{}
	r.Options = opts
	return
}

func (r *RatelimitService) Limit(ctx context.Context, body RatelimitLimitParams, opts ...option.RequestOption) (res *RatelimitLimitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/ratelimits.limit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RatelimitLimitResponse struct {
	// How many requests are allowed within a window.
	Limit int64 `json:"limit,required"`
	// How many requests can still be made in the current window.
	Remaining int64 `json:"remaining,required"`
	// A unix millisecond timestamp when the limits reset.
	Reset int64 `json:"reset,required"`
	// Returns true if the request should be processed, false if it was rejected.
	Success bool                       `json:"success,required"`
	JSON    ratelimitLimitResponseJSON `json:"-"`
}

// ratelimitLimitResponseJSON contains the JSON metadata for the struct
// [RatelimitLimitResponse]
type ratelimitLimitResponseJSON struct {
	Limit       apijson.Field
	Remaining   apijson.Field
	Reset       apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatelimitLimitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratelimitLimitResponseJSON) RawJSON() string {
	return r.raw
}

type RatelimitLimitParams struct {
	// The window duration in milliseconds
	Duration param.Field[int64] `json:"duration,required"`
	// Identifier of your user, this can be their userId, an email, an ip or anything
	// else.
	Identifier param.Field[string] `json:"identifier,required"`
	// How many requests may pass in a given window.
	Limit param.Field[int64] `json:"limit,required"`
	// Async will return a response immediately, lowering latency at the cost of
	// accuracy.
	Async param.Field[bool] `json:"async"`
	// Expensive requests may use up more tokens. You can specify a cost to the request
	// here and we'll deduct this many tokens in the current window. If there are not
	// enough tokens left, the request is denied.
	//
	// Set it to 0 to receive the current limit without changing anything.
	Cost param.Field[int64] `json:"cost"`
	// Attach any metadata to this request
	Meta param.Field[map[string]RatelimitLimitParamsMetaUnion] `json:"meta"`
	// Namespaces group different limits together for better analytics. You might have
	// a namespace for your public API and one for internal tRPC routes.
	Namespace param.Field[string] `json:"namespace"`
	// Resources that are about to be accessed by the user
	Resources param.Field[[]RatelimitLimitParamsResource] `json:"resources"`
}

func (r RatelimitLimitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type RatelimitLimitParamsMetaUnion interface {
	ImplementsRatelimitLimitParamsMetaUnion()
}

type RatelimitLimitParamsResource struct {
	// The unique identifier for the resource
	ID param.Field[string] `json:"id,required"`
	// The type of resource
	Type param.Field[string] `json:"type,required"`
	// Attach any metadata to this resources
	Meta param.Field[map[string]RatelimitLimitParamsResourcesMetaUnion] `json:"meta"`
	// A human readable name for this resource
	Name param.Field[string] `json:"name"`
}

func (r RatelimitLimitParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type RatelimitLimitParamsResourcesMetaUnion interface {
	ImplementsRatelimitLimitParamsResourcesMetaUnion()
}
