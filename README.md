# semrel-plugins

Official first-party plugin catalog for [semrel](https://github.com/SemRels/semrel) — the semantic release tool.

All plugins follow the semrel plugin contract: they are standalone binaries that communicate via environment variables, write JSON to stdout (analyzers only), and exit with a meaningful code.

## Official plugins

### Analyzers

Analyzers inspect commits and determine the next semantic version.

| Plugin | Repository | Description |
|---|---|---|
| `analyzer-conventional` | [SemRels/analyzer-conventional](https://github.com/SemRels/analyzer-conventional) | Conventional Commits (feat/fix/BREAKING CHANGE) |
| `analyzer-default` | [SemRels/analyzer-default](https://github.com/SemRels/analyzer-default) | Simple keyword-based bump detection |

### Generators

Generators produce release artefacts from commit history.

| Plugin | Repository | Description |
|---|---|---|
| `generator-changelog-md` | [SemRels/generator-changelog-md](https://github.com/SemRels/generator-changelog-md) | Markdown changelog (CHANGELOG.md) |
| `generator-changelog-html` | [SemRels/generator-changelog-html](https://github.com/SemRels/generator-changelog-html) | HTML changelog |
| `generator-release-notes` | [SemRels/generator-release-notes](https://github.com/SemRels/generator-release-notes) | GitHub/GitLab release notes body |

### Providers

Providers publish the release to a hosting platform.

| Plugin | Repository | Description |
|---|---|---|
| `provider-github` | [SemRels/provider-github](https://github.com/SemRels/provider-github) | GitHub Releases + asset upload |
| `provider-gitlab` | [SemRels/provider-gitlab](https://github.com/SemRels/provider-gitlab) | GitLab Releases |
| `provider-gitea` | [SemRels/provider-gitea](https://github.com/SemRels/provider-gitea) | Gitea Releases |
| `provider-bitbucket` | [SemRels/provider-bitbucket](https://github.com/SemRels/provider-bitbucket) | Bitbucket Cloud |
| `provider-git` | [SemRels/provider-git](https://github.com/SemRels/provider-git) | Plain git tag push |

### Publishers

Publishers push language- or registry-specific artifacts after semrel has prepared the release.

| Plugin | Repository | Description |
|---|---|---|
| `publisher-crates` | [SemRels/publisher-crates](https://github.com/SemRels/publisher-crates) | Publish crate to crates.io |
| `publisher-npm` | [SemRels/publisher-npm](https://github.com/SemRels/publisher-npm) | Publish package to the npm registry |

### Conditions

Conditions gate the release pipeline — a non-zero exit aborts the release.

| Plugin | Repository | Description |
|---|---|---|
| `condition-github-actions` | [SemRels/condition-github-actions](https://github.com/SemRels/condition-github-actions) | Ensures running in GitHub Actions on the correct branch |
| `condition-gitlab-ci` | [SemRels/condition-gitlab-ci](https://github.com/SemRels/condition-gitlab-ci) | Ensures running in GitLab CI on the correct branch |
| `condition-gitea-actions` | [SemRels/condition-gitea-actions](https://github.com/SemRels/condition-gitea-actions) | Ensures running in Gitea Actions |
| `condition-generic` | [SemRels/condition-generic](https://github.com/SemRels/condition-generic) | Generic branch/env condition checks |

### Hooks

Hooks run before or after release for notifications and integrations.

| Plugin | Repository | Description |
|---|---|---|
| `hook-slack` | [SemRels/hook-slack](https://github.com/SemRels/hook-slack) | Slack release notification |
| `hook-teams` | [SemRels/hook-teams](https://github.com/SemRels/hook-teams) | Microsoft Teams notification |
| `hook-email` | [SemRels/hook-email](https://github.com/SemRels/hook-email) | Email notification via SMTP |
| `hook-matrix` | [SemRels/hook-matrix](https://github.com/SemRels/hook-matrix) | Matrix (Element) notification |
| `hook-jira` | [SemRels/hook-jira](https://github.com/SemRels/hook-jira) | Jira version tracking |
| `hook-gitplugin` | [SemRels/hook-gitplugin](https://github.com/SemRels/hook-gitplugin) | Git-based plugin hooks |

### Updaters

Updaters bump version strings in project files after the release version is determined.

| Plugin | Repository | Description |
|---|---|---|
| `updater-npm` | [SemRels/updater-npm](https://github.com/SemRels/updater-npm) | `package.json` version field |
| `updater-go` | [SemRels/updater-go](https://github.com/SemRels/updater-go) | Go module version constant |
| `updater-cargo` | [SemRels/updater-cargo](https://github.com/SemRels/updater-cargo) | `Cargo.toml` version field |
| `updater-python` | [SemRels/updater-python](https://github.com/SemRels/updater-python) | `pyproject.toml` / `setup.py` version |
| `updater-maven` | [SemRels/updater-maven](https://github.com/SemRels/updater-maven) | `pom.xml` version element |
| `updater-gradle` | [SemRels/updater-gradle](https://github.com/SemRels/updater-gradle) | `build.gradle` / `build.gradle.kts` version |
| `updater-docker` | [SemRels/updater-docker](https://github.com/SemRels/updater-docker) | `Dockerfile` ARG/LABEL version |
| `updater-helm` | [SemRels/updater-helm](https://github.com/SemRels/updater-helm) | Helm `Chart.yaml` version |
| `updater-terraform` | [SemRels/updater-terraform](https://github.com/SemRels/updater-terraform) | Terraform module version constraint |
| `updater-nuget` | [SemRels/updater-nuget](https://github.com/SemRels/updater-nuget) | NuGet `.csproj` version |
| `updater-homebrew` | [SemRels/updater-homebrew](https://github.com/SemRels/updater-homebrew) | Homebrew formula version |

### Publishers

Publishers push built artifacts/packages to a package registry after the version bump.

| Plugin | Repository | Description |
|---|---|---|
| `publisher-generic-http` | [SemRels/publisher-generic-http](https://github.com/SemRels/publisher-generic-http) | Upload artifacts to a generic HTTP(S) endpoint |
| `publisher-oci` | [SemRels/publisher-oci](https://github.com/SemRels/publisher-oci) | Publish artifacts to OCI registries via `oras push` |

### Packagers

Packagers build distributable OS/language packages from release artifacts.

| Plugin | Repository | Description |
|---|---|---|
| `packager-nfpm` | [SemRels/packager-nfpm](https://github.com/SemRels/packager-nfpm) | Build `.deb`/`.rpm`/`.apk` packages via nFPM |

---

## Plugin contract

See the [plugin development guide](https://semrel.io/guide/plugin-development) for a step-by-step tutorial on writing custom plugins.

All plugins communicate via:

- **Environment variables** → plugin: release context (`SEMREL_*`) and plugin config (`SEMREL_PLUGIN_*`)
- **stdout** (analyzers only): JSON analysis result
- **stderr**: logs and `plugin_schema_version=N` announcement
- **Exit code**: 0 = success, non-zero = abort release

---

## Contributing

Each plugin lives in its own repository under the [SemRels](https://github.com/SemRels) organization.
See [CONTRIBUTING.md](CONTRIBUTING.md) if present, or open an issue in the relevant plugin repo.
