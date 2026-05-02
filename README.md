# Terraform Provider for NiCE CXone

[![Tests](https://github.com/fakhtar/terraform-provider-cxone/actions/workflows/test.yml/badge.svg)](https://github.com/fakhtar/terraform-provider-cxone/actions/workflows/test.yml)
[![Lint](https://github.com/fakhtar/terraform-provider-cxone/actions/workflows/lint.yml/badge.svg)](https://github.com/fakhtar/terraform-provider-cxone/actions/workflows/lint.yml)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)
[![Go Version](https://img.shields.io/github/go-mod/go-version/fakhtar/terraform-provider-cxone)](go.mod)

> ⚠️ **Community Provider** — This is an independently developed community
> provider and is not officially affiliated with, endorsed by, or supported
> by NiCE Ltd. or Cognigy GmbH. For issues please
> [open a GitHub issue](https://github.com/fakhtar/terraform-provider-cxone/issues).

---

## Overview

This provider enables teams to manage **NiCE CXone** and **Cognigy.AI**
resources using Terraform, bringing Infrastructure as Code workflows to
enterprise AI contact center deployments.

Built on the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework),
this provider manages NiCE CXone Mpower and Cognigy.AI resources as
first-class Terraform objects — giving enterprise teams the same IaC
workflows they use for the rest of their infrastructure.

### Why This Exists

NiCE CXone and Cognigy.AI are world-class AI contact center platforms with
no native Terraform support. Enterprise teams building on CXone Mpower today
configure their contact center resources manually or through ad-hoc scripting.
This provider changes that.

---

## Requirements

| Tool | Version |
|---|---|
| [Terraform](https://developer.hashicorp.com/terraform/downloads) | >= 1.0 |
| [Go](https://golang.org/doc/install) | >= 1.21 (for local development) |
| NiCE CXone tenant | With API access key credentials |
| Cognigy.AI instance | With API key |

---

## Quick Start

```hcl
terraform {
  required_providers {
    cxone = {
      source  = "fakhtar/cxone"
      version = "~> 0.1"
    }
  }
  required_version = ">= 1.0"
}

provider "cxone" {
  region      = "na1"
  cognigy_url = "https://api.cognigy.ai"

  # Sensitive credentials should be set via environment variables:
  #   CXONE_ACCESS_KEY_ID
  #   CXONE_ACCESS_KEY_SECRET
  #   COGNIGY_API_KEY
}
```

---

## Authentication

### NiCE CXone

CXone uses OAuth2 server-to-server authentication. You will need an
**Access Key ID** and **Access Key Secret** generated from a dedicated
API user account in the CXone Mpower Admin UI.

Set credentials via environment variables (recommended):

```bash
export CXONE_ACCESS_KEY_ID="your-access-key-id"
export CXONE_ACCESS_KEY_SECRET="your-access-key-secret"
export CXONE_REGION="na1"
```

Or configure directly in the provider block (not recommended for production):

```hcl
provider "cxone" {
  access_key_id     = "your-access-key-id"
  access_key_secret = "your-access-key-secret"
  region            = "na1"
}
```

### Cognigy.AI

Cognigy.AI uses API key authentication. Generate a key from your
Cognigy.AI user profile.

```bash
export COGNIGY_API_KEY="your-cognigy-api-key"
export COGNIGY_URL="https://api.cognigy.ai"
```

---

## MVP Resource Roadmap

The initial focus is a minimum viable conversational AI flow — the
equivalent of: *phone call arrives → Cognigy AI handles it → escalates
to a human agent.*

| Resource | Type | Status | Description |
|---|---|---|---|
| `cxone_business_unit` | Data Source | 🔄 In Progress | Read existing CXone tenant |
| `cxone_hours_of_operation` | Resource | 📋 Planned | When agents are available |
| `cxone_acd_skill` | Resource | 📋 Planned | The routing queue |
| `cxone_point_of_contact` | Resource | 📋 Planned | Phone number entry point |
| `cxone_virtual_agent_hub` | Resource | 📋 Planned | Cognigy ↔ CXone connector |
| `cxone_studio_script` | Resource | 📋 Planned | Contact routing logic |
| `cognigy_flow` | Resource | 📋 Planned | AI conversation design |
| `cognigy_endpoint` | Resource | 📋 Planned | Deploy flow to CXone channel |
| `cognigy_handover_provider` | Resource | 📋 Planned | Escalation to human agent |

### Production Resources (v2+)

| Resource | Type | Status | Description |
|---|---|---|---|
| `cxone_campaign` | Resource | 🗓 Backlog | Groups skills for reporting |
| `cxone_acd_user` | Resource | 🗓 Backlog | Agent user management |
| `cognigy_snapshot` | Resource | 🗓 Backlog | Versioned flow for production |
| `cxone_agent_assist_hub` | Resource | 🗓 Backlog | Cognigy Copilot for agents |

---

## Resource Dependency Order

data.cxone_business_unit
↓
cxone_hours_of_operation
↓
cxone_acd_skill
↓
cognigy_flow → cognigy_endpoint
↓
cxone_virtual_agent_hub
↓
cxone_studio_script
↓
cxone_point_of_contact
↓
cognigy_handover_provider


---

## Development

### Requirements

- [Go](https://golang.org/doc/install) >= 1.21
- [Git Bash](https://git-scm.com/download/win) (Windows)
- [make](https://community.chocolatey.org/packages/make) >= 4.4

### Build

```bash
make build
```

### Test

```bash
# Unit tests
make test

# Acceptance tests (requires live tenant credentials)
export CXONE_ACCESS_KEY_ID="..."
export CXONE_ACCESS_KEY_SECRET="..."
export CXONE_REGION="..."
export COGNIGY_API_KEY="..."
export COGNIGY_URL="..."
make testacc
```

### Lint

```bash
make lint
```

---

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for the full development guide
including local setup, coding standards, and the pull request process.

---

## License

[Mozilla Public License 2.0](./LICENSE)

---

## Acknowledgements

This provider was initiated by [@fakhtar](https://github.com/fakhtar) to
bring Infrastructure as Code practices to NiCE CXone and Cognigy.AI
deployments. The long-term goal is for this provider to be adopted
officially by NiCE and maintained as a Partner provider on the Terraform
Registry.