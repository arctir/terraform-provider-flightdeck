// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_catalog_provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var catalogProviderConverter = generated.CatalogProviderConverterImpl{}

func NewCatalogProviderResource() resource.Resource {
	return &catalogProviderResource{}
}

type catalogProviderResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *catalogProviderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_provider"
}

func (r *catalogProviderResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_catalog_provider.CatalogProviderResourceSchema(ctx)
}

func (r *catalogProviderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *catalogProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfCatalogProvider resource_catalog_provider.CatalogProviderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfCatalogProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	catalogProvider, err := catalogProviderConverter.ConvertToApiCatalogProvider(tfCatalogProvider)
	if err != nil {
		resp.Diagnostics.AddError("could not convert catalog provider", err.Error())
		return
	}

	res, err := r.client.CreateCatalogProviderWithResponse(ctx, tfCatalogProvider.OrganizationId.ValueString(), tfCatalogProvider.PortalName.ValueString(), *catalogProvider)
	if err != nil {
		resp.Diagnostics.AddError("catalog provider creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("catalog provider creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := catalogProviderConverter.ConvertToTfCatalogProvider(*res.JSON201)
	if err != nil {
		resp.Diagnostics.AddError("catalog provider creation failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *catalogProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfCatalogProvider resource_catalog_provider.CatalogProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfCatalogProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetCatalogProviderWithResponse(ctx, tfCatalogProvider.OrganizationId.ValueString(), tfCatalogProvider.PortalName.ValueString(), tfCatalogProvider.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("catalog provider read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("catalog provider read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err := catalogProviderConverter.ConvertToTfCatalogProvider(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("catalog provider read failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *catalogProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_catalog_provider.CatalogProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var tfCatalogProvider resource_catalog_provider.CatalogProviderModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfCatalogProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	catalogProvider, err := catalogProviderConverter.ConvertToApiCatalogProvider(tfCatalogProvider)
	if err != nil {
		resp.Diagnostics.AddError("could not convert catalog provider", err.Error())
		return
	}
	res, err := r.client.UpdateCatalogProviderWithResponse(ctx, tfCatalogProvider.OrganizationId.ValueString(), tfCatalogProvider.PortalName.ValueString(), state.Name.ValueString(), *catalogProvider)
	if err != nil {
		resp.Diagnostics.AddError("catalog provider update failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("catalog provider update failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state, err = catalogProviderConverter.ConvertToTfCatalogProvider(*res.JSON200)
	if err != nil {
		resp.Diagnostics.AddError("catalog provider update failed", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *catalogProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfCatalogProvider resource_catalog_provider.CatalogProviderModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfCatalogProvider)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteCatalogProviderWithResponse(ctx, tfCatalogProvider.OrganizationId.ValueString(), tfCatalogProvider.PortalName.ValueString(), tfCatalogProvider.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("catalog provider delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("catalog provider delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
