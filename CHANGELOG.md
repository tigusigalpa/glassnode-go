# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-01-15

### Added
- Initial release of `glassnode-go`.
- Core `Client` with functional options: `WithAPIKey`, `WithAuthMode`, `WithBaseURL`, `WithHTTPClient`, `WithTimeout`, `WithRetry`, `WithUserAgent`, `WithAppID`.
- `NewClientFromEnv` constructor reading `GLASSNODE_API_KEY`.
- Header (`X-Api-Key`) and query-string (`api_key`) authentication modes.
- `MetadataService` — `Assets`, `Metrics`, `Metric` endpoints.
- `UserService` — `APIUsage` endpoint.
- `MetricsService` — generic `Get`, `GetTimePoints`, `GetObjectPoints`, `GetBulk`, `GetBulkRaw`.
- 25 category services with typed convenience methods: Addresses, Bridges, Blockchain, Breakdowns, DeFi, Derivatives, Distribution, Entities, ETH2, Fees, Global, Indicators, Institutions, Lightning, Macro, Market, Mempool, Mining, Options, PointInTime, Protocols, Signals, Supply, Transactions, Treasuries.
- Exported error types: `APIError`, `ErrBadRequest`, `ErrUnauthorized`, `ErrNotFound`, `ErrRateLimited`, `ErrMissingAPIKey`, `ErrInvalidMetricPath`.
- Automatic retry on HTTP 429 with `x-rate-limit-reset` header support and exponential backoff.
- `ResponseMetadata` with rate-limit headers, request ID, duration, and redacted URL.
- `MetricQuery` and `BulkQuery` parameter models with `ToValues()`.
- Typed response models: `TimePoint`, `ObjectPoint`, `PITTimePoint`, `PITObjectPoint`, `BulkPoint`, `BulkResponse`, `Asset`, `MetricMetadata`, `APIUsage`.
- Comprehensive test suite (39 tests) covering authentication, retries, errors, URL encoding, bulk params, context cancellation, concurrency, and more.
- `docs/endpoint-coverage.md` documenting all covered endpoints.
