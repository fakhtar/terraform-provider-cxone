## Description

<!-- 
Describe what this PR does and why. Link to any related issues.
Example: "Closes #42 — adds cxone_acd_skill resource"
-->



## Type of Change

<!-- Check all that apply -->

- [ ] New resource
- [ ] New data source
- [ ] Bug fix
- [ ] Documentation update
- [ ] Dependency update
- [ ] Refactor / code quality
- [ ] CI/CD / tooling

## Related Issues

<!-- 
Link any related issues here.
Example: Closes #42
-->

## Testing

<!-- 
Describe how you tested this change. For new resources or bug fixes,
acceptance tests are required. For documentation changes, a manual
review is sufficient.
-->

- [ ] Unit tests pass (`make test`)
- [ ] Lint passes (`make lint`)
- [ ] Acceptance tests added or updated (for new resources / bug fixes)
- [ ] Manually tested against a live NiCE CXone or Cognigy.AI tenant

## Checklist

- [ ] Code follows the conventions in [CONTRIBUTING.md](../CONTRIBUTING.md)
- [ ] All new attributes have `MarkdownDescription` set
- [ ] Sensitive attributes have `Sensitive: true`
- [ ] New resources implement `ImportState`
- [ ] `CHANGELOG.md` updated under `Unreleased`
- [ ] Documentation generated (`make generate`) if schema changed
- [ ] No API keys, tokens, or secrets included in this PR