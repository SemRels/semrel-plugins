// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package sdk

import (
	"io"
	"os"

	"github.com/hashicorp/go-hclog"
)

// NewLogger returns a logger configured for plugin binaries.
//
// ADR-001 requires stdout to stay reserved for the go-plugin handshake,
// therefore logs are always written to stderr.
func NewLogger(name string) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:   name,
		Level:  hclog.Info,
		Output: os.Stderr,
	})
}

// NewLoggerWithOutput allows tests/custom setups while preserving sane defaults.
func NewLoggerWithOutput(name string, output io.Writer) hclog.Logger {
	if output == nil {
		output = os.Stderr
	}

	return hclog.New(&hclog.LoggerOptions{
		Name:   name,
		Level:  hclog.Info,
		Output: output,
	})
}
