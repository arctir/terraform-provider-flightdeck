// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	clusterdatasource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_cluster"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type ClusterConverter interface {
	ConvertToTfCluster(source flightdeckv1.Cluster) clusterdatasource.ClusterModel
}
