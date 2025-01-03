// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_plugin_definition

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func PluginDefinitionResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"backend": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"config_schema": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"packages": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
					},
					"ui_schema": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
				},
				CustomType: BackendType{
					ObjectType: types.ObjectType{
						AttrTypes: BackendValue{}.AttributeTypes(ctx),
					},
				},
				Optional: true,
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"display_name": schema.StringAttribute{
				Required: true,
			},
			"frontend": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"components": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_schema": schema.StringAttribute{
									Optional: true,
									Computed: true,
								},
								"description": schema.StringAttribute{
									Required: true,
								},
								"path": schema.StringAttribute{
									Required: true,
								},
								"title": schema.StringAttribute{
									Required: true,
								},
								"type": schema.StringAttribute{
									Required: true,
								},
							},
							CustomType: ComponentsType{
								ObjectType: types.ObjectType{
									AttrTypes: ComponentsValue{}.AttributeTypes(ctx),
								},
							},
						},
						Optional: true,
						Computed: true,
					},
					"packages": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
					},
				},
				CustomType: FrontendType{
					ObjectType: types.ObjectType{
						AttrTypes: FrontendValue{}.AttributeTypes(ctx),
					},
				},
				Optional: true,
				Computed: true,
			},
			"icon_name": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"organization_id": schema.StringAttribute{
				Required: true,
			},
			"version": schema.Int64Attribute{
				Required: true,
			},
		},
	}
}

type PluginDefinitionModel struct {
	Backend        BackendValue  `tfsdk:"backend"`
	CreatedAt      types.String  `tfsdk:"created_at"`
	Description    types.String  `tfsdk:"description"`
	DisplayName    types.String  `tfsdk:"display_name"`
	Frontend       FrontendValue `tfsdk:"frontend"`
	IconName       types.String  `tfsdk:"icon_name"`
	Id             types.String  `tfsdk:"id"`
	Name           types.String  `tfsdk:"name"`
	OrganizationId types.String  `tfsdk:"organization_id"`
	Version        types.Int64   `tfsdk:"version"`
}

var _ basetypes.ObjectTypable = BackendType{}

type BackendType struct {
	basetypes.ObjectType
}

func (t BackendType) Equal(o attr.Type) bool {
	other, ok := o.(BackendType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t BackendType) String() string {
	return "BackendType"
}

func (t BackendType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	configSchemaAttribute, ok := attributes["config_schema"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`config_schema is missing from object`)

		return nil, diags
	}

	configSchemaVal, ok := configSchemaAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`config_schema expected to be basetypes.StringValue, was: %T`, configSchemaAttribute))
	}

	packagesAttribute, ok := attributes["packages"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`packages is missing from object`)

		return nil, diags
	}

	packagesVal, ok := packagesAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`packages expected to be basetypes.ListValue, was: %T`, packagesAttribute))
	}

	uiSchemaAttribute, ok := attributes["ui_schema"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ui_schema is missing from object`)

		return nil, diags
	}

	uiSchemaVal, ok := uiSchemaAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ui_schema expected to be basetypes.StringValue, was: %T`, uiSchemaAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return BackendValue{
		ConfigSchema: configSchemaVal,
		Packages:     packagesVal,
		UiSchema:     uiSchemaVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewBackendValueNull() BackendValue {
	return BackendValue{
		state: attr.ValueStateNull,
	}
}

func NewBackendValueUnknown() BackendValue {
	return BackendValue{
		state: attr.ValueStateUnknown,
	}
}

func NewBackendValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (BackendValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing BackendValue Attribute Value",
				"While creating a BackendValue value, a missing attribute value was detected. "+
					"A BackendValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("BackendValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid BackendValue Attribute Type",
				"While creating a BackendValue value, an invalid attribute value was detected. "+
					"A BackendValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("BackendValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("BackendValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra BackendValue Attribute Value",
				"While creating a BackendValue value, an extra attribute value was detected. "+
					"A BackendValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra BackendValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewBackendValueUnknown(), diags
	}

	configSchemaAttribute, ok := attributes["config_schema"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`config_schema is missing from object`)

		return NewBackendValueUnknown(), diags
	}

	configSchemaVal, ok := configSchemaAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`config_schema expected to be basetypes.StringValue, was: %T`, configSchemaAttribute))
	}

	packagesAttribute, ok := attributes["packages"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`packages is missing from object`)

		return NewBackendValueUnknown(), diags
	}

	packagesVal, ok := packagesAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`packages expected to be basetypes.ListValue, was: %T`, packagesAttribute))
	}

	uiSchemaAttribute, ok := attributes["ui_schema"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ui_schema is missing from object`)

		return NewBackendValueUnknown(), diags
	}

	uiSchemaVal, ok := uiSchemaAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ui_schema expected to be basetypes.StringValue, was: %T`, uiSchemaAttribute))
	}

	if diags.HasError() {
		return NewBackendValueUnknown(), diags
	}

	return BackendValue{
		ConfigSchema: configSchemaVal,
		Packages:     packagesVal,
		UiSchema:     uiSchemaVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewBackendValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) BackendValue {
	object, diags := NewBackendValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewBackendValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t BackendType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewBackendValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewBackendValueUnknown(), nil
	}

	if in.IsNull() {
		return NewBackendValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewBackendValueMust(BackendValue{}.AttributeTypes(ctx), attributes), nil
}

