[![Build Status](https://github.com/foomo/go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/foomo/go/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/go)](https://goreportcard.com/report/github.com/foomo/go)
[![GoDoc](https://godoc.org/github.com/foomo/go?status.svg)](https://godoc.org/github.com/foomo/go)

<p align="center">
  <img alt="go" src="docs/public/logo.png" width="400" height="400"/>
</p>

# go

Go standard library extension, adding the missing parts used in the foomo ecosystem to keep [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself).

## Features

- **fmt** — Template string formatting with `%{.key}` syntax
- **net** — Free port allocation and TCP connection wait helpers (`FreePort`, `FreePorts`, `WaitFor`, `WaitForFreePort`)
- **options** — Generic functional options pattern (`Option`, `OptionE`, `Builder`, `BuilderE`)
- **os** — Typed `Getenv`/`MustGetenv` for scalars, slices, and maps with defaults; path `Expand` (`~/`, env vars)
- **runtime** — Enriched caller introspection (`Caller`, `CallerFunc`, `StackTrace`) and panic recovery
- **sec** — Safe path joining to prevent directory traversal (gosec G304)
- **slices** — Generic slice utilities: `Filter`, `Map`, `GroupBy` (with error variants)
- **slog** — Test-friendly `slog.Handler` that writes to `testing.TB` output
- **strings** — Case conversions, padding, validation, prefix/suffix matching, and composition
- **testing** — Tag-based test filtering via `GO_TEST_TAGS`, crypto key helpers, `ExampleTB`
- **time** — Context-aware `Sleep` and polling `WaitFor`
- **types** — Common interface contracts (`Closer`, `Starter`, `Stopper`, …) with function adapters and `As<X>` helpers

## How to Contribute

Contributions are welcome! Please read the [contributing guide](docs/CONTRIBUTING.md).

![Contributors](https://contributors-table.vercel.app/image?repo=foomo/go&width=50&columns=15)

## License

Distributed under MIT License, please read the [license file](LICENSE) for more details.

_Made with ♥ [foomo](https://www.foomo.org) by [bestbytes](https://www.bestbytes.com)_
