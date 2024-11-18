// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_organization"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var organizationConverter = generated.OrganizationConverterImpl{}

func NewOrganizationDataSource() datasource.DataSource {
	return &organizationDatasource{}
}

type organizationDatasource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *organizationDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (r *organizationDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_organization.OrganizationDataSourceSchema(ctx)
}

func (r *organizationDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *organizationDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	data := datasource_organization.OrganizationModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	res, err := r.client.GetOrganizationByIDWithResponse(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("uanble to read organization", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("uanble to read organization", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}

	state := organizationConverter.ConvertToTfOrganization((*res.JSON200))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
