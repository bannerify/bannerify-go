// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsUnion()           {}
func (UnionString) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndUnion()     {}
func (UnionString) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsAndAndOrOrUnion() {}
func (UnionString) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrUnion()       {}
func (UnionString) ImplementsV1KeyVerifyKeyNewParamsAuthorizationPermissionsOrOrAndAndUnion() {}
func (UnionString) ImplementsRatelimitLimitParamsMetaUnion()                                  {}
func (UnionString) ImplementsRatelimitLimitParamsResourcesMetaUnion()                         {}

type UnionBool bool

func (UnionBool) ImplementsRatelimitLimitParamsMetaUnion()          {}
func (UnionBool) ImplementsRatelimitLimitParamsResourcesMetaUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsRatelimitLimitParamsMetaUnion()          {}
func (UnionFloat) ImplementsRatelimitLimitParamsResourcesMetaUnion() {}
