# Contributing to terraform-provider-cxone

Thank you for your interest in contributing. This provider brings
Infrastructure as Code to NiCE CXone and Cognigy.AI — your contributions
help make that vision a reality.

---

## Code of Conduct

This project follows the
[HashiCorp Community Guidelines](https://www.hashicorp.com/community-guidelines).
Be respectful and constructive in all interactions.

---

## Requirements

| Tool | Version | Install |
|---|---|---|
| Go | >= 1.21 | [golang.org/dl](https://golang.org/dl) |
| Terraform | >= 1.0 | [developer.hashicorp.com/terraform/downloads](https://developer.hashicorp.com/terraform/downloads) |
| golangci-lint | Latest | [golangci-lint.run](https://golangci-lint.run/welcome/install/) |
| Git | Any recent | [git-scm.com](https://git-scm.com/download/win) |
| make | Any recent | `choco install make` |

> **Windows users:** All commands in this guide use **Git Bash**.
> Open Git Bash from the Start menu or from within VS Code by selecting
> Git Bash as your terminal shell.

---

## Getting Started

### 1. Fork and Clone

```bash
git clone https://github.com/YOUR_USERNAME/terraform-provider-cxone.git
cd terraform-provider-cxone
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Build

```bash
make build
```

### 4. Configure Local Dev Override

To test the provider locally against a real Terraform configuration,
add a dev override to your Terraform CLI config file.

On Windows the file is located at `$APPDATA/terraform.rc`. Create it
if it does not exist:

```hcl
provider_installation {
  dev_overrides {
    "fakhtar/cxone" = "C:/Users/YOUR_USERNAME/go/bin"
  }
  direct {}
}
```

Then install the provider binary locally:

```bash
go install .
```

Terraform will now use your local build instead of downloading from
the Registry. Note: `terraform init` is not required when using
dev overrides.

---

## Development Workflow

### Common Commands

| Command | What It Does |
|---|---|
| `make build` | Compiles the provider |
| `make install` | Builds and installs to `$GOPATH/bin` |
| `make lint` | Runs golangci-lint |
| `make fmt` | Formats all Go code with gofmt |
| `make test` | Runs unit tests |
| `make testacc` | Runs acceptance tests (requires live tenant) |

### Running Unit Tests

```bash
make test
```

### Running Acceptance Tests

Acceptance tests create real resources in NiCE CXone and Cognigy.AI.
You will need valid credentials set as environment variables:

```bash
export CXONE_API_KEY="your-cxone-api-key"
export CXONE_REGION="us-east-1"
export COGNIGY_API_KEY="your-cognigy-api-key"
export COGNIGY_URL="https://api.cognigy.ai"
make testacc
```

---

## Pull Request Process

1. Open an issue first for significant changes to discuss the approach
2. Fork the repo and create a branch: `git checkout -b feat/my-feature`
3. Make your changes following the coding standards below
4. Run `make lint` and `make test` — both must pass cleanly
5. Update `CHANGELOG.md` under `## 0.x.x (Unreleased)` with your change
6. Open a pull request with a clear description of what and why

### Branch Naming

| Type | Pattern | Example |
|---|---|---|
| Feature | `feat/` | `feat/cognigy-flow-resource` |
| Bug fix | `fix/` | `fix/acd-skill-import` |
| Documentation | `docs/` | `docs/update-readme` |
| Chore | `chore/` | `chore/update-dependencies` |

---

## Coding Standards

- Follow standard Go conventions and idioms
- All resources must implement `ImportState` for `terraform import` support
- All attributes must have `MarkdownDescription` set
- Sensitive attributes (API keys, tokens) must have `Sensitive: true`
- New resources must have at least one acceptance test
- Run `make fmt` before committing — unformatted code will fail CI

---

## Reporting Issues

Use the GitHub issue templates:
- **Bug report** — something is not working as documented
- **Feature request** — a new resource or data source you need

Please search existing issues before opening a new one.

---

## Questions

If you are new to Terraform provider development, the
[HashiCorp Plugin Framework tutorial](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework)
is the best starting point.