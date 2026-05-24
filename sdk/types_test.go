package sdk

import "testing"

func TestPluginTypesAreStable(t *testing.T) {
	if PluginTypeProvider != "provider" {
		t.Fatalf("unexpected provider type: %q", PluginTypeProvider)
	}
	if PluginTypeCICondition != "ci-condition" {
		t.Fatalf("unexpected ci type: %q", PluginTypeCICondition)
	}
	if PluginTypeCommitAnalyzer != "commit-analyzer" {
		t.Fatalf("unexpected commit analyzer type: %q", PluginTypeCommitAnalyzer)
	}
	if PluginTypeChangelogGenerator != "changelog-generator" {
		t.Fatalf("unexpected changelog type: %q", PluginTypeChangelogGenerator)
	}
	if PluginTypeFilesUpdater != "files-updater" {
		t.Fatalf("unexpected files updater type: %q", PluginTypeFilesUpdater)
	}
	if PluginTypeHooks != "hooks" {
		t.Fatalf("unexpected hooks type: %q", PluginTypeHooks)
	}
}
