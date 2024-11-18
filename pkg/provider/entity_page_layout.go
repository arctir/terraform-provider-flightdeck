// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_entity_page_layout"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var entityPageLayoutConverter = generated.EntityPageLayoutConverterImpl{}

func NewEntityPageLayoutResource() resource.Resource {
	return &entityPageLayoutResource{}
}

type entityPageLayoutResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *entityPageLayoutResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entity_page_layout"
}

func (r *entityPageLayoutResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_entity_page_layout.EntityPageLayoutResourceSchema(ctx)
}

func (r *entityPageLayoutResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *entityPageLayoutResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfEntityPageLayout resource_entity_page_layout.EntityPageLayoutModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfEntityPageLayout)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entityPageLayout, err := entityPageLayoutConverter.ConvertToApiEntityPageLayout(tfEntityPageLayout)
	if err != nil {
		resp.Diagnostics.AddError("could not convert entity page layout", err.Error())
		return
	}

	res, err := r.client.CreateEntityPageLayoutWithResponse(ctx, tfEntityPageLayout.OrganizationId.ValueString(), tfEntityPageLayout.PortalName.ValueString(), entityPageLayout)
	if err != nil {
		resp.Diagnostics.AddError("entity page layout creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("entity page layout creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := entityPageLayoutConverter.ConvertToTfEntityPageLayout(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("could not convert entity page layout", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *entityPageLayoutResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfEntityPageLayout resource_entity_page_layout.EntityPageLayoutModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfEntityPageLayout)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetEntityPageLayoutWithResponse(ctx, tfEntityPageLayout.OrganizationId.ValueString(), tfEntityPageLayout.PortalName.ValueString(), tfEntityPageLayout.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("entity page layout read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("entity page layout read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := entityPageLayoutConverter.ConvertToTfEntityPageLayout(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert entity page layout", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *entityPageLayoutResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_entity_page_layout.EntityPageLayoutModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfEntityPageLayout resource_entity_page_layout.EntityPageLayoutModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfEntityPageLayout)...)
	if resp.Diagnostics.HasError() {
		return
	}

	entityPageLayout, err := entityPageLayoutConverter.ConvertToApiEntityPageLayout(tfEntityPageLayout)
	if err != nil {
		resp.Diagnostics.AddError("could not convert entity page layout", err.Error())
		return
	}
	res, err := r.client.UpdateEntityPageLayoutWithResponse(ctx, tfEntityPageLayout.OrganizationId.ValueString(), tfEntityPageLayout.PortalName.ValueString(), state.Name.ValueString(), entityPageLayout)
	if err != nil {
		resp.Diagnostics.AddError("entity page layout update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("entity page layout update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = entityPageLayoutConverter.ConvertToTfEntityPageLayout(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("could not convert entity page layout", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *entityPageLayoutResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfEntityPageLayout resource_entity_page_layout.EntityPageLayoutModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfEntityPageLayout)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteEntityPageLayoutWithResponse(ctx, tfEntityPageLayout.OrganizationId.ValueString(), tfEntityPageLayout.PortalName.ValueString(), tfEntityPageLayout.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("entity page layout delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("entity page layout delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
