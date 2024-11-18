// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_portal"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var portalConverter = generated.PortalConverterImpl{}

func NewPortalResource() resource.Resource {
	return &portalResource{}
}

type portalResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *portalResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal"
}

func (r *portalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_portal.PortalResourceSchema(ctx)
}

func (r *portalResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *portalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfPortal resource_portal.PortalModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfPortal)...)
	if resp.Diagnostics.HasError() {
		return
	}

	portal, err := portalConverter.ConvertToApiPortal(tfPortal)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal", err.Error())
		return
	}
	res, err := r.client.CreatePortalWithResponse(ctx, tfPortal.OrganizationId.ValueString(), portal)
	if err != nil {
		resp.Diagnostics.AddError("portal creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("portal creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := portalConverter.ConvertToTfPortal(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *portalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfPortal resource_portal.PortalModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfPortal)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetPortalWithResponse(ctx, tfPortal.OrganizationId.ValueString(), tfPortal.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("portal read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("portal read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := portalConverter.ConvertToTfPortal(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *portalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_portal.PortalModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfPortal resource_portal.PortalModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfPortal)...)
	if resp.Diagnostics.HasError() {
		return
	}

	portal, err := portalConverter.ConvertToApiPortal(tfPortal)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal", err.Error())
		return
	}
	res, err := r.client.UpdatePortalWithResponse(ctx, state.OrganizationId.ValueString(), state.Name.ValueString(), portal)
	if err != nil {
		resp.Diagnostics.AddError("portal update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("portal update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = portalConverter.ConvertToTfPortal(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert portal", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *portalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfPortal resource_portal.PortalModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfPortal)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeletePortalWithResponse(ctx, tfPortal.OrganizationId.ValueString(), tfPortal.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("portal delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("portal delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
