# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Roadmap

To be defined.

## [0.0.11] - 2022-03-4
### Changed
- Updated deps

## [0.0.10] - 2022-03-4
### Changed
- Update OS signalling.

## [0.0.9] - 2022-02-25
### Changed
- Standardizes WithXYZOptions to set params instead of adding.
- `New` is basic server, `NewDefault` is loaded one.

## [0.0.8] - 2022-02-25
### Changed
- Removes any external dependency on `expvar`.

## [0.0.7] - 2022-02-25
### Added
- Added missing `webserver.Stop`.

### Changed
- `Address` validation: only port is required.

## [0.0.6] - 2022-02-25
### Added
- Added mime type constants.

## [0.0.5] - 2022-02-25
### Added
- Option to set the base router.
- New Metric type.

### Changed
- Setting a feature's option enables the feature.
- `WithMetrics`, `WithHandlers` and `WithReadiness` now accepts multiple params.
- `New`'s params for handler, and metrics are validated.

## [0.0.4] - 2022-02-23
### Changed
- `ReadinessState` is now called `ReadinessDeterminer`.

## [0.0.3] - 2022-02-23
### Changed
- `ReadinessHandler` now provides a `ReadinessState` determiner.
- Readiness determination now support multiple `ReadinessState` determiners.

## [0.0.2] - 2022-02-22
### Added
- Added `GetTelemetry` to the `IServer` interface.
- Added more tests.

### Changed
- Fixed `NewBasic` options processing.

## [0.0.1] - 2022-02-22
### Added
- First release.
