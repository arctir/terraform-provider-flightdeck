// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	tenantresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_tenant"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type TenantConverter interface {
	ConvertToApiTenant(source tenantresource.TenantModel) flightdeckv1.TenantInput
	ConvertToTfTenant(source flightdeckv1.Tenant) tenantresource.TenantModel
}
