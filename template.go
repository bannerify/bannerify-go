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

// TemplateService contains methods and other services that help with interacting
// with the bannerify API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r *TemplateService) {
	r = &TemplateService{}
	r.Options = opts
	return
}

// Create an image from a template
func (r *TemplateService) NewImage(ctx context.Context, body TemplateNewImageParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	path := "v1/templates/createImage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a signed URL for a template
func (r *TemplateService) Signedurl(ctx context.Context, query TemplateSignedurlParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	path := "v1/templates/signedurl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TemplateNewImageParams struct {
	APIKey param.Field[string] `json:"apiKey,required"`
	// Your template id
	TemplateID param.Field[string] `json:"templateId,required"`
	// Only for debug purpose, it draws bounding box for each layer
	Debug         param.Field[string]                               `json:"_debug"`
	Format        param.Field[TemplateNewImageParamsFormat]         `json:"format"`
	Modifications param.Field[[]TemplateNewImageParamsModification] `json:"modifications"`
}

func (r TemplateNewImageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TemplateNewImageParamsFormat string

const (
	TemplateNewImageParamsFormatPng TemplateNewImageParamsFormat = "png"
	TemplateNewImageParamsFormatSvg TemplateNewImageParamsFormat = "svg"
)

func (r TemplateNewImageParamsFormat) IsKnown() bool {
	switch r {
	case TemplateNewImageParamsFormatPng, TemplateNewImageParamsFormatSvg:
		return true
	}
	return false
}

// A modification (aka override) to apply to the layer in image
type TemplateNewImageParamsModification struct {
	// The layer name of the modification
	Name param.Field[string] `json:"name,required"`
	// Modify the barcode layer content with this field
	Barcode param.Field[string] `json:"barcode"`
	// Update chart layer's data, follow chart.js data structure
	Chart param.Field[map[string]interface{}] `json:"chart"`
	// The color for the modification
	Color param.Field[string] `json:"color"`
	// Table columns
	Columns param.Field[[]string] `json:"columns"`
	// Table height mode
	HeightMode param.Field[TemplateNewImageParamsModificationsHeightMode] `json:"heightMode"`
	// Modify the qrcode layer content with this field
	Qrcode param.Field[string] `json:"qrcode"`
	// Table rows
	Rows param.Field[[]interface{}] `json:"rows"`
	// The source image for the modification
	Src param.Field[string] `json:"src"`
	// Star value
	Star param.Field[float64] `json:"star"`
	// You can modify the text layer with this field
	Text param.Field[string] `json:"text"`
	// Table theme
	Theme param.Field[TemplateNewImageParamsModificationsTheme] `json:"theme"`
	// Set the visibility of the field
	Visible param.Field[bool] `json:"visible"`
	// Table width mode
	WidthMode param.Field[TemplateNewImageParamsModificationsWidthMode] `json:"widthMode"`
}

func (r TemplateNewImageParamsModification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Table height mode
type TemplateNewImageParamsModificationsHeightMode string

const (
	TemplateNewImageParamsModificationsHeightModeStandard TemplateNewImageParamsModificationsHeightMode = "standard"
	TemplateNewImageParamsModificationsHeightModeAdaptive TemplateNewImageParamsModificationsHeightMode = "adaptive"
)

func (r TemplateNewImageParamsModificationsHeightMode) IsKnown() bool {
	switch r {
	case TemplateNewImageParamsModificationsHeightModeStandard, TemplateNewImageParamsModificationsHeightModeAdaptive:
		return true
	}
	return false
}

// Table theme
type TemplateNewImageParamsModificationsTheme string

const (
	TemplateNewImageParamsModificationsThemeNone     TemplateNewImageParamsModificationsTheme = "NONE"
	TemplateNewImageParamsModificationsThemeDefault  TemplateNewImageParamsModificationsTheme = "DEFAULT"
	TemplateNewImageParamsModificationsThemeBright   TemplateNewImageParamsModificationsTheme = "BRIGHT"
	TemplateNewImageParamsModificationsThemeSimplify TemplateNewImageParamsModificationsTheme = "SIMPLIFY"
	TemplateNewImageParamsModificationsThemeArco     TemplateNewImageParamsModificationsTheme = "ARCO"
)

func (r TemplateNewImageParamsModificationsTheme) IsKnown() bool {
	switch r {
	case TemplateNewImageParamsModificationsThemeNone, TemplateNewImageParamsModificationsThemeDefault, TemplateNewImageParamsModificationsThemeBright, TemplateNewImageParamsModificationsThemeSimplify, TemplateNewImageParamsModificationsThemeArco:
		return true
	}
	return false
}

// Table width mode
type TemplateNewImageParamsModificationsWidthMode string

const (
	TemplateNewImageParamsModificationsWidthModeStandard TemplateNewImageParamsModificationsWidthMode = "standard"
	TemplateNewImageParamsModificationsWidthModeAdaptive TemplateNewImageParamsModificationsWidthMode = "adaptive"
)

func (r TemplateNewImageParamsModificationsWidthMode) IsKnown() bool {
	switch r {
	case TemplateNewImageParamsModificationsWidthModeStandard, TemplateNewImageParamsModificationsWidthModeAdaptive:
		return true
	}
	return false
}

type TemplateSignedurlParams struct {
	// SHA256 hash of the query params, read more at
	// https://documentation.bannerify.co/api#signing-requests
	Sign param.Field[string] `query:"sign,required"`
	// Your template id
	TemplateID param.Field[string] `query:"templateId,required"`
	// Only for debug purpose, it draws bounding box for each layer
	Debug param.Field[string] `query:"_debug"`
	// Sha256 hash of the API key (use this)
	APIKeyHashed param.Field[string] `query:"apiKeyHashed"`
	// MD5 hash of the API key
	APIKeyMd5 param.Field[string]                        `query:"apiKeyMd5"`
	Format    param.Field[TemplateSignedurlParamsFormat] `query:"format"`
	// A JSON string of modifications object
	Modifications param.Field[string] `query:"modifications"`
	// By default, we cache the image in the CDN for 1 day to save your bandwidth, use
	// this field to disable cache so you can get the latest image
	Nocache param.Field[string] `query:"nocache"`
}

// URLQuery serializes [TemplateSignedurlParams]'s query parameters as
// `url.Values`.
func (r TemplateSignedurlParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TemplateSignedurlParamsFormat string

const (
	TemplateSignedurlParamsFormatPng TemplateSignedurlParamsFormat = "png"
	TemplateSignedurlParamsFormatSvg TemplateSignedurlParamsFormat = "svg"
)

func (r TemplateSignedurlParamsFormat) IsKnown() bool {
	switch r {
	case TemplateSignedurlParamsFormatPng, TemplateSignedurlParamsFormatSvg:
		return true
	}
	return false
}