func (t BackendType) ValueType(ctx context.Context) attr.Value {
	return BackendValue{}
}

var _ basetypes.ObjectValuable = BackendValue{}

type BackendValue struct {
	ConfigSchema basetypes.StringValue `tfsdk:"config_schema"`
	Packages     basetypes.ListValue   `tfsdk:"packages"`
	UiSchema     basetypes.StringValue `tfsdk:"ui_schema"`
	state        attr.ValueState
}

func (v BackendValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["config_schema"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["packages"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["ui_schema"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.ConfigSchema.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["config_schema"] = val

		val, err = v.Packages.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["packages"] = val

		val, err = v.UiSchema.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ui_schema"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v BackendValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v BackendValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v BackendValue) String() string {
	return "BackendValue"
}

func (v BackendValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var packagesVal basetypes.ListValue
	switch {
	case v.Packages.IsUnknown():
		packagesVal = types.ListUnknown(types.StringType)
	case v.Packages.IsNull():
		packagesVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		packagesVal, d = types.ListValue(types.StringType, v.Packages.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"config_schema": basetypes.StringType{},
			"packages": basetypes.ListType{
				ElemType: types.StringType,
			},
			"ui_schema": basetypes.StringType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"config_schema": basetypes.StringType{},
		"packages": basetypes.ListType{
			ElemType: types.StringType,
		},
		"ui_schema": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"config_schema": v.ConfigSchema,
			"packages":      packagesVal,
			"ui_schema":     v.UiSchema,
		})

	return objVal, diags
}

func (v BackendValue) Equal(o attr.Value) bool {
	other, ok := o.(BackendValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.ConfigSchema.Equal(other.ConfigSchema) {
		return false
	}

	if !v.Packages.Equal(other.Packages) {
		return false
	}

	if !v.UiSchema.Equal(other.UiSchema) {
		return false
	}

	return true
}

func (v BackendValue) Type(ctx context.Context) attr.Type {
	return BackendType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v BackendValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"config_schema": basetypes.StringType{},
		"packages": basetypes.ListType{
			ElemType: types.StringType,
		},
		"ui_schema": basetypes.StringType{},
	}
}

var _ basetypes.ObjectTypable = FrontendType{}

type FrontendType struct {
	basetypes.ObjectType
}

func (t FrontendType) Equal(o attr.Type) bool {
	other, ok := o.(FrontendType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t FrontendType) String() string {
	return "FrontendType"
}

func (t FrontendType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	componentsAttribute, ok := attributes["components"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`components is missing from object`)

		return nil, diags
	}

	componentsVal, ok := componentsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`components expected to be basetypes.ListValue, was: %T`, componentsAttribute))
	}

	packagesAttribute, ok := attributes["packages"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`packages is missing from object`)

		return nil, diags
	}

	packagesVal, ok := packagesAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`packages expected to be basetypes.ListValue, was: %T`, packagesAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return FrontendValue{
		Components: componentsVal,
		Packages:   packagesVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewFrontendValueNull() FrontendValue {
	return FrontendValue{
		state: attr.ValueStateNull,
	}
}

func NewFrontendValueUnknown() FrontendValue {
	return FrontendValue{
		state: attr.ValueStateUnknown,
	}
}

func NewFrontendValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (FrontendValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing FrontendValue Attribute Value",
				"While creating a FrontendValue value, a missing attribute value was detected. "+
					"A FrontendValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("FrontendValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid FrontendValue Attribute Type",
				"While creating a FrontendValue value, an invalid attribute value was detected. "+
					"A FrontendValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("FrontendValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("FrontendValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra FrontendValue Attribute Value",
				"While creating a FrontendValue value, an extra attribute value was detected. "+
					"A FrontendValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra FrontendValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewFrontendValueUnknown(), diags
	}

	componentsAttribute, ok := attributes["components"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`components is missing from object`)

		return NewFrontendValueUnknown(), diags
	}

	componentsVal, ok := componentsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`components expected to be basetypes.ListValue, was: %T`, componentsAttribute))
	}

	packagesAttribute, ok := attributes["packages"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`packages is missing from object`)

		return NewFrontendValueUnknown(), diags
	}

	packagesVal, ok := packagesAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`packages expected to be basetypes.ListValue, was: %T`, packagesAttribute))
	}

	if diags.HasError() {
		return NewFrontendValueUnknown(), diags
	}

	return FrontendValue{
		Components: componentsVal,
		Packages:   packagesVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewFrontendValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) FrontendValue {
	object, diags := NewFrontendValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewFrontendValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t FrontendType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewFrontendValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewFrontendValueUnknown(), nil
	}

	if in.IsNull() {
		return NewFrontendValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewFrontendValueMust(FrontendValue{}.AttributeTypes(ctx), attributes), nil
}

