// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package sdk

import (
	"path/filepath"
	"runtime"
)

const LocalPluginRoot = ".semrel"

// LocalPluginPath returns the ADR-001 compliant local plugin path:
// .semrel/<GOOS>_<GOARCH>/<plugin-name>/<version>/<binary>
func LocalPluginPath(pluginName, version, binaryName string) string {
	platform := runtime.GOOS + "_" + runtime.GOARCH
	return filepath.Join(LocalPluginRoot, platform, pluginName, version, binaryName)
}
