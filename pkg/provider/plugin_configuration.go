// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_plugin_configuration"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var pluginConfigurationConverter = generated.PluginConfigurationConverterImpl{}

func NewPluginConfigurationResource() resource.Resource {
	return &pluginConfigurationResource{}
}

type pluginConfigurationResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *pluginConfigurationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_configuration"
}

func (r *pluginConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_plugin_configuration.PluginConfigurationResourceSchema(ctx)
}

func (r *pluginConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *pluginConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfPluginConfiguration resource_plugin_configuration.PluginConfigurationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfPluginConfiguration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	pluginConfiguration, err := pluginConfigurationConverter.ConvertToApiPluginConfiguration(tfPluginConfiguration)
	if err != nil {
		resp.Diagnostics.AddError("could not convert plugin configuration", err.Error())
		return
	}

	res, err := r.client.CreatePluginConfigurationWithResponse(ctx, tfPluginConfiguration.OrganizationId.ValueString(), tfPluginConfiguration.PortalName.ValueString(), pluginConfiguration)
	if err != nil {
		resp.Diagnostics.AddError("plugin configuration creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("plugin configuration creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := pluginConfigurationConverter.ConvertToTfPluginConfiguration(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert plugin configuration", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *pluginConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfPluginConfiguration resource_plugin_configuration.PluginConfigurationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfPluginConfiguration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetPluginConfigurationWithResponse(ctx, tfPluginConfiguration.OrganizationId.ValueString(), tfPluginConfiguration.PortalName.ValueString(), tfPluginConfiguration.Definition.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("plugin configuration read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("plugin configuration read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := pluginConfigurationConverter.ConvertToTfPluginConfiguration(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert plugin configuration", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *pluginConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_plugin_configuration.PluginConfigurationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfPluginConfiguration resource_plugin_configuration.PluginConfigurationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfPluginConfiguration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	pluginConfiguration, err := pluginConfigurationConverter.ConvertToApiPluginConfiguration(tfPluginConfiguration)
	if err != nil {
		resp.Diagnostics.AddError("could not convert plugin configuration", err.Error())
		return
	}
	res, err := r.client.UpdatePluginConfigurationWithResponse(ctx, tfPluginConfiguration.OrganizationId.ValueString(), tfPluginConfiguration.PortalName.ValueString(), state.Definition.Name.ValueString(), pluginConfiguration)
	if err != nil {
		resp.Diagnostics.AddError("plugin configuration update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("plugin configuration update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = pluginConfigurationConverter.ConvertToTfPluginConfiguration(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert plugin configuration", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *pluginConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfPluginConfiguration resource_plugin_configuration.PluginConfigurationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfPluginConfiguration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeletePluginConfigurationWithResponse(ctx, tfPluginConfiguration.OrganizationId.ValueString(), tfPluginConfiguration.PortalName.ValueString(), tfPluginConfiguration.Definition.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("plugin configuration delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("plugin configuration delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
