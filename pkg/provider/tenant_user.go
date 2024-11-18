// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_tenant_user"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var tenantUserConverter = generated.TenantUserConverterImpl{}

func NewTenantUserResource() resource.Resource {
	return &tenantUserResource{}
}

type tenantUserResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *tenantUserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_user"
}

func (r *tenantUserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_tenant_user.TenantUserResourceSchema(ctx)
}

func (r *tenantUserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *tenantUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfTenantUser resource_tenant_user.TenantUserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfTenantUser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user := tenantUserConverter.ConvertToApiTenantUser(tfTenantUser)
	res, err := r.client.CreateTenantUserWithResponse(ctx, tfTenantUser.OrganizationId.ValueString(), tfTenantUser.TenantName.ValueString(), user)
	if err != nil {
		resp.Diagnostics.AddError("tenant user creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("tenant user creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state := tenantUserConverter.ConvertToTfTenantUser(*res.JSON201)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tenantUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfTenantUser resource_tenant_user.TenantUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfTenantUser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetTenantUserWithResponse(ctx, tfTenantUser.OrganizationId.ValueString(), tfTenantUser.TenantName.ValueString(), tfTenantUser.Username.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("tenant user read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("tenant user read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state := tenantUserConverter.ConvertToTfTenantUser(*res.JSON200)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *tenantUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	/*
	   var state resource_tenant_user.TenantUserModel
	   resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	   	if resp.Diagnostics.HasError() {
	   		return
	   	}

	   var tfTenantUser resource_tenant_user.TenantUserModel
	   resp.Diagnostics.Append(req.Plan.Get(ctx, &tfTenantUser)...)

	   	if resp.Diagnostics.HasError() {
	   		return
	   	}

	   converter := generated.ConverterImpl{}
	   tenant user := tenantUserConverter.ConvertToApiTenantUser(tfTenantUser)
	   res, err := r.client.UpdateTenantUserByNameWithResponse(ctx, tfTenantUser.OrgId.ValueString(), tfTenantUser.PortalName.ValueString(), state.Name.ValueString(), tenant user)

	   	if err != nil {
	   		resp.Diagnostics.AddError("tenant user update failed", err.Error())
	   		return
	   	}

	   	if res.HTTPResponse.StatusCode != http.StatusOK {
	   		resp.Diagnostics.AddError("tenant user update failed", http.StatusText(res.HTTPResponse.StatusCode))
	   		return
	   	}

	   state = tenantUserConverter.ConvertToTfTenantUser(*res.JSON200)
	   resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	*/
}

func (r *tenantUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfTenantUser resource_tenant_user.TenantUserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfTenantUser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteTenantUserWithResponse(ctx, tfTenantUser.OrganizationId.ValueString(), tfTenantUser.TenantName.ValueString(), tfTenantUser.Username.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("tenant user delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("tenant user delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
