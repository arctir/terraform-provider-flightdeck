// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	portalresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_portal"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
// goverter:extend ListValueStringToSliceString
// goverter:extend SliceStringToListValueString
// goverter:useZeroValueOnPointerInconsistency
type PortalConverter interface {
	ConvertToApiPortal(source portalresource.PortalModel) (flightdeckv1.PortalInput, error)
	// goverter:ignore Status
	ConvertToTfPortal(source flightdeckv1.Portal) (portalresource.PortalModel, error)
}
