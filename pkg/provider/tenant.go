// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_tenant"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var tenantConverter = generated.TenantConverterImpl{}

func NewTenantResource() resource.Resource {
	return &tenantResource{}
}

type tenantResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *tenantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

func (r *tenantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_tenant.TenantResourceSchema(ctx)
}

func (r *tenantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *tenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfTenant resource_tenant.TenantModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfTenant)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tenant := tenantConverter.ConvertToApiTenant(tfTenant)
	res, err := r.client.CreateTenantWithResponse(ctx, tfTenant.OrganizationId.ValueString(), tenant)
	if err != nil {
		resp.Diagnostics.AddError("tenant creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("tenant creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state := tenantConverter.ConvertToTfTenant(*res.JSON201)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfTenant resource_tenant.TenantModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfTenant)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetTenantWithResponse(ctx, tfTenant.OrganizationId.ValueString(), tfTenant.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("tenant read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("tenant read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state := tenantConverter.ConvertToTfTenant(*res.JSON200)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_tenant.TenantModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfTenant resource_tenant.TenantModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfTenant)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tenant := tenantConverter.ConvertToApiTenant(tfTenant)
	res, err := r.client.UpdateTenantWithResponse(ctx, state.OrganizationId.ValueString(), state.Name.ValueString(), tenant)
	if err != nil {
		resp.Diagnostics.AddError("tenant update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("tenant update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state = tenantConverter.ConvertToTfTenant(*res.JSON200)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfTenant resource_tenant.TenantModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfTenant)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteTenantWithResponse(ctx, tfTenant.OrganizationId.ValueString(), tfTenant.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("tenant delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("tenant delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
