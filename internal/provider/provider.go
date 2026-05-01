// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure CxoneProvider satisfies the provider.Provider interface.
var _ provider.Provider = &CxoneProvider{}

// CxoneProvider defines the provider implementation.
type CxoneProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// New returns a function that creates a new provider instance.
// Called from main.go to start the provider server.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CxoneProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name and version.
func (p *CxoneProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cxone"
	resp.Version = p.version
}

// Schema defines the provider-level configuration block.
// This is what users configure in their provider "cxone" {} block.
func (p *CxoneProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The CXone provider enables management of NiCE CXone and Cognigy.AI resources using Terraform.",
		Attributes: map[string]schema.Attribute{
			"access_key_id": schema.StringAttribute{
				MarkdownDescription: "Access Key ID for NiCE CXone OAuth2 authentication. Can also be set via the `CXONE_ACCESS_KEY_ID` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"access_key_secret": schema.StringAttribute{
				MarkdownDescription: "Access Key Secret for NiCE CXone OAuth2 authentication. Can also be set via the `CXONE_ACCESS_KEY_SECRET` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "NiCE CXone region for the authentication endpoint (e.g. `us-east-1`). Can also be set via the `CXONE_REGION` environment variable.",
				Optional:            true,
			},
			"cognigy_api_key": schema.StringAttribute{
				MarkdownDescription: "API key for Cognigy.AI authentication. Can also be set via the `COGNIGY_API_KEY` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"cognigy_url": schema.StringAttribute{
				MarkdownDescription: "Base URL for the Cognigy.AI API (e.g. `https://api.cognigy.ai`). Can also be set via the `COGNIGY_URL` environment variable.",
				Optional:            true,
			},
		},
	}
}

// Configure sets up any provider-level shared clients or data.
// Runs before any resource or data source operations.
func (p *CxoneProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	// HTTP client setup will be added here when internal/client is implemented.
}

// Resources returns the list of resources this provider supports.
func (p *CxoneProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// Resources will be registered here as they are built.
	}
}

// DataSources returns the list of data sources this provider supports.
func (p *CxoneProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Data sources will be registered here as they are built.
	}
}
