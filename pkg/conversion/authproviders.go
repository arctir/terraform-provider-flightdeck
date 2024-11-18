// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"context"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	authproviderresource "github.com/arctir/terraform-provider-flightdeck/pkg/provider/resource_auth_provider"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// goverter:converter
// goverter:extend StringValueToString
// goverter:extend StringToStringValue
// goverter:extend TimeToStringValue
// goverter:extend UUIDToStringValue
// goverter:extend StringValueToUUID
type AuthProviderConverter interface {
	// goverter:map . ProviderConfig | SetApiAuthProviderConfig
	ConvertToApiAuthProvider(source authproviderresource.AuthProviderModel) (flightdeckv1.AuthProviderInput, error)
	// goverter:ignoreMissing
	// foverter:ignore Github Gitlab
	// goverter:default SetTfAuthProviderConfig
	ConvertToTfAuthProvider(source flightdeckv1.AuthProvider) (authproviderresource.AuthProviderModel, error)
}

func SetApiAuthProviderConfig(source authproviderresource.AuthProviderModel) (flightdeckv1.AuthProviderInput_ProviderConfig, error) {
	config := flightdeckv1.AuthProviderInput_ProviderConfig{}

	if !source.Github.IsUnknown() && !source.Gitlab.IsUnknown() {
		return config, nil
	}
	if !source.Github.IsUnknown() {
		var gh flightdeckv1.GithubAuthProvider
		o, diag := source.Github.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github authprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gh, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse github authprovider %+v", diag.Errors())
		}
		gh.ConfigType = "github"
		config.FromGithubAuthProvider(gh)
	}
	if !source.Gitlab.IsUnknown() {
		var gl flightdeckv1.GitlabAuthProvider
		o, diag := source.Gitlab.ToObjectValue(context.TODO())
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab authprovider %+v", diag.Errors())
		}
		diag = o.As(context.TODO(), &gl, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if diag.HasError() {
			return config, fmt.Errorf("could not parse gitlab authprovider %+v", diag.Errors())
		}
		gl.ConfigType = "gitlab"
		config.FromGitlabAuthProvider(gl)
	}
	return config, nil
}

func SetTfAuthProviderConfig(c AuthProviderConverter, input flightdeckv1.AuthProvider) (authproviderresource.AuthProviderModel, error) {
	authprovider := authproviderresource.AuthProviderModel{
		OrganizationId: basetypes.NewStringValue(input.OrganizationId.String()),
		PortalName:     basetypes.NewStringValue(input.PortalName),
		Github:         authproviderresource.NewGithubValueNull(),
		Gitlab:         authproviderresource.NewGitlabValueNull(),
	}

	githubTypes := authproviderresource.GithubValue{}.AttributeTypes(context.TODO())
	gitlabTypes := authproviderresource.GitlabValue{}.AttributeTypes(context.TODO())

	configType, err := input.ProviderConfig.Discriminator()
	if err != nil {
		return authprovider, err
	}
	switch configType {
	case "github":
		gh, err := input.ProviderConfig.AsGithubAuthProvider()
		if err != nil {
			return authprovider, err
		}
		github, diag := basetypes.NewObjectValueFrom(context.TODO(), githubTypes, &gh)
		if diag.HasError() {
			return authprovider, fmt.Errorf("could not parse github authprovider %+v", diag.Errors())
		}
		ghv, diag := authproviderresource.NewGithubValue(githubTypes, github.Attributes())
		if diag.HasError() {
			return authprovider, fmt.Errorf("could not parse github authprovider %+v", diag.Errors())
		}
		authprovider.Github = ghv
	case "gitlab":
		gl, err := input.ProviderConfig.AsGitlabAuthProvider()
		if err != nil {
			return authprovider, err
		}
		gitlab, diag := basetypes.NewObjectValueFrom(context.TODO(), gitlabTypes, &gl)
		if diag.HasError() {
			return authprovider, fmt.Errorf("could not parse gitlab authprovider %+v", diag.Errors())
		}
		glv, diag := authproviderresource.NewGithubValue(gitlabTypes, gitlab.Attributes())
		if diag.HasError() {
			return authprovider, fmt.Errorf("could not parse gitlab authprovider %+v", diag.Errors())
		}
		authprovider.Github = glv
	}
	return authprovider, nil
}
