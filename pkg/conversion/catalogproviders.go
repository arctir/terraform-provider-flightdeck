// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	catalogproviderresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_catalog_provider"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"k8s.io/utils/ptr"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type CatalogProviderConverter interface {
	// goverter:map . ProviderConfig | SetCatalogProviderConfig
	ConvertToApiCatalogProvider(source catalogproviderresource.CatalogProviderModel) (*flightdeckv1.CatalogProviderInput, error)
	// goverter:ignoreMissing
	// goverter:default SetTfCatalogProviderConfig
	// goverter:ignore Github Gitlab
	ConvertToTfCatalogProvider(source flightdeckv1.CatalogProvider) (catalogproviderresource.CatalogProviderModel, error)
}

// FIXME: schedules do not work with the tf generator, so hardcode for now
var defaultSchedule = flightdeckv1.TaskScheduleDefinitionConfig{
	Frequency: &flightdeckv1.TaskScheduleDefinitionTimeConfigFrequency{
		Minutes: ptr.To(1),
	},
	Timeout: &flightdeckv1.TaskScheduleDefinitionTimeConfigTimeout{
		Minutes: ptr.To(1),
	},
}

func SetCatalogProviderConfig(source catalogproviderresource.CatalogProviderModel) (flightdeckv1.CatalogProviderInput_ProviderConfig, error) {
	config := flightdeckv1.CatalogProviderInput_ProviderConfig{}

	if !source.Github.IsUnknown() && !source.Gitlab.IsUnknown() {
		return config, nil
	}
	if !source.Github.IsUnknown() {
		var gh flightdeckv1.GithubCatalogProvider
		o, diag := source.Github.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github catalogprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gh, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github catalogprovider %+v", diag.Errors())
		}
		gh.ConfigType = "github"
		gh.Schedule = &defaultSchedule
		config.FromGithubCatalogProvider(gh)
	}
	if !source.Gitlab.IsUnknown() {
		var gl flightdeckv1.GitlabCatalogProvider
		o, diag := source.Gitlab.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab catalogprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gl, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab catalogprovider %+v", diag.Errors())
		}
		gl.ConfigType = "gitlab"
		gl.Schedule = &defaultSchedule
		config.FromGitlabCatalogProvider(gl)
	}
	if !source.Location.IsUnknown() {
		var l flightdeckv1.LocationCatalogProvider
		o, diag := source.Location.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse location catalogprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &l, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse location catalogprovider %+v", diag.Errors())
		}
		config.FromLocationCatalogProvider(l)
	}
	return config, nil
}

func SetTfCatalogProviderConfig(c CatalogProviderConverter, input flightdeckv1.CatalogProvider) (catalogproviderresource.CatalogProviderModel, error) {
	catalogprovider := catalogproviderresource.CatalogProviderModel{
		OrganizationId: basetypes.NewStringValue(input.OrganizationId.String()),
		PortalName:     basetypes.NewStringValue(input.PortalName),
		Github:         catalogproviderresource.NewGithubValueNull(),
		Gitlab:         catalogproviderresource.NewGitlabValueNull(),
	}

	githubTypes := catalogproviderresource.GithubValue{}.AttributeTypes(context.TODO())
	gitlabTypes := catalogproviderresource.GitlabValue{}.AttributeTypes(context.TODO())
	locationTypes := catalogproviderresource.LocationValue{}.AttributeTypes(context.TODO())

	configType, err := input.ProviderConfig.Discriminator()
	if err != nil {
		return catalogprovider, err
	}
	switch configType {
	case "github":
		gh, err := input.ProviderConfig.AsGithubCatalogProvider()
		if err != nil {
			return catalogprovider, err
		}
		github, diag := basetypes.NewObjectValueFrom(context.TODO(), githubTypes, gh)
		if diag.HasError() {
			return catalogprovider, fmt.Errorf("could not parse github catalogprovider %+v", diag.Errors())
		}
		ghv, diag := catalogproviderresource.NewGithubValue(githubTypes, github.Attributes())
		catalogprovider.Github = ghv
	case "gitlab":
		gl, err := input.ProviderConfig.AsGitlabCatalogProvider()
		if err != nil {
			return catalogprovider, err
		}
		gitlab, diag := basetypes.NewObjectValueFrom(context.TODO(), gitlabTypes, gl)
		if diag.HasError() {
			return catalogprovider, fmt.Errorf("could not parse gitlab catalogprovider %+v", diag.Errors())
		}
		glv, diag := catalogproviderresource.NewGithubValue(gitlabTypes, gitlab.Attributes())
		catalogprovider.Github = glv
	case "location":
		l, err := input.ProviderConfig.AsLocationCatalogProvider()
		if err != nil {
			return catalogprovider, err
		}
		location, diag := basetypes.NewObjectValueFrom(context.TODO(), locationTypes, l)
		if diag.HasError() {
			return catalogprovider, fmt.Errorf("could not parse location catalogprovider %+v", diag.Errors())
		}
		lv, diag := catalogproviderresource.NewLocationValue(locationTypes, location.Attributes())
		catalogprovider.Location = lv
	}

	return catalogprovider, nil
}
