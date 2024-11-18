// Copyright (c) Arctir, Inc.
// SPDX-License-Identifier: Apache-2.0



package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/arctir/terraform-provider-flightdeck/pkg/provider"
	_ "github.com/jmattheis/goverter" 
)

func main() {
	var debug bool
	var apiEndpoint string
	var configPath string

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.StringVar(&apiEndpoint, "api-endpoint", "", "the arctir api endpoint to use")
	flag.StringVar(&configPath, "config-path", "", "the path to the arctir configuration")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/arctir/flightdeck",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(
		basetypes.NewStringValue(apiEndpoint),
		basetypes.NewStringValue(configPath)), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
