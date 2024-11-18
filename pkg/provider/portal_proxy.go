// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_portal_proxy"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var portalProxyConverter = generated.PortalProxyConverterImpl{}

func NewPortalProxyResource() resource.Resource {
	return &portalProxyProviderResource{}
}

type portalProxyProviderResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *portalProxyProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_proxy"
}

func (r *portalProxyProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_portal_proxy.PortalProxyResourceSchema(ctx)
}

func (r *portalProxyProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *portalProxyProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfPortalProxy resource_portal_proxy.PortalProxyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfPortalProxy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	portalProxyProvider, err := portalProxyConverter.ConvertToApiPortalProxy(tfPortalProxy)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal proxy", err.Error())
		return
	}

	res, err := r.client.CreatePortalProxyWithResponse(ctx, tfPortalProxy.OrganizationId.ValueString(), tfPortalProxy.PortalName.ValueString(), portalProxyProvider)
	if err != nil {
		resp.Diagnostics.AddError("portal proxy creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("portal proxy creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := portalProxyConverter.ConvertToTfPortalProxy(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal proxy", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *portalProxyProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfPortalProxy resource_portal_proxy.PortalProxyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfPortalProxy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetPortalProxyWithResponse(ctx, tfPortalProxy.OrganizationId.ValueString(), tfPortalProxy.PortalName.ValueString(), tfPortalProxy.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("portal proxy read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("portal proxy read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := portalProxyConverter.ConvertToTfPortalProxy(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal proxy", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *portalProxyProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_portal_proxy.PortalProxyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfPortalProxy resource_portal_proxy.PortalProxyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfPortalProxy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	portalProxyProvider, err := portalProxyConverter.ConvertToApiPortalProxy(tfPortalProxy)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal proxy", err.Error())
		return
	}
	res, err := r.client.UpdatePortalProxyWithResponse(ctx, tfPortalProxy.OrganizationId.ValueString(), tfPortalProxy.PortalName.ValueString(), state.Name.ValueString(), portalProxyProvider)
	if err != nil {
		resp.Diagnostics.AddError("portal proxy update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("portal proxy update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = portalProxyConverter.ConvertToTfPortalProxy(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal proxy", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *portalProxyProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfPortalProxy resource_portal_proxy.PortalProxyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfPortalProxy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeletePortalProxyWithResponse(ctx, tfPortalProxy.OrganizationId.ValueString(), tfPortalProxy.PortalName.ValueString(), tfPortalProxy.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("portal proxy delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("portal proxy delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
