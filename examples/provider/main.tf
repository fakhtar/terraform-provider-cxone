# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    cxone = {
      source  = "fakhtar/cxone"
      version = "~> 0.1"
    }
  }
  required_version = ">= 1.0"
}

# Authentication values should be provided via environment variables:
#   CXONE_ACCESS_KEY_ID     - NiCE CXone Access Key ID
#   CXONE_ACCESS_KEY_SECRET - NiCE CXone Access Key Secret
#   CXONE_REGION            - NiCE CXone region code (e.g. na1, eu1, au1)
#   COGNIGY_API_KEY         - Cognigy.AI API key
#   COGNIGY_URL             - Cognigy.AI base URL

provider "cxone" {
  region      = "na1"
  cognigy_url = "https://api.cognigy.ai"

  # access_key_id, access_key_secret, and cognigy_api_key are sensitive
  # and should be set via environment variables rather than in configuration.
}