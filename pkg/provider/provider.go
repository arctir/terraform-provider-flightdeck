// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"os"
	"path"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	flightdeckclient "github.com/arctir/go-flightdeck/pkg/client"
)

const defaultApiEndpoint = "https://api.arctir.cloud/v1"

var _ provider.Provider = (*flightdeckProvider)(nil)

func New(apiEndpoint, configPath types.String) func() provider.Provider {
	return func() provider.Provider {
		return &flightdeckProvider{
			ApiEndpoint: apiEndpoint,
			ConfigPath:  configPath,
		}
	}
}

type flightdeckProvider struct {
	ApiEndpoint types.String `tfsdk:"api_endpoint"`
	ConfigPath  types.String `tfsdk:"config_path"`
}

func (p *flightdeckProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "The endpoint for the flightdeck API",
			},
			"config_path": schema.StringAttribute{
				Optional:    true,
				Description: "The path to the flightdeck config file",
			},
		},
	}
}

func (p *flightdeckProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config flightdeckProvider
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var configPath string
	if config.ConfigPath != basetypes.NewStringNull() {
		configPath = config.ConfigPath.String()
	} else {
		if p, ok := os.LookupEnv("FLIGHTDECK_CONFIG"); ok {
			configPath = p
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				resp.Diagnostics.AddError("unable to get home directory", err.Error())
			} else {
				configPath = path.Join(home, "/.flightdeck/config.yaml")
			}
		}
	}

	apiConfig, err := flightdeckclient.ReadConfig(configPath)
	if err != nil {
		resp.Diagnostics.AddError("Unable to create arctir api client", err.Error())
		return
	}

	apiEndpoint := config.ApiEndpoint.ValueString()
	if apiEndpoint == "" {
		apiEndpoint = defaultApiEndpoint
	}

	client, err := flightdeckclient.NewClient(apiEndpoint, *apiConfig)
	if err != nil {
		resp.Diagnostics.AddError("Unable to create arctir api client", err.Error())
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *flightdeckProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "flightdeck"
}

func (p *flightdeckProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewClusterDataSource,
		NewOrganizationDataSource,
		NewPortalVersionDataSource,
	}
}

func (p *flightdeckProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTenantResource,
		NewTenantUserResource,
		NewPortalResource,
		NewIntegrationResource,
		NewCatalogProviderResource,
		NewConnectionResource,
		NewAuthProviderResource,
		NewPortalProxyResource,
		NewEntityPageLayoutResource,
		NewPluginConfigurationResource,
	}
}
