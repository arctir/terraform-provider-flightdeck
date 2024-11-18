// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_auth_provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var authProviderConverter = generated.AuthProviderConverterImpl{}

func NewAuthProviderResource() resource.Resource {
	return &authProviderResource{}
}

type authProviderResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *authProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_provider"
}

func (r *authProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_auth_provider.AuthProviderResourceSchema(ctx)
}

func (r *authProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *authProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfAuthProvider resource_auth_provider.AuthProviderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfAuthProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	authProvider, err := authProviderConverter.ConvertToApiAuthProvider(tfAuthProvider)
	if err != nil {
		resp.Diagnostics.AddError("could not convert auth provider", err.Error())
		return
	}

	res, err := r.client.CreateAuthProviderWithResponse(ctx, tfAuthProvider.OrganizationId.ValueString(), tfAuthProvider.PortalName.ValueString(), authProvider)
	if err != nil {
		resp.Diagnostics.AddError("auth provider creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("auth provider creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := authProviderConverter.ConvertToTfAuthProvider(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert auth provider", err.Error())
		return
	}

	state.Github.ClientSecret = tfAuthProvider.Github.ClientSecret
	state.Gitlab.ClientSecret = tfAuthProvider.Gitlab.ClientSecret
	state.Google.ClientSecret = tfAuthProvider.Google.ClientSecret

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *authProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfAuthProvider resource_auth_provider.AuthProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfAuthProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetAuthProviderWithResponse(ctx, tfAuthProvider.OrganizationId.ValueString(), tfAuthProvider.PortalName.ValueString(), tfAuthProvider.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("auth provider read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("auth provider read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := authProviderConverter.ConvertToTfAuthProvider(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert auth provider", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *authProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_auth_provider.AuthProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfAuthProvider resource_auth_provider.AuthProviderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfAuthProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	authProvider, err := authProviderConverter.ConvertToApiAuthProvider(tfAuthProvider)
	if err != nil {
		resp.Diagnostics.AddError("could not convert auth provider", err.Error())
		return
	}
	res, err := r.client.UpdateAuthProviderWithResponse(ctx, tfAuthProvider.OrganizationId.ValueString(), tfAuthProvider.PortalName.ValueString(), state.Name.ValueString(), authProvider)
	if err != nil {
		resp.Diagnostics.AddError("auth provider update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("auth provider update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = authProviderConverter.ConvertToTfAuthProvider(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert auth provider", err.Error())
		return
	}

	state.Github.ClientSecret = tfAuthProvider.Github.ClientSecret
	state.Gitlab.ClientSecret = tfAuthProvider.Gitlab.ClientSecret
	state.Google.ClientSecret = tfAuthProvider.Google.ClientSecret

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *authProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfAuthProvider resource_auth_provider.AuthProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfAuthProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteAuthProviderWithResponse(ctx, tfAuthProvider.OrganizationId.ValueString(), tfAuthProvider.PortalName.ValueString(), tfAuthProvider.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("auth provider delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("auth provider delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
