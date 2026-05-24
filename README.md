# go-semrel-plugins

Shared SDK helpers for semrel out-of-process plugins.

This repository contains the Go-side bootstrap utilities used by plugin binaries
defined in ADR-001 (go-plugin + gRPC transport).

## Install

```bash
go get github.com/GoSemantics/go-semrel-plugins
```

## Available SDK helpers

- `sdk.HandshakeConfig`: shared go-plugin handshake config (protocol + cookie)
- `sdk.NewLogger`: logger configured to write to stderr (stdout is handshake-only)
- `sdk.LocalPluginPath`: local discovery path builder
- `sdk.PluginType`: constants for supported plugin categories

## Example

```go
package main

import (
	goplugin "github.com/hashicorp/go-plugin"

	"github.com/GoSemantics/go-semrel-plugins/sdk"
)

func main() {
	logger := sdk.NewLogger("my-plugin")

	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: sdk.HandshakeConfig,
		Plugins:         map[string]goplugin.Plugin{},
		GRPCServer:      goplugin.DefaultGRPCServer,
		Logger:          logger,
	})
}
```
