// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/fakhtar/terraform-provider-cxone/internal/provider"
)

// Run "go generate" to format example terraform files and generate the
// docs for the registry/website.
//
// If you do not have terraform installed, you can remove the formatting
// command, but it is suggested to ensure the documentation is formatted
// properly.
//
//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name cxone

// version is set by the GoReleaser configuration using ldflags at build time.
// It defaults to "dev" for local development builds.
var version string = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/fakhtar/cxone",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}