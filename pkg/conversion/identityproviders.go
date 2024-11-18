// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	identityproviderresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_identity_provider"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type IdentityProviderConverter interface {
	// goverter:map . ProviderConfig | SetApiIdentityProviderConfig
	ConvertToApiIdentityProvider(source identityproviderresource.IdentityProviderModel) (flightdeckv1.IdentityProviderInput, error)
	// goverter:ignoreMissing
	// goverter:ignore Github Gitlab
	ConvertToTfIdentityProvider(source flightdeckv1.IdentityProvider) (identityproviderresource.IdentityProviderModel, error)
}

func SetApiIdentityProviderConfig(source identityproviderresource.IdentityProviderModel) (flightdeckv1.IdentityProviderInput_ProviderConfig, error) {
	config := flightdeckv1.IdentityProviderInput_ProviderConfig{}

	if !source.Github.IsUnknown() && !source.Gitlab.IsUnknown() {
		return config, nil
	}
	if !source.Github.IsUnknown() {
		var gh flightdeckv1.GithubIdentityProvider
		o, diag := source.Github.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github identityprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gh, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github identityprovider %+v", diag.Errors())
		}
		gh.ConfigType = "github"
		config.FromGithubIdentityProvider(gh)
	}
	if !source.Gitlab.IsUnknown() {
		var gl flightdeckv1.GitlabIdentityProvider
		o, diag := source.Gitlab.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab identityprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gl, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab identityprovider %+v", diag.Errors())
		}
		gl.ConfigType = "gitlab"
		config.FromGitlabIdentityProvider(gl)
	}
	return config, nil
}

func SetTfIdentityProviderConfig(c IdentityProviderConverter, input flightdeckv1.IdentityProvider) (identityproviderresource.IdentityProviderModel, error) {
	identityprovider := identityproviderresource.IdentityProviderModel{
		OrganizationId: basetypes.NewStringValue(input.OrganizationId.String()),
		TenantName:     basetypes.NewStringValue(input.TenantName),
		Github:         identityproviderresource.NewGithubValueNull(),
		Gitlab:         identityproviderresource.NewGitlabValueNull(),
	}

	githubTypes := identityproviderresource.GithubValue{}.AttributeTypes(context.TODO())
	gitlabTypes := identityproviderresource.GitlabValue{}.AttributeTypes(context.TODO())

	configType, err := input.ProviderConfig.Discriminator()
	if err != nil {
		return identityprovider, err
	}
	switch configType {
	case "github":
		gh, err := input.ProviderConfig.AsGithubIdentityProvider()
		if err != nil {
			return identityprovider, err
		}
		github, diag := basetypes.NewObjectValueFrom(context.TODO(), githubTypes, gh)
		if diag.HasError() {
			return identityprovider, fmt.Errorf("could not parse github identityprovider %+v", diag.Errors())
		}
		ghv, diag := identityproviderresource.NewGithubValue(githubTypes, github.Attributes())
		diag = github.As(context.TODO(), github, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return identityprovider, fmt.Errorf("could not parse github identityprovider %+v", diag.Errors())
		}
		identityprovider.Github = ghv
	case "gitlab":
		gl, err := input.ProviderConfig.AsGitlabIdentityProvider()
		if err != nil {
			return identityprovider, err
		}
		gitlab, diag := basetypes.NewObjectValueFrom(context.TODO(), gitlabTypes, gl)
		if diag.HasError() {
			return identityprovider, fmt.Errorf("could not parse gitlab identityprovider %+v", diag.Errors())
		}
		glv, diag := identityproviderresource.NewGithubValue(gitlabTypes, gitlab.Attributes())
		diag = gitlab.As(context.TODO(), gitlab, basetypes.ObjectAsOptions{})
		if diag.HasError() {
			return identityprovider, fmt.Errorf("could not parse gitlab identityprovider %+v", diag.Errors())
		}
		identityprovider.Github = glv
	}
	return identityprovider, nil
}
