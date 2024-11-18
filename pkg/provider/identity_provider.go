// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_identity_provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var identityProviderConverter = generated.IdentityProviderConverterImpl{}

func NewIdentityProviderResource() resource.Resource {
	return &identityProviderResource{}
}

type identityProviderResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *identityProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity_provider"
}

func (r *identityProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_identity_provider.IdentityProviderResourceSchema(ctx)
}

func (r *identityProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *identityProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfIdentityProvider resource_identity_provider.IdentityProviderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfIdentityProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	identityProvider, err := identityProviderConverter.ConvertToApiIdentityProvider(tfIdentityProvider)
	if err != nil {
		resp.Diagnostics.AddError("could not convert identity provider", err.Error())
		return
	}

	res, err := r.client.CreateIdentityProviderWithResponse(ctx, tfIdentityProvider.OrganizationId.ValueString(), tfIdentityProvider.TenantName.ValueString(), identityProvider)
	if err != nil {
		resp.Diagnostics.AddError("identity provider creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("identity provider creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := identityProviderConverter.ConvertToTfIdentityProvider(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert identity provider", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *identityProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfIdentityProvider resource_identity_provider.IdentityProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfIdentityProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetIdentityProviderWithResponse(ctx, tfIdentityProvider.OrganizationId.ValueString(), tfIdentityProvider.TenantName.ValueString(), tfIdentityProvider.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("identity provider read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("identity provider read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := identityProviderConverter.ConvertToTfIdentityProvider(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert identity provider", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *identityProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_identity_provider.IdentityProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfIdentityProvider resource_identity_provider.IdentityProviderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfIdentityProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	identityProvider, err := identityProviderConverter.ConvertToApiIdentityProvider(tfIdentityProvider)
	if err != nil {
		resp.Diagnostics.AddError("could not convert identity provider", err.Error())
		return
	}
	res, err := r.client.UpdateIdentityProviderWithResponse(ctx, tfIdentityProvider.OrganizationId.ValueString(), tfIdentityProvider.TenantName.ValueString(), state.Name.ValueString(), identityProvider)
	if err != nil {
		resp.Diagnostics.AddError("identity provider update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("identity provider update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = identityProviderConverter.ConvertToTfIdentityProvider(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert identity provider", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *identityProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfIdentityProvider resource_identity_provider.IdentityProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfIdentityProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteIdentityProviderWithResponse(ctx, tfIdentityProvider.OrganizationId.ValueString(), tfIdentityProvider.TenantName.ValueString(), tfIdentityProvider.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("identity provider delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("identity provider delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
