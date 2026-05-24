// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package sdk

import goplugin "github.com/hashicorp/go-plugin"

const (
	// ProtocolVersion identifies the go-plugin protocol for semantic_release.v1.
	ProtocolVersion = 1

	// MagicCookieKey and MagicCookieValue protect against accidental execution
	// of non-semrel binaries by requiring a matching handshake cookie.
	MagicCookieKey   = "SEMREL_PLUGIN"
	MagicCookieValue = "v1"
)

// HandshakeConfig is the shared go-plugin handshake used by host and plugins.
var HandshakeConfig = goplugin.HandshakeConfig{
	ProtocolVersion:  ProtocolVersion,
	MagicCookieKey:   MagicCookieKey,
	MagicCookieValue: MagicCookieValue,
}
