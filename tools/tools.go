// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build tools

package tools

import (
	// tfplugindocs generates documentation for the Terraform Registry
	// from provider schema and examples.
	// Run: go generate ./...
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)