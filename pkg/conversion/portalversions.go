// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_portal_version"
)

// goverter:converter
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend IntToInt64Value
type PortalVersionConverter interface {
	ConvertToTfPortalVersion(source flightdeckv1.PortalVersion) datasource_portal_version.PortalVersionModel
}
