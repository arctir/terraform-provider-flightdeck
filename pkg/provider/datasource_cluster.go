// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/datasource_cluster"
	"github.com/aws/smithy-go/ptr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var clusterConverter = generated.ClusterConverterImpl{}

func NewClusterDataSource() datasource.DataSource {
	return &clusterDatasource{}
}

type clusterDatasource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *clusterDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *clusterDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_cluster.ClusterDataSourceSchema(ctx)
}

func (r *clusterDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *clusterDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	data := datasource_cluster.ClusterModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	params := flightdeckv1.GetClustersParams{
		Name: ptr.String(data.Name.ValueString()),
	}
	res, err := r.client.GetClustersWithResponse(ctx, &params)
	if err != nil {
		resp.Diagnostics.AddError("uanble to read cluster", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("uanble to read cluster", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}

	if len(*res.JSON200.Items) == 0 {
		resp.Diagnostics.AddError("unable to read cluster", http.StatusText(http.StatusNotFound))
		return
	}

	state := clusterConverter.ConvertToTfCluster((*res.JSON200.Items)[0])
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
