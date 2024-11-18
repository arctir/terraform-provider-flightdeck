// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	tenantuserresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_tenant_user"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	oapitypes "github.com/oapi-codegen/runtime/types"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
// goverter:extend StringValueToEmail
// goverter:extend EmailToTfStringValue
type TenantUserConverter interface {
	ConvertToApiTenantUser(source tenantuserresource.TenantUserModel) flightdeckv1.TenantUserInput
	ConvertToTfTenantUser(source flightdeckv1.TenantUser) tenantuserresource.TenantUserModel
}

func StringValueToEmail(input basetypes.StringValue) oapitypes.Email {
	return oapitypes.Email(input.ValueString())
}

func EmailToTfStringValue(input oapitypes.Email) basetypes.StringValue {
	return basetypes.NewStringValue(string(input))
}
