// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_portal_version"
	"github.com/aws/smithy-go/ptr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var portalVersionConverter = generated.PortalVersionConverterImpl{}

func NewPortalVersionDataSource() datasource.DataSource {
	return &portalVersionDatasource{}
}

type portalVersionDatasource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *portalVersionDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_version"
}

func (r *portalVersionDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_portal_version.PortalVersionDataSourceSchema(ctx)
}

func (r *portalVersionDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *portalVersionDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	data := datasource_portal_version.PortalVersionModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	params := flightdeckv1.GetPortalVersionsParams{
		VersionName: ptr.String(data.Version.ValueString()),
	}
	res, err := r.client.GetPortalVersionsWithResponse(ctx, &params)
	if err != nil {
		resp.Diagnostics.AddError("unable to read portal version", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("unable to read portal version", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}

	if len(*res.JSON200.Items) == 0 {
		resp.Diagnostics.AddError("unable to read portal version", http.StatusText(http.StatusNotFound))
		return
	}

	state := portalVersionConverter.ConvertToTfPortalVersion((*res.JSON200.Items)[0])
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
