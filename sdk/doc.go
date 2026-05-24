// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

// Package sdk contains shared helpers for semrel out-of-process plugins.
//
// It intentionally keeps a small surface so plugin binaries can bootstrap
// quickly while the generated gRPC bindings evolve in lockstep with the
// semrel protocol.
package sdk