func (t FrontendType) ValueType(ctx context.Context) attr.Value {
	return FrontendValue{}
}

var _ basetypes.ObjectValuable = FrontendValue{}

type FrontendValue struct {
	Components basetypes.ListValue `tfsdk:"components"`
	Packages   basetypes.ListValue `tfsdk:"packages"`
	state      attr.ValueState
}

func (v FrontendValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["components"] = basetypes.ListType{
		ElemType: ComponentsValue{}.Type(ctx),
	}.TerraformType(ctx)
	attrTypes["packages"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Components.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["components"] = val

		val, err = v.Packages.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["packages"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v FrontendValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v FrontendValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v FrontendValue) String() string {
	return "FrontendValue"
}

func (v FrontendValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	components := types.ListValueMust(
		ComponentsType{
			basetypes.ObjectType{
				AttrTypes: ComponentsValue{}.AttributeTypes(ctx),
			},
		},
		v.Components.Elements(),
	)

	if v.Components.IsNull() {
		components = types.ListNull(
			ComponentsType{
				basetypes.ObjectType{
					AttrTypes: ComponentsValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	if v.Components.IsUnknown() {
		components = types.ListUnknown(
			ComponentsType{
				basetypes.ObjectType{
					AttrTypes: ComponentsValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	var packagesVal basetypes.ListValue
	switch {
	case v.Packages.IsUnknown():
		packagesVal = types.ListUnknown(types.StringType)
	case v.Packages.IsNull():
		packagesVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		packagesVal, d = types.ListValue(types.StringType, v.Packages.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"components": basetypes.ListType{
				ElemType: ComponentsValue{}.Type(ctx),
			},
			"packages": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"components": basetypes.ListType{
			ElemType: ComponentsValue{}.Type(ctx),
		},
		"packages": basetypes.ListType{
			ElemType: types.StringType,
		},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"components": components,
			"packages":   packagesVal,
		})

	return objVal, diags
}

func (v FrontendValue) Equal(o attr.Value) bool {
	other, ok := o.(FrontendValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Components.Equal(other.Components) {
		return false
	}

	if !v.Packages.Equal(other.Packages) {
		return false
	}

	return true
}

func (v FrontendValue) Type(ctx context.Context) attr.Type {
	return FrontendType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v FrontendValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"components": basetypes.ListType{
			ElemType: ComponentsValue{}.Type(ctx),
		},
		"packages": basetypes.ListType{
			ElemType: types.StringType,
		},
	}
}

var _ basetypes.ObjectTypable = ComponentsType{}

type ComponentsType struct {
	basetypes.ObjectType
}

func (t ComponentsType) Equal(o attr.Type) bool {
	other, ok := o.(ComponentsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ComponentsType) String() string {
	return "ComponentsType"
}

func (t ComponentsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	configSchemaAttribute, ok := attributes["config_schema"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`config_schema is missing from object`)

		return nil, diags
	}

	configSchemaVal, ok := configSchemaAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`config_schema expected to be basetypes.StringValue, was: %T`, configSchemaAttribute))
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return nil, diags
	}

	descriptionVal, ok := descriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be basetypes.StringValue, was: %T`, descriptionAttribute))
	}

	pathAttribute, ok := attributes["path"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`path is missing from object`)

		return nil, diags
	}

	pathVal, ok := pathAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`path expected to be basetypes.StringValue, was: %T`, pathAttribute))
	}

	titleAttribute, ok := attributes["title"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`title is missing from object`)

		return nil, diags
	}

	titleVal, ok := titleAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`title expected to be basetypes.StringValue, was: %T`, titleAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return nil, diags
	}

	typeVal, ok := typeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, typeAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ComponentsValue{
		ConfigSchema:   configSchemaVal,
		Description:    descriptionVal,
		Path:           pathVal,
		Title:          titleVal,
		ComponentsType: typeVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewComponentsValueNull() ComponentsValue {
	return ComponentsValue{
		state: attr.ValueStateNull,
	}
}

func NewComponentsValueUnknown() ComponentsValue {
	return ComponentsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewComponentsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ComponentsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ComponentsValue Attribute Value",
				"While creating a ComponentsValue value, a missing attribute value was detected. "+
					"A ComponentsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ComponentsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ComponentsValue Attribute Type",
				"While creating a ComponentsValue value, an invalid attribute value was detected. "+
					"A ComponentsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ComponentsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ComponentsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ComponentsValue Attribute Value",
				"While creating a ComponentsValue value, an extra attribute value was detected. "+
					"A ComponentsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ComponentsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewComponentsValueUnknown(), diags
	}

	configSchemaAttribute, ok := attributes["config_schema"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`config_schema is missing from object`)

		return NewComponentsValueUnknown(), diags
	}

	configSchemaVal, ok := configSchemaAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`config_schema expected to be basetypes.StringValue, was: %T`, configSchemaAttribute))
	}

	descriptionAttribute, ok := attributes["description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`description is missing from object`)

		return NewComponentsValueUnknown(), diags
	}

	descriptionVal, ok := descriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`description expected to be basetypes.StringValue, was: %T`, descriptionAttribute))
	}

	pathAttribute, ok := attributes["path"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`path is missing from object`)

		return NewComponentsValueUnknown(), diags
	}

	pathVal, ok := pathAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`path expected to be basetypes.StringValue, was: %T`, pathAttribute))
	}

	titleAttribute, ok := attributes["title"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`title is missing from object`)

		return NewComponentsValueUnknown(), diags
	}

	titleVal, ok := titleAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`title expected to be basetypes.StringValue, was: %T`, titleAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return NewComponentsValueUnknown(), diags
	}

	typeVal, ok := typeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, typeAttribute))
	}

	if diags.HasError() {
		return NewComponentsValueUnknown(), diags
	}

	return ComponentsValue{
		ConfigSchema:   configSchemaVal,
		Description:    descriptionVal,
		Path:           pathVal,
		Title:          titleVal,
		ComponentsType: typeVal,
		state:          attr.ValueStateKnown,
	}, diags
}

func NewComponentsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ComponentsValue {
	object, diags := NewComponentsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewComponentsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ComponentsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewComponentsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewComponentsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewComponentsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewComponentsValueMust(ComponentsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ComponentsType) ValueType(ctx context.Context) attr.Value {
	return ComponentsValue{}
}

var _ basetypes.ObjectValuable = ComponentsValue{}

type ComponentsValue struct {
	ConfigSchema   basetypes.StringValue `tfsdk:"config_schema"`
	Description    basetypes.StringValue `tfsdk:"description"`
	Path           basetypes.StringValue `tfsdk:"path"`
	Title          basetypes.StringValue `tfsdk:"title"`
	ComponentsType basetypes.StringValue `tfsdk:"type"`
	state          attr.ValueState
}

func (v ComponentsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 5)

	var val tftypes.Value
	var err error

	attrTypes["config_schema"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["description"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["path"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["title"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["type"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 5)

		val, err = v.ConfigSchema.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["config_schema"] = val

		val, err = v.Description.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["description"] = val

		val, err = v.Path.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["path"] = val

		val, err = v.Title.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["title"] = val

		val, err = v.ComponentsType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["type"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ComponentsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ComponentsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ComponentsValue) String() string {
	return "ComponentsValue"
}

func (v ComponentsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"config_schema": basetypes.StringType{},
		"description":   basetypes.StringType{},
		"path":          basetypes.StringType{},
		"title":         basetypes.StringType{},
		"type":          basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"config_schema": v.ConfigSchema,
			"description":   v.Description,
			"path":          v.Path,
			"title":         v.Title,
			"type":          v.ComponentsType,
		})

	return objVal, diags
}

func (v ComponentsValue) Equal(o attr.Value) bool {
	other, ok := o.(ComponentsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.ConfigSchema.Equal(other.ConfigSchema) {
		return false
	}

	if !v.Description.Equal(other.Description) {
		return false
	}

	if !v.Path.Equal(other.Path) {
		return false
	}

	if !v.Title.Equal(other.Title) {
		return false
	}

	if !v.ComponentsType.Equal(other.ComponentsType) {
		return false
	}

	return true
}

func (v ComponentsValue) Type(ctx context.Context) attr.Type {
	return ComponentsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ComponentsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"config_schema": basetypes.StringType{},
		"description":   basetypes.StringType{},
		"path":          basetypes.StringType{},
		"title":         basetypes.StringType{},
		"type":          basetypes.StringType{},
	}
}
