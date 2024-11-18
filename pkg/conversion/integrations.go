// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	integrationresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_integration"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type IntegrationConverter interface {
	// goverter:map . IntegrationConfig | SetApiIntegrationConfig
	ConvertToApiIntegration(source integrationresource.IntegrationModel) (*flightdeckv1.IntegrationInput, error)
	// goverter:ignoreMissing
	// goverter:default SetTfIntegrationConfig
	ConvertToTfIntegration(source flightdeckv1.Integration) (integrationresource.IntegrationModel, error)
}

func SetApiIntegrationConfig(c IntegrationConverter, source integrationresource.IntegrationModel) (flightdeckv1.IntegrationInput_IntegrationConfig, error) {
	config := flightdeckv1.IntegrationInput_IntegrationConfig{}

	if !source.Github.IsUnknown() && !source.Gitlab.IsUnknown() {
		return config, nil
	}
	if !source.Github.IsUnknown() {
		var gh flightdeckv1.GithubIntegration
		o, diag := source.Github.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github integration %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gh, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github integration %+v", diag.Errors())
		}
		gh.ConfigType = "github"
		config.FromGithubIntegration(gh)
	}
	if !source.Gitlab.IsUnknown() {
		var gl flightdeckv1.GitlabIntegration
		o, diag := source.Gitlab.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab integration %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gl, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab integration %+v", diag.Errors())
		}
		gl.ConfigType = "gitlab"
		config.FromGitlabIntegration(gl)
	}
	return config, nil
}

func SetTfIntegrationConfig(c IntegrationConverter, input flightdeckv1.Integration) (integrationresource.IntegrationModel, error) {
	integration := integrationresource.IntegrationModel{
		OrganizationId: basetypes.NewStringValue(input.OrganizationId.String()),
		PortalName:     basetypes.NewStringValue(input.PortalName),
		Github:         integrationresource.NewGithubValueNull(),
		Gitlab:         integrationresource.NewGitlabValueNull(),
	}

	githubTypes := integrationresource.GithubValue{}.AttributeTypes(context.TODO())
	gitlabTypes := integrationresource.GitlabValue{}.AttributeTypes(context.TODO())

	configType, err := input.IntegrationConfig.Discriminator()
	if err != nil {
		return integration, err
	}
	switch configType {
	case "github":
		gh, err := input.IntegrationConfig.AsGithubIntegration()
		if err != nil {
			return integration, err
		}
		github, diag := basetypes.NewObjectValueFrom(context.TODO(), githubTypes, gh)
		if diag.HasError() {
			return integration, fmt.Errorf("could not parse github integration %+v", diag.Errors())
		}
		ghv, diag := integrationresource.NewGithubValue(githubTypes, github.Attributes())
		if diag.HasError() {
			return integration, fmt.Errorf("could not parse gitlab integration %+v", diag.Errors())
		}
		integration.Github = ghv
	case "gitlab":
		gl, err := input.IntegrationConfig.AsGitlabIntegration()
		if err != nil {
			return integration, err
		}
		gitlab, diag := basetypes.NewObjectValueFrom(context.TODO(), gitlabTypes, gl)
		if diag.HasError() {
			return integration, fmt.Errorf("could not parse gitlab integration %+v", diag.Errors())
		}
		glv, diag := integrationresource.NewGitlabValue(gitlabTypes, gitlab.Attributes())
		if diag.HasError() {
			return integration, fmt.Errorf("could not parse gitlab integration %+v", diag.Errors())
		}
		integration.Gitlab = glv
	}
	return integration, nil
}
