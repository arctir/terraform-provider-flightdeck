// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

/*
import (
	"context"
	"net/http"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/arctir/terraform-provider-flightdeck/pkg/conversion/generated"
	"github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_catalog_location"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewCatalogLocationResource() resource.Resource {
	return &catalogLocationResource{}
}

type catalogLocationResource struct {
	client *flightdeckv1.ClientWithResponses
}

func (r *catalogLocationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_location"
}

func (r *catalogLocationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_catalog_location.CatalogLocationResourceSchema(ctx)
}

func (r *catalogLocationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	client := configure(ctx, req.ProviderData, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	r.client = client
}

func (r *catalogLocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var tfCatalogLocation resource_catalog_location.CatalogLocationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &tfCatalogLocation)...)
	if resp.Diagnostics.HasError() {
		return
	}

	converter := generated.ConverterImpl{}
	catalogLocation := converter.ConvertToApiCatalogLocation(tfCatalogLocation)
	res, err := r.client.CreateCatalogLocationWithResponse(ctx, tfCatalogLocation.OrganizationId.ValueString(), tfCatalogLocation.PortalName.ValueString(), catalogLocation)
	if err != nil {
		resp.Diagnostics.AddError("catalog location creation failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusCreated {
		resp.Diagnostics.AddError("catalog location creation failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state := converter.ConvertToTfCatalogLocation(*res.JSON201)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *catalogLocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var tfCatalogLocation resource_catalog_location.CatalogLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfCatalogLocation)...)
	if resp.Diagnostics.HasError() {
		return
	}

	converter := generated.ConverterImpl{}
	res, err := r.client.GetCatalogLocationByNameWithResponse(ctx, tfCatalogLocation.OrganizationId.ValueString(), tfCatalogLocation.PortalName.ValueString(), tfCatalogLocation.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("catalog location read failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("catalog location read failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
	state := converter.ConvertToTfCatalogLocation(*res.JSON200)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *catalogLocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	/*
	   var state resource_catalog_location.CatalogLocationModel
	   resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	   	if resp.Diagnostics.HasError() {
	   		return
	   	}

	   var tfCatalogLocation resource_catalog_location.CatalogLocationModel
	   resp.Diagnostics.Append(req.Plan.Get(ctx, &tfCatalogLocation)...)

	   	if resp.Diagnostics.HasError() {
	   		return
	   	}

	   converter := generated.ConverterImpl{}
	   catalogLocation := converter.ConvertToApiCatalogLocation(tfCatalogLocation)
	   res, err := r.client.UpdateCatalogLocationByNameWithResponse(ctx, tfCatalogLocation.OrgId.ValueString(), tfCatalogLocation.PortalName.ValueString(), state.Name.ValueString(), catalogLocation)

	   	if err != nil {
	   		resp.Diagnostics.AddError("catalog location update failed", err.Error())
	   		return
	   	}

	   	if res.HTTPResponse.StatusCode != http.StatusOK {
	   		resp.Diagnostics.AddError("catalog location update failed", http.StatusText(res.HTTPResponse.StatusCode))
	   		return
	   	}

	   state = converter.ConvertToTfCatalogLocation(*res.JSON200)
	   resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	*
}

func (r *catalogLocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var tfCatalogLocation resource_catalog_location.CatalogLocationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &tfCatalogLocation)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.DeleteCatalogLocationByNameWithResponse(ctx, tfCatalogLocation.OrganizationId.ValueString(), tfCatalogLocation.PortalName.ValueString(), tfCatalogLocation.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("catalog location delete failed", err.Error())
		return
	}
	if res.HTTPResponse.StatusCode != http.StatusNoContent {
		resp.Diagnostics.AddError("catalog location delete failed", http.StatusText(res.HTTPResponse.StatusCode))
		return
	}
}
*/
