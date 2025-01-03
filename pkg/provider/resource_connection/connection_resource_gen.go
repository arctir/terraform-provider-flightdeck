// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_connection

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

func ConnectionResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed:            true,
				Description:         "The date and time of the resources creation.",
				MarkdownDescription: "The date and time of the resources creation.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of the Flightdeck resource.",
				MarkdownDescription: "The ID of the Flightdeck resource.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the Flightdeck Connection resource.",
				MarkdownDescription: "The name of the Flightdeck Connection resource.",
			},
			"organization_id": schema.StringAttribute{
				Required:            true,
				Description:         "The ID of the Flightdeck Organization resource.",
				MarkdownDescription: "The ID of the Flightdeck Organization resource.",
			},
			"portal_name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the Flightdeck Portal resource.",
				MarkdownDescription: "The name of the Flightdeck Portal resource.",
			},
			"tailscale": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"auth_token": schema.StringAttribute{
						Required:            true,
						Description:         "The Tailscale connection auth token.",
						MarkdownDescription: "The Tailscale connection auth token.",
					},
					"hosts": schema.ListAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						Computed:            true,
						Description:         "The Tailscale MagicDNS hosts to make available.",
						MarkdownDescription: "The Tailscale MagicDNS hosts to make available.",
					},
				},
				CustomType: TailscaleType{
					ObjectType: types.ObjectType{
						AttrTypes: TailscaleValue{}.AttributeTypes(ctx),
					},
				},
				Optional: true,
				Computed: true,
			},
		},
		Description: "Represents a Flightdeck Connection resource.",
	}
}

type ConnectionModel struct {
	CreatedAt      types.String   `tfsdk:"created_at"`
	Id             types.String   `tfsdk:"id"`
	Name           types.String   `tfsdk:"name"`
	OrganizationId types.String   `tfsdk:"organization_id"`
	PortalName     types.String   `tfsdk:"portal_name"`
	Tailscale      TailscaleValue `tfsdk:"tailscale"`
}

var _ basetypes.ObjectTypable = TailscaleType{}

type TailscaleType struct {
	basetypes.ObjectType
}

func (t TailscaleType) Equal(o attr.Type) bool {
	other, ok := o.(TailscaleType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t TailscaleType) String() string {
	return "TailscaleType"
}

func (t TailscaleType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	authTokenAttribute, ok := attributes["auth_token"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`auth_token is missing from object`)

		return nil, diags
	}

	authTokenVal, ok := authTokenAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`auth_token expected to be basetypes.StringValue, was: %T`, authTokenAttribute))
	}

	hostsAttribute, ok := attributes["hosts"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`hosts is missing from object`)

		return nil, diags
	}

	hostsVal, ok := hostsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`hosts expected to be basetypes.ListValue, was: %T`, hostsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return TailscaleValue{
		AuthToken: authTokenVal,
		Hosts:     hostsVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewTailscaleValueNull() TailscaleValue {
	return TailscaleValue{
		state: attr.ValueStateNull,
	}
}

func NewTailscaleValueUnknown() TailscaleValue {
	return TailscaleValue{
		state: attr.ValueStateUnknown,
	}
}

func NewTailscaleValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (TailscaleValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing TailscaleValue Attribute Value",
				"While creating a TailscaleValue value, a missing attribute value was detected. "+
					"A TailscaleValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("TailscaleValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid TailscaleValue Attribute Type",
				"While creating a TailscaleValue value, an invalid attribute value was detected. "+
					"A TailscaleValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("TailscaleValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("TailscaleValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra TailscaleValue Attribute Value",
				"While creating a TailscaleValue value, an extra attribute value was detected. "+
					"A TailscaleValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra TailscaleValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewTailscaleValueUnknown(), diags
	}

	authTokenAttribute, ok := attributes["auth_token"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`auth_token is missing from object`)

		return NewTailscaleValueUnknown(), diags
	}

	authTokenVal, ok := authTokenAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`auth_token expected to be basetypes.StringValue, was: %T`, authTokenAttribute))
	}

	hostsAttribute, ok := attributes["hosts"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`hosts is missing from object`)

		return NewTailscaleValueUnknown(), diags
	}

	hostsVal, ok := hostsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`hosts expected to be basetypes.ListValue, was: %T`, hostsAttribute))
	}

	if diags.HasError() {
		return NewTailscaleValueUnknown(), diags
	}

	return TailscaleValue{
		AuthToken: authTokenVal,
		Hosts:     hostsVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewTailscaleValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) TailscaleValue {
	object, diags := NewTailscaleValue(attributeTypes, attributes)

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

		panic("NewTailscaleValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t TailscaleType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewTailscaleValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewTailscaleValueUnknown(), nil
	}

	if in.IsNull() {
		return NewTailscaleValueNull(), nil
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

	return NewTailscaleValueMust(TailscaleValue{}.AttributeTypes(ctx), attributes), nil
}

func (t TailscaleType) ValueType(ctx context.Context) attr.Value {
	return TailscaleValue{}
}

var _ basetypes.ObjectValuable = TailscaleValue{}

type TailscaleValue struct {
	AuthToken basetypes.StringValue `tfsdk:"auth_token"`
	Hosts     basetypes.ListValue   `tfsdk:"hosts"`
	state     attr.ValueState
}

func (v TailscaleValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["auth_token"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["hosts"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.AuthToken.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["auth_token"] = val

		val, err = v.Hosts.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["hosts"] = val

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

func (v TailscaleValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v TailscaleValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v TailscaleValue) String() string {
	return "TailscaleValue"
}

func (v TailscaleValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var hostsVal basetypes.ListValue
	switch {
	case v.Hosts.IsUnknown():
		hostsVal = types.ListUnknown(types.StringType)
	case v.Hosts.IsNull():
		hostsVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		hostsVal, d = types.ListValue(types.StringType, v.Hosts.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"auth_token": basetypes.StringType{},
			"hosts": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"auth_token": basetypes.StringType{},
		"hosts": basetypes.ListType{
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
			"auth_token": v.AuthToken,
			"hosts":      hostsVal,
		})

	return objVal, diags
}

func (v TailscaleValue) Equal(o attr.Value) bool {
	other, ok := o.(TailscaleValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.AuthToken.Equal(other.AuthToken) {
		return false
	}

	if !v.Hosts.Equal(other.Hosts) {
		return false
	}

	return true
}

func (v TailscaleValue) Type(ctx context.Context) attr.Type {
	return TailscaleType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v TailscaleValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"auth_token": basetypes.StringType{},
		"hosts": basetypes.ListType{
			ElemType: types.StringType,
		},
	}
}
