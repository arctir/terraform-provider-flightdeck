// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"encoding/json"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	pluginconfigurationresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_plugin_configuration"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
// goverter:extend StringValueToJSON
// goverter:extend JSONToStringValue
// goverter:extend PtrJSONToStringValue
// goverter:extend StringValueToPtrJSON
// goverter:extend BoolValueToBool
// goverter:extend BoolToBoolValue
type PluginConfigurationConverter interface {
	ConvertToApiPluginConfiguration(source pluginconfigurationresource.PluginConfigurationModel) (flightdeckv1.PluginConfigurationInput, error)
	// goverter:map Definition | PluginConfigurationDefinitionSpecToDefinitionValue
	ConvertToTfPluginConfiguration(source flightdeckv1.PluginConfiguration) (pluginconfigurationresource.PluginConfigurationModel, error)
}

func StringValueToJSON(source basetypes.StringValue) (map[string]interface{}, error) {
	var out map[string]interface{}
	err := json.Unmarshal([]byte(source.ValueString()), &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func JSONToStringValue(source map[string]interface{}) basetypes.StringValue {
	b, err := json.Marshal(source)
	if err != nil {
		return basetypes.NewStringValue("")
	}
	return basetypes.NewStringValue(string(b))
}

func PtrJSONToStringValue(source *map[string]interface{}) basetypes.StringValue {
	if source == nil {
		return basetypes.NewStringValue("")
	}
	return JSONToStringValue(*source)
}

func StringValueToPtrJSON(source basetypes.StringValue) *map[string]interface{} {
	if source.ValueString() == "" {
		return nil
	}
	out, err := StringValueToJSON(source)
	if err != nil {
		return nil
	}
	return &out
}

func PluginConfigurationDefinitionSpecToDefinitionValue(source flightdeckv1.PluginConfigurationDefinitionSpec) (pluginconfigurationresource.DefinitionValue, error) {
	attrTypes := pluginconfigurationresource.DefinitionValue{}.AttributeTypes(context.Background())
	val, diag := pluginconfigurationresource.NewDefinitionValue(attrTypes, map[string]attr.Value{
		"name":              basetypes.NewStringValue(source.Name),
		"portal_version_id": basetypes.NewStringValue(source.PortalVersionId),
	})
	if diag.HasError() {
		return pluginconfigurationresource.DefinitionValue{}, fmt.Errorf("failed to create definition value")
	}
	return val, nil
}
