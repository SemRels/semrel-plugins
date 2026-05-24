// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package sdk

// PluginType represents one of the supported plugin categories.
type PluginType string

const (
	PluginTypeProvider           PluginType = "provider"
	PluginTypeCICondition        PluginType = "ci-condition"
	PluginTypeCommitAnalyzer     PluginType = "commit-analyzer"
	PluginTypeChangelogGenerator PluginType = "changelog-generator"
	PluginTypeFilesUpdater       PluginType = "files-updater"
	PluginTypeHooks              PluginType = "hooks"
)
