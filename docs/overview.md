# Overview

[简体中文](overview.zh-CN.md) | [Docs Index](README.md)

## Purpose

Unix domain stream and datagram transport for Gnalloy, including Linux peer credentials and fd passing.

This module owns an I/O boundary. It creates or adapts Gnalloy Channels for a concrete transport while protocol parsing, business handlers, TLS policy, and observability remain in other modules.

## Repository Identity

- Module path: `gnalloy.org/transport-unix`
- GitHub repository: `github.com/gnalloy/transport-unix`
- Default branch: `dev`
- License: Apache-2.0

## Package Map
- `gnalloy.org/transport-unix` (`unix`)

## Direct Gnalloy Dependencies
- `gnalloy.org/gnalloy`

## Direct Dependents in the Current Module Plan
- No repository in the current module plan depends on this module directly.

## Architecture Position

Gnalloy keeps the core small and dependency-light. This repository is a replaceable module around one responsibility, connected through explicit Go packages instead of runtime discovery.

## Compatibility

The public import path is `gnalloy.org/transport-unix`. Until the first stable tag is published, use `@dev` or an explicit pseudo-version selected by your dependency policy.
