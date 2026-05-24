package sdk

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestLocalPluginPath(t *testing.T) {
	got := LocalPluginPath("commit-analyzer-default", "v0.1.0", "plugin")
	want := filepath.Join(
		LocalPluginRoot,
		runtime.GOOS+"_"+runtime.GOARCH,
		"commit-analyzer-default",
		"v0.1.0",
		"plugin",
	)
	if got != want {
		t.Fatalf("unexpected local plugin path: got %q want %q", got, want)
	}
}
