// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	orgdatasource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_organization"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type OrganizationConverter interface {
	// goverter:ignore Subscription
	ConvertToTfOrganization(source flightdeckv1.Organization) orgdatasource.OrganizationModel
}
