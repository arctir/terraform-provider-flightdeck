// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"

	flightdeckv1 "github.com/arctir/go-flightdeck/pkg/api/v1"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func configure(ctx context.Context, providerData any, diagnostics *diag.Diagnostics) *flightdeckv1.ClientWithResponses {
	if providerData == nil {
		return nil
	}

	client, ok := providerData.(*flightdeckv1.ClientWithResponses)
	if !ok {
		diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *arctir.ClientWithResponses, got: %T. Please report this issue to the provider developers.", providerData),
		)
		return nil
	}
	return client
}
