# Official plugin integration matrix

This is the durable 43-plugin baseline for the first-party catalog. Names use
the canonical migration format `@semrel/<type>-<name>`. Only canonical,
direct-child repositories in the workspace count: `repos/` and `tmp-repos/`
are duplicate checkouts and are deliberately excluded.

The SDK test helpers in `semrel/pkg/sdk/sdktest` are not imported here: these
independent modules cannot depend on an unreleased sibling through a local
`replace` without leaking that replacement to consumers. Their standard-library
HTTP contract tests provide the same hermetic isolation until the helper is
released as a versioned module.

| Canonical plugin | Direct repository | Risk / deterministic primary-effect coverage | Gate |
|---|---|---|---|
| `@semrel/analyzer-conventional` | `analyzer-conventional` | JSON bump conformance; `TestRunWritesAnalysisResult`, invalid commits | local |
| `@semrel/analyzer-default` | `analyzer-default` | JSON bump conformance; `TestRunWritesAnalysisResult`, invalid patterns | local |
| `@semrel/condition-generic` | `condition-generic` | condition command/env conformance; `TestRun_Success`, `TestRun_Failure` | local |
| `@semrel/condition-gitea-actions` | `condition-gitea-actions` | CI environment conformance; `TestRun_Success`, `TestRun_Failure` | local |
| `@semrel/condition-github-actions` | `condition-github-actions` | CI environment conformance; `TestRun_Success`, `TestRun_Failure` | local |
| `@semrel/condition-gitlab-ci` | `condition-gitlab-ci` | CI environment conformance; `TestRun_Success`, `TestRun_Failure` | local |
| `@semrel/generator-changelog-html` | `generator-changelog-html` | golden HTML output; `TestRunWritesHTML`, invalid commit JSON | local |
| `@semrel/generator-changelog-md` | `generator-changelog-md` | changelog artifact mutation/no-mutation; `TestRunWithCompressionWritesChangelog`, `TestRunWithCompressionDryRunSkipsWrite` | local |
| `@semrel/generator-release-notes` | `generator-release-notes` | golden release-notes output; `TestRunWritesReleaseNotes`, invalid commit JSON | local |
| `@semrel/hook-discord` | `hook-discord` | webhook payload, rate-limit retry, error and dry-run; `TestDiscordNotifierNotifyHandlesRateLimitRetry`, `TestRunDryRunPrintsPayloadJSON` | local |
| `@semrel/hook-email` | `hook-email` | SMTP envelope/payload, validation and dry-run; `TestMailerNotify_Success`, `TestRun_DryRun`; real Mailpit delivery uses `integration` tag | local + Linux CI |
| `@semrel/hook-gitplugin` | `hook-gitplugin` | git tag/commit/push and dry-run; `TestCreateTagLightweight`, `TestRunDryRun` | local |
| `@semrel/hook-jira` | `hook-jira` | API create/release errors and dry-run; `TestCreateVersionSuccess`, `TestRunDryRun` | local |
| `@semrel/hook-matrix` | `hook-matrix` | webhook payload, retry/error/dry-run; `TestMatrixNotifierNotifyRetriesOnServerError`, `TestRunDryRun` | local |
| `@semrel/hook-slack` | `hook-slack` | webhook payload, retry/error/dry-run; `TestSlackNotifierNotifyRetriesOnServerError`, `TestRunDryRun` | local |
| `@semrel/hook-teams` | `hook-teams` | webhook payload, retry/error/dry-run; `TestTeamsNotifierNotifyRetriesOnServerError`, `TestRunSuccess` | local |
| `@semrel/packager-nfpm` | `packager-nfpm` | command plan/no-mutation; `TestIntegrationDryRunBuildPlanAndInvalidVersion`; real `.deb` build and inspection use `integration` tag | local + Linux CI |
| `@semrel/provider-bitbucket` | `provider-bitbucket` | authenticated release/download API behavior, error and dry-run; `TestCreateRelease_Success`, `TestRunDryRun` | local |
| `@semrel/provider-git` | `provider-git` | real local bare remote tag/branch push; `TestPushTagPushesTagToRemoteRepository`, `TestRunSkipsPushesInDryRunMode` | local |
| `@semrel/provider-gitea` | `provider-gitea` | release API contracts, error and dry-run; `TestCreateReleaseSuccess`, `TestRunDryRun` | local |
| `@semrel/provider-github` | `provider-github` | release/asset API contracts, error and dry-run; `TestCreateReleaseSuccess`, `TestRunDryRun` | local |
| `@semrel/provider-gitlab` | `provider-gitlab` | release API contracts, error and dry-run; `TestCreateReleaseSuccess`, `TestRunDryRun` | local |
| `@semrel/publisher-crates` | `publisher-crates` | Cargo invocation/token/error/dry-run conformance; `TestBuildInvocationsWithSinglePackageManifestAndDryRun`, `TestPublishStopsOnFirstPackageFailure` | local |
| `@semrel/publisher-docker` | `publisher-docker` | strict config/reference/version handling, inspect/tag skip-or-replace/verification, exactly one push, digest fallback, cancellation and secret-safe failures; real image build, registry publish/digest comparison, remove/pull/inspect, and no-mutation dry-run use the `integration` tag | local + Linux CI |
| `@semrel/publisher-generic-http` | `publisher-generic-http` | strict stateful HTTP request/body/auth contract, 429 retry, error, dry-run, and secret non-disclosure; `TestIntegrationPublishRetriesRateLimitWithExactRequestContract` | local |
| `@semrel/publisher-npm` | `publisher-npm` | package manager/token/error/dry-run conformance; `TestRunDryRunDoesNotExecute`, `TestRunRequiresTokenOutsideDryRun` | local |
| `@semrel/publisher-oci` | `publisher-oci` | OCI reference/artifact plan, invalid config and dry-run; real distribution-registry publish/manifest inspection uses `integration` tag | local + Linux CI |
| `@semrel/publisher-pypi` | `publisher-pypi` | build/upload command, token/error/dry-run conformance; `TestPublishDryRunBuildsAndSkipsUpload`, `TestPublishPropagatesUploadFailure` | local |
| `@semrel/updater-cargo` | `updater-cargo` | `Cargo.toml` mutation/no-mutation; `TestRunUpdatesCargoToml`, `TestRunDryRun` | local |
| `@semrel/updater-composer` | `updater-composer` | `composer.json` mutation/no-mutation; `TestRunUpdatesComposerJSON`, `TestRunDryRunDoesNotWriteFile` | local |
| `@semrel/updater-docker` | `updater-docker` | Dockerfile mutation/no-mutation; `TestRunUpdatesDockerfile`, `TestRunDryRun` | local |
| `@semrel/updater-go` | `updater-go` | Go version-file mutation/no-mutation; `TestRunUpdatesFile`, `TestRunDryRun` | local |
| `@semrel/updater-gradle` | `updater-gradle` | Gradle mutation/no-mutation; `TestRunUpdatesGradleFile`, `TestRunDryRun` | local |
| `@semrel/updater-helm` | `updater-helm` | Chart mutation/no-mutation; `TestRunUpdatesVersionOnlyByDefault`, `TestRunDryRun` | local |
| `@semrel/updater-homebrew` | `updater-homebrew` | Formula mutation/no-mutation; `TestRunUpdatesFormula`, `TestRunDryRun` | local |
| `@semrel/updater-maven` | `updater-maven` | POM mutation/no-mutation; `TestRunUpdatesPom`, `TestRunDryRun` | local |
| `@semrel/updater-npm` | `updater-npm` | `package.json` mutation/no-mutation; `TestRunUpdatesPackageJSON`, `TestRunDryRun` | local |
| `@semrel/updater-nuget` | `updater-nuget` | project-file mutation/no-mutation; `TestRunUpdatesProjectFile`, `TestRunDryRun` | local |
| `@semrel/updater-pubspec` | `updater-pubspec` | `pubspec.yaml` mutation/no-mutation; `TestRunUpdatesPubspec`, `TestRunDryRunPrintsComputedVersionLine` | local |
| `@semrel/updater-python` | `updater-python` | Python metadata mutation/no-mutation; `TestRunUpdatesPyproject`, `TestRunDryRun` | local |
| `@semrel/updater-terraform` | `updater-terraform` | Terraform mutation/no-mutation; `TestRunUpdatesTerraformFile`, `TestRunDryRun` | local |
| `@semrel/condition-bitbucket-pipelines` | `condition-bitbucket-pipelines` | CI environment conformance; `TestCheck_HappyPath`, `TestRun_Failure` | local |
| `@semrel/condition-circleci` | `condition-circleci` | CI environment conformance; `TestCheck_HappyPath`, `TestRun_Failure` | local |

## Validation policy

Run `go test ./...` in each direct repository, with a per-repository timeout.
No test is represented as passing merely because it is skipped. Docker-only
tests have the `integration` build tag and are required by Linux CI jobs: a
real distribution registry plus Docker for image publication, a registry plus
ORAS for OCI artifacts, Mailpit SMTP/API, and an actual nFPM package inspected
with `dpkg-deb`. These jobs fail when their service or tool is unavailable;
they do not silently skip.

## 2026-07-23 validation record

The timed `go test -timeout 90s ./...` sweep completed for the eight cloned
repositories (8/8). The prior direct-repository sweep passed for 33/34
repositories; `packager-nfpm/cmd/plugin` passed and its tagged integration test
compiled, while Windows application control blocks only its internal-plugin
generated test executable. That host limitation is not a production gap:
Linux CI now performs the actual nFPM build and inspection. Tagged OCI, SMTP,
and nFPM tests compile locally and run only in their required Linux CI gates.
