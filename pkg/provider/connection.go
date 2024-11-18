// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_connection"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var connectionConverter = generated.ConnectionConverterImpl{}

func NewConnectionResource() resource.Resource {
	return &connectionResource{}
}

type connectionResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *connectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection"
}

func (r *connectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_connection.ConnectionResourceSchema(ctx)
}

func (r *connectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *connectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfConnection resource_connection.ConnectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfConnection)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connection, err := connectionConverter.ConvertToApiConnection(tfConnection)
	if err != nil {
		resp.Diagnostics.AddError("could not convert connection", err.Error())
		return
	}

	res, err := r.client.CreateConnectionWithResponse(ctx, tfConnection.OrganizationId.ValueString(), tfConnection.PortalName.ValueString(), *connection)
	if err != nil {
		resp.Diagnostics.AddError("connection creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("connection creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := connectionConverter.ConvertToTfConnection(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert connection", err.Error())
		return
	}
	state.Tailscale.AuthToken = tfConnection.Tailscale.AuthToken
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *connectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfConnection resource_connection.ConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfConnection)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetConnectionWithResponse(ctx, tfConnection.OrganizationId.ValueString(), tfConnection.PortalName.ValueString(), tfConnection.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("connection read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("connection read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := connectionConverter.ConvertToTfConnection(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert connection", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *connectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_connection.ConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfConnection resource_connection.ConnectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfConnection)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connection, err := connectionConverter.ConvertToApiConnection(tfConnection)
	if err != nil {
		resp.Diagnostics.AddError("could not convert connection", err.Error())
		return
	}
	res, err := r.client.UpdateConnectionWithResponse(ctx, tfConnection.OrganizationId.ValueString(), tfConnection.PortalName.ValueString(), state.Name.ValueString(), *connection)
	if err != nil {
		resp.Diagnostics.AddError("connection update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("connection update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = connectionConverter.ConvertToTfConnection(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert connection", err.Error())
		return
	}
	state.Tailscale.AuthToken = tfConnection.Tailscale.AuthToken
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *connectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfConnection resource_connection.ConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfConnection)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteConnectionWithResponse(ctx, tfConnection.OrganizationId.ValueString(), tfConnection.PortalName.ValueString(), tfConnection.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("connection delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("connection delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
