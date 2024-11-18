// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_integration"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var integrationConverter = generated.IntegrationConverterImpl{}

func NewIntegrationResource() resource.Resource {
	return &integrationResource{}
}

type integrationResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *integrationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_integration"
}

func (r *integrationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_integration.IntegrationResourceSchema(ctx)
}

func (r *integrationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *integrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfIntegration resource_integration.IntegrationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfIntegration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	integration, err := integrationConverter.ConvertToApiIntegration(tfIntegration)
	if err != nil {
		resp.Diagnostics.AddError("could not convert integration", err.Error())
		return
	}

	res, err := r.client.CreateIntegrationWithResponse(ctx, tfIntegration.OrganizationId.ValueString(), tfIntegration.PortalName.ValueString(), *integration)
	if err != nil {
		resp.Diagnostics.AddError("integration creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("integration creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := integrationConverter.ConvertToTfIntegration(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert integration", err.Error())
		return
	}
	state.Github.Token = tfIntegration.Github.Token
	state.Gitlab.Token = tfIntegration.Gitlab.Token
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *integrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfIntegration resource_integration.IntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfIntegration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetIntegrationWithResponse(ctx, tfIntegration.OrganizationId.ValueString(), tfIntegration.PortalName.ValueString(), tfIntegration.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("integration read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}

	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("integration read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := integrationConverter.ConvertToTfIntegration(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert integration", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *integrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_integration.IntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfIntegration resource_integration.IntegrationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfIntegration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	integration, err := integrationConverter.ConvertToApiIntegration(tfIntegration)
	if err != nil {
		resp.Diagnostics.AddError("could not convert integration", err.Error())
		return
	}

	res, err := r.client.UpdateIntegrationWithResponse(ctx, tfIntegration.OrganizationId.ValueString(), tfIntegration.PortalName.ValueString(), state.Name.ValueString(), *integration)
	if err != nil {
		resp.Diagnostics.AddError("integration update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("integration update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = integrationConverter.ConvertToTfIntegration(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert integration", err.Error())
		return
	}
	state.Github.Token = tfIntegration.Github.Token
	state.Gitlab.Token = tfIntegration.Gitlab.Token
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *integrationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfIntegration resource_integration.IntegrationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfIntegration)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteIntegrationWithResponse(ctx, tfIntegration.OrganizationId.ValueString(), tfIntegration.PortalName.ValueString(), tfIntegration.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("integration delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("integration delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
