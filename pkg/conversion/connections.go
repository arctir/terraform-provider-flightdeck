// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	connectionresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_connection"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type ConnectionConverter interface {
	// goverter:map . ConnectionConfig | SetApiConnectionConfig
	ConvertToApiConnection(source connectionresource.ConnectionModel) (*flightdeckv1.ConnectionInput, error)
	// goverter:ignoreMissing
	// goverter:default SetTfConnectionConfig
	// goverter:ignore Tailscale
	ConvertToTfConnection(source flightdeckv1.Connection) (connectionresource.ConnectionModel, error)
}

func SetApiConnectionConfig(c ConnectionConverter, source connectionresource.ConnectionModel) (flightdeckv1.ConnectionInput_ConnectionConfig, error) {
	config := flightdeckv1.ConnectionInput_ConnectionConfig{}

	/*
		if source.Tailscale.IsUnknown() {
			return config, nil
		}*/
	if !source.Tailscale.IsUnknown() {
		var tc flightdeckv1.TailscaleConnection
		o, diag := source.Tailscale.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse tailscale connection %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &tc, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse tailscale connection %+v", diag.Errors())
		}
		tc.ConfigType = "tailscale"
		config.FromTailscaleConnection(tc)
	}
	return config, nil
}

func SetTfConnectionConfig(c ConnectionConverter, input flightdeckv1.Connection) (connectionresource.ConnectionModel, error) {
	connection := connectionresource.ConnectionModel{
		OrganizationId: basetypes.NewStringValue(input.OrganizationId.String()),
		PortalName:     basetypes.NewStringValue(input.PortalName),
		Tailscale:      connectionresource.NewTailscaleValueNull(),
	}

	tailscaleTypes := connectionresource.TailscaleValue{}.AttributeTypes(context.TODO())

	configType, err := input.ConnectionConfig.Discriminator()
	if err != nil {
		return connection, err
	}
	switch configType {
	case "tailscale":
		tc, err := input.ConnectionConfig.AsTailscaleConnection()
		if err != nil {
			return connection, err
		}
		tailscale, diag := basetypes.NewObjectValueFrom(context.TODO(), tailscaleTypes, tc)
		if diag.HasError() {
			return connection, fmt.Errorf("could not parse tailscale connection %+v", diag.Errors())
		}
		tsv, diag := connectionresource.NewTailscaleValue(tailscaleTypes, tailscale.Attributes())
		if diag.HasError() {
			return connection, fmt.Errorf("could not parse tailscale connection %+v", diag.Errors())
		}
		connection.Tailscale = tsv
	}
	return connection, nil
}
