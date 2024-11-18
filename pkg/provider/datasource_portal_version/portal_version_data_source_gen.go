// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_portal_version

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PortalVersionDataSourceSchema(ctx context.Context) schema.Schema {
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
			"major": schema.Int64Attribute{
				Computed:            true,
				Description:         "The version's major release number.",
				MarkdownDescription: "The version's major release number.",
			},
			"minor": schema.Int64Attribute{
				Computed:            true,
				Description:         "The version's minor release number.",
				MarkdownDescription: "The version's minor release number.",
			},
			"patch": schema.Int64Attribute{
				Computed:            true,
				Description:         "The version's patch release number.",
				MarkdownDescription: "The version's patch release number.",
			},
			"rev": schema.Int64Attribute{
				Computed:            true,
				Description:         "The version's revision release number.",
				MarkdownDescription: "The version's revision release number.",
			},
			"version": schema.StringAttribute{
				Required:            true,
				Description:         "The human-readable semver version of the Flightdeck Portal.",
				MarkdownDescription: "The human-readable semver version of the Flightdeck Portal.",
			},
		},
		Description: "Represents a Flightdeck Portal Version resource.",
	}
}

type PortalVersionModel struct {
	CreatedAt types.String `tfsdk:"created_at"`
	Id        types.String `tfsdk:"id"`
	Major     types.Int64  `tfsdk:"major"`
	Minor     types.Int64  `tfsdk:"minor"`
	Patch     types.Int64  `tfsdk:"patch"`
	Rev       types.Int64  `tfsdk:"rev"`
	Version   types.String `tfsdk:"version"`
}