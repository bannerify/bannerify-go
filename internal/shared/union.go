// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsRatelimitLimitParamsMetaUnion()          {}
func (UnionString) ImplementsRatelimitLimitParamsResourcesMetaUnion() {}

type UnionBool bool

func (UnionBool) ImplementsRatelimitLimitParamsMetaUnion()          {}
func (UnionBool) ImplementsRatelimitLimitParamsResourcesMetaUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsRatelimitLimitParamsMetaUnion()          {}
func (UnionFloat) ImplementsRatelimitLimitParamsResourcesMetaUnion() {}
