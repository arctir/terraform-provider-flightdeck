// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"encoding/json"
	"errors"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	entitypagelayoutresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_entity_page_layout"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend BoolValueToPtrBool
// goverter:extend PtrBoolToBoolValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
// goverter:extend ListValueToSliceEntityPageCardSpec
// goverter:extend ListValueToSliceEntityPageContentSpec
// goverter:extend SliceEntityPageCardSpecToListValue
// goverter:extend SliceEntityPageContentSpecToListValue
type EntityPageLayoutConverter interface {
	ConvertToApiEntityPageLayout(source entitypagelayoutresource.EntityPageLayoutModel) (flightdeckv1.EntityPageLayoutInput, error)
	// goverter:map CardOrder | SliceEntityPageCardSpecToListValue
	// goverter:map ContentOrder | SliceEntityPageContentSpecToListValue
	ConvertToTfEntityPageLayout(source flightdeckv1.EntityPageLayout) (entitypagelayoutresource.EntityPageLayoutModel, error)
}

func ListValueToSliceEntityPageCardSpec(source basetypes.ListValue) ([]flightdeckv1.EntityPageCardSpec, error) {
	out := make([]flightdeckv1.EntityPageCardSpec, 0, len(source.Elements()))
	for _, element := range source.Elements() {
		v, ok := element.(entitypagelayoutresource.CardOrderValue)
		if ok {
			var config map[string]interface{}
			configData := v.Config.ValueString()
			if configData != "" {
				err := json.Unmarshal([]byte(v.Config.ValueString()), &config)
				if err != nil {
					return nil, err
				}
			}
			out = append(out, flightdeckv1.EntityPageCardSpec{
				Path:   v.Path.ValueString(),
				Config: &config,
			})
		}
	}
	return out, nil
}

func ListValueToSliceEntityPageContentSpec(source basetypes.ListValue) ([]flightdeckv1.EntityPageContentSpec, error) {
	out := make([]flightdeckv1.EntityPageContentSpec, 0, len(source.Elements()))
	for _, element := range source.Elements() {
		v, ok := element.(entitypagelayoutresource.ContentOrderValue)
		if ok {
			var config map[string]interface{}
			configData := v.Config.ValueString()
			if configData != "" {
				err := json.Unmarshal([]byte(v.Config.ValueString()), &config)
				if err != nil {
					return nil, err
				}
			}
			out = append(out, flightdeckv1.EntityPageContentSpec{
				Path:   v.Path.ValueString(),
				Config: &config,
			})
		}
	}
	return out, nil
}

func SliceEntityPageCardSpecToListValue(source *[]flightdeckv1.EntityPageCardSpec) (basetypes.ListValue, error) {
	types := entitypagelayoutresource.CardOrderValue{}.Type(context.TODO())
	if source == nil {
		return basetypes.NewListNull(types), nil
	}

	out := make([]entitypagelayoutresource.CardOrderValue, 0, len(*source))
	for _, i := range *source {
		var config []byte
		var err error
		if i.Config != nil {
			config, err = json.Marshal(i.Config)
			if err != nil {
				return basetypes.NewListNull(types), err
			}
		}
		elem, diag := entitypagelayoutresource.NewCardOrderValue(
			entitypagelayoutresource.CardOrderValue{}.AttributeTypes(context.TODO()),
			map[string]attr.Value{
				"path":    basetypes.NewStringValue(i.Path),
				"filters": basetypes.NewListNull(basetypes.StringType{}),
				"config":  basetypes.NewStringValue(string(config)),
			})
		if diag.HasError() {
			return basetypes.NewListNull(types), nil
		}
		out = append(out, elem)
	}
	val, diag := basetypes.NewListValueFrom(context.TODO(), types, out)
	if diag.HasError() {
		return val, errors.New("could not convert entitypagecardspec slice to list value")
	}
	return val, nil
}

func SliceEntityPageContentSpecToListValue(source *[]flightdeckv1.EntityPageContentSpec) (basetypes.ListValue, error) {
	types := entitypagelayoutresource.ContentOrderValue{}.Type(context.TODO())
	if source == nil {
		return basetypes.NewListNull(types), nil
	}

	out := make([]entitypagelayoutresource.ContentOrderValue, 0, len(*source))
	for _, i := range *source {
		var config []byte
		var err error
		if i.Config != nil {
			config, err = json.Marshal(i.Config)
			if err != nil {
				return basetypes.NewListNull(types), err
			}
		}
		elem, diag := entitypagelayoutresource.NewContentOrderValue(
			entitypagelayoutresource.ContentOrderValue{}.AttributeTypes(context.TODO()),
			map[string]attr.Value{
				"path":    basetypes.NewStringValue(i.Path),
				"filters": basetypes.NewListNull(basetypes.StringType{}),
				"config":  basetypes.NewStringValue(string(config)),
			})
		if diag.HasError() {
			return basetypes.NewListNull(types), nil
		}
		out = append(out, elem)
	}
	val, diag := basetypes.NewListValueFrom(context.TODO(), types, out)
	if diag.HasError() {
		return val, errors.New("could not convert entitypageContentspec slice to list value")
	}
	return val, nil
}
