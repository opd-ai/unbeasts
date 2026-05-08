# untemplate

[![CI](https://github.com/opd-ai/untemplate/actions/workflows/ci.yml/badge.svg)](https://github.com/opd-ai/untemplate/actions/workflows/ci.yml)

Generic repository template for planned `opd-ai/unsuite` projects.

## Overview

`untemplate` mirrors the directory, module/library, and documentation structure used by `opd-ai/unpeople`, while keeping implementation intentionally minimal and generic.

## Structure

- Root Go library package (`github.com/opd-ai/untemplate`)
- CLI binaries under `cmd/untemplated` and `cmd/untemplate-server`
- Integration package under `kaiju`
- Example app under `example`
- Project docs under `docs`
- CI workflow under `.github/workflows/ci.yml`

## Quick Start

```bash
go test ./...
go build ./...
```

## Documentation

- [Architecture](docs/architecture.md)
- [Attachment Slots](docs/attachment-slots.md)
- [Face Mesh Template](docs/face-mesh-template.md)
- [Kaiju Integration](docs/kaiju-integration.md)
- [REST API](docs/rest-api.md)
- [Vertex Merging](docs/vertex-merging.md)
