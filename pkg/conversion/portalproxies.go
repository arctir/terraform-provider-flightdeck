// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"errors"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	portalproxyresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_portal_proxy"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend BoolValueToBool
// goverter:extend BoolToBoolValue
// goverter:extend BoolValueToPtrBool
// goverter:extend PtrBoolToBoolValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend ListValueStringToSliceString
// goverter:extend SliceStringToListValueString
// goverter:extend ListValueStringToPtrSliceString
// goverter:extend PtrSliceStringToListValueString
// goverter:extend StringValueToPortalProxyInputCredentials
// goverter:extend PortalProxyCredentialsToStringValue
// goverter:extend ListValueToSlicePortalProxyHeader
// goverter:extend ListValueToSlicePortalProxyPathRewrite
// goverter:extend SlicePortalProxyPathRewriteToListValue
type PortalProxyConverter interface {
	ConvertToApiPortalProxy(source portalproxyresource.PortalProxyModel) (flightdeckv1.PortalProxyInput, error)
	// goverter:map HttpHeaders | SlicePortalProxyHeadersToListValue
	// goverter:map PathRewrite | SlicePortalProxyPathRewriteToListValue
	ConvertToTfPortalProxy(source flightdeckv1.PortalProxy) (portalproxyresource.PortalProxyModel, error)
}

func StringValueToPortalProxyInputCredentials(source basetypes.StringValue) flightdeckv1.PortalProxyInputCredentials {
	return flightdeckv1.PortalProxyInputCredentials(source.ValueString())
}

func PortalProxyCredentialsToStringValue(source flightdeckv1.PortalProxyCredentials) basetypes.StringValue {
	return basetypes.NewStringValue(string(source))
}

func ListValueToSlicePortalProxyHeader(source basetypes.ListValue) ([]flightdeckv1.PortalProxyHeader, error) {
	out := make([]flightdeckv1.PortalProxyHeader, 0, len(source.Elements()))
	diag := source.ElementsAs(context.TODO(), &out, true)
	if diag.HasError() {
		return nil, errors.New("could not convert listvalue to portalproxyheader slice")
	}
	return out, nil
}

func ListValueToSlicePortalProxyPathRewrite(source basetypes.ListValue) ([]flightdeckv1.PortalProxyPathRewrite, error) {
	out := make([]flightdeckv1.PortalProxyPathRewrite, 0, len(source.Elements()))
	diag := source.ElementsAs(context.TODO(), &out, true)
	if diag.HasError() {
		return nil, errors.New("could not convert listvalue to portalproxypathrewrite slice")
	}
	return out, nil
}

func SlicePortalProxyHeadersToListValue(source *[]flightdeckv1.PortalProxyHeader) (basetypes.ListValue, error) {
	types := portalproxyresource.HttpHeadersValue{}.Type(context.TODO())
	if source == nil {
		return basetypes.NewListNull(types), nil
	}
	out, diag := basetypes.NewListValueFrom(context.TODO(), types, source)
	if diag.HasError() {
		return out, errors.New("could not convert portalproxyheader slice to list value")
	}
	return out, nil
}

func SlicePortalProxyPathRewriteToListValue(source *[]flightdeckv1.PortalProxyPathRewrite) (basetypes.ListValue, error) {
	types := portalproxyresource.PathRewriteValue{}.Type(context.TODO())
	if source == nil {
		return basetypes.NewListNull(types), nil
	}
	out, diag := basetypes.NewListValueFrom(context.TODO(), types, source)
	if diag.HasError() {
		return out, errors.New("could not convert portalproxypathrewrite slice to list value")
	}
	return out, nil
}
