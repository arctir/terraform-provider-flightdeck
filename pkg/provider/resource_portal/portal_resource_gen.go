// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_portal

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

func PortalResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alternate_domains": schema.ListAttribute{
				ElementType:         types.StringType,
				Required:            true,
				Description:         "A list of alternate domains for the Portal.",
				MarkdownDescription: "A list of alternate domains for the Portal.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				Description:         "The date and time of the resources creation.",
				MarkdownDescription: "The date and time of the resources creation.",
			},
			"domain": schema.StringAttribute{
				Required:            true,
				Description:         "The primary domain of the Portal.",
				MarkdownDescription: "The primary domain of the Portal.",
			},
			"hostname": schema.StringAttribute{
				Computed:            true,
				Description:         "The hostname of the Portal.",
				MarkdownDescription: "The hostname of the Portal.",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of the Flightdeck resource.",
				MarkdownDescription: "The ID of the Flightdeck resource.",
			},
			"identifier": schema.StringAttribute{
				Computed:            true,
				Description:         "The identifier of the Portal.",
				MarkdownDescription: "The identifier of the Portal.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the Portal.",
				MarkdownDescription: "The name of the Portal.",
			},
			"organization_id": schema.StringAttribute{
				Required:            true,
				Description:         "The ID of the Flightdeck Organization resource.",
				MarkdownDescription: "The ID of the Flightdeck Organization resource.",
			},
			"organization_name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the Organization operating this Portal.",
				MarkdownDescription: "The name of the Organization operating this Portal.",
			},
			"status": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"detail": schema.StringAttribute{
						Required:            true,
						Description:         "A detailed message about the current status of the Portal.",
						MarkdownDescription: "A detailed message about the current status of the Portal.",
					},
					"status": schema.StringAttribute{
						Required:            true,
						Description:         "The current status of the Portal.",
						MarkdownDescription: "The current status of the Portal.",
					},
				},
				CustomType: StatusType{
					ObjectType: types.ObjectType{
						AttrTypes: StatusValue{}.AttributeTypes(ctx),
					},
				},
				Optional: true,
				Computed: true,
			},
			"tenant_name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the Tenant providing user idenity to this Portal.",
				MarkdownDescription: "The name of the Tenant providing user idenity to this Portal.",
			},
			"title": schema.StringAttribute{
				Required:            true,
				Description:         "The HTML title of the Portal.",
				MarkdownDescription: "The HTML title of the Portal.",
			},
			"url": schema.StringAttribute{
				Computed:            true,
				Description:         "The primary URL of the Portal.",
				MarkdownDescription: "The primary URL of the Portal.",
			},
			"version_id": schema.StringAttribute{
				Required:            true,
				Description:         "The ID of the Portal Version.",
				MarkdownDescription: "The ID of the Portal Version.",
			},
		},
		Description: "Represents a Portal resource.",
	}
}

type PortalModel struct {
	AlternateDomains types.List   `tfsdk:"alternate_domains"`
	CreatedAt        types.String `tfsdk:"created_at"`
	Domain           types.String `tfsdk:"domain"`
	Hostname         types.String `tfsdk:"hostname"`
	Id               types.String `tfsdk:"id"`
	Identifier       types.String `tfsdk:"identifier"`
	Name             types.String `tfsdk:"name"`
	OrganizationId   types.String `tfsdk:"organization_id"`
	OrganizationName types.String `tfsdk:"organization_name"`
	Status           StatusValue  `tfsdk:"status"`
	TenantName       types.String `tfsdk:"tenant_name"`
	Title            types.String `tfsdk:"title"`
	Url              types.String `tfsdk:"url"`
	VersionId        types.String `tfsdk:"version_id"`
}

var _ basetypes.ObjectTypable = StatusType{}

type StatusType struct {
	basetypes.ObjectType
}

func (t StatusType) Equal(o attr.Type) bool {
	other, ok := o.(StatusType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t StatusType) String() string {
	return "StatusType"
}

func (t StatusType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	detailAttribute, ok := attributes["detail"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`detail is missing from object`)

		return nil, diags
	}

	detailVal, ok := detailAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`detail expected to be basetypes.StringValue, was: %T`, detailAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return nil, diags
	}

	statusVal, ok := statusAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be basetypes.StringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return StatusValue{
		Detail: detailVal,
		Status: statusVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewStatusValueNull() StatusValue {
	return StatusValue{
		state: attr.ValueStateNull,
	}
}

func NewStatusValueUnknown() StatusValue {
	return StatusValue{
		state: attr.ValueStateUnknown,
	}
}

func NewStatusValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (StatusValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing StatusValue Attribute Value",
				"While creating a StatusValue value, a missing attribute value was detected. "+
					"A StatusValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("StatusValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid StatusValue Attribute Type",
				"While creating a StatusValue value, an invalid attribute value was detected. "+
					"A StatusValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("StatusValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("StatusValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra StatusValue Attribute Value",
				"While creating a StatusValue value, an extra attribute value was detected. "+
					"A StatusValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra StatusValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewStatusValueUnknown(), diags
	}

	detailAttribute, ok := attributes["detail"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`detail is missing from object`)

		return NewStatusValueUnknown(), diags
	}

	detailVal, ok := detailAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`detail expected to be basetypes.StringValue, was: %T`, detailAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewStatusValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be basetypes.StringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return NewStatusValueUnknown(), diags
	}

	return StatusValue{
		Detail: detailVal,
		Status: statusVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewStatusValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) StatusValue {
	object, diags := NewStatusValue(attributeTypes, attributes)

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

		panic("NewStatusValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t StatusType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewStatusValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewStatusValueUnknown(), nil
	}

	if in.IsNull() {
		return NewStatusValueNull(), nil
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

	return NewStatusValueMust(StatusValue{}.AttributeTypes(ctx), attributes), nil
}

func (t StatusType) ValueType(ctx context.Context) attr.Value {
	return StatusValue{}
}

var _ basetypes.ObjectValuable = StatusValue{}

type StatusValue struct {
	Detail basetypes.StringValue `tfsdk:"detail"`
	Status basetypes.StringValue `tfsdk:"status"`
	state  attr.ValueState
}

func (v StatusValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["detail"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Detail.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["detail"] = val

		val, err = v.Status.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["status"] = val

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

func (v StatusValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v StatusValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v StatusValue) String() string {
	return "StatusValue"
}

func (v StatusValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"detail": basetypes.StringType{},
		"status": basetypes.StringType{},
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
			"detail": v.Detail,
			"status": v.Status,
		})

	return objVal, diags
}

func (v StatusValue) Equal(o attr.Value) bool {
	other, ok := o.(StatusValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Detail.Equal(other.Detail) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	return true
}

func (v StatusValue) Type(ctx context.Context) attr.Type {
	return StatusType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v StatusValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"detail": basetypes.StringType{},
		"status": basetypes.StringType{},
	}
}
