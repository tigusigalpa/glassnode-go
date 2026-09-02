# GlassNode Golang SDK

![GlassNode Golang SDK](https://i.postimg.cc/853CkbFB/glassnode-golang-github.jpg)

[![CI](https://github.com/tigusigalpa/glassnode-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/glassnode-go/actions/workflows/ci.yml)
[![Tests](https://github.com/tigusigalpa/glassnode-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/glassnode-go/actions/workflows/test.yml)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)
[![CodeQL](https://github.com/tigusigalpa/glassnode-go/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/glassnode-go/actions/workflows/codeql.yml)
[![Codecov](https://codecov.io/gh/tigusigalpa/glassnode-go/graph/badge.svg)](https://codecov.io/gh/tigusigalpa/glassnode-go)
[![GitHub Release](https://img.shields.io/github/v/release/tigusigalpa/glassnode-go?style=flat-square)](https://github.com/tigusigalpa/glassnode-go/releases)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/tigusigalpa/glassnode-go)

> **Unofficial** Go SDK for the [Glassnode Basic API](https://docs.glassnode.com/basic-api/). Not affiliated with or
> endorsed by Glassnode — just a community-built tool I wish had existed when I started working with on-chain data.

A dependency-light, production-oriented Go module for the Glassnode Basic API, built with the kind of care you'd want
from a library you depend on every day.

If you've ever found yourself hand-crafting `net/http` calls to Glassnode, parsing JSON into `map[string]interface{}`,
and string-building metric paths at 2 AM to ship a dashboard — you already know it gets tedious fast. This SDK wraps all
of that into a clean, idiomatic Go package with strongly-typed structs, `context.Context` propagation, and sensible
defaults, so you can spend your time on application logic instead of HTTP plumbing.

The SDK exposes the complete documented metric surface through three layers of API, each one building on the last:

1. **Typed transport/config** — a `Client` with functional options, secure defaults, `context.Context` support, and zero
   external dependencies
2. **Metadata & generic metrics** — a `MetadataService` for runtime discovery and a `MetricsService` that works with any
   present or future metric path
3. **Ergonomic category services** — 25 typed service structs with convenience methods that map 1:1 to Glassnode's
   endpoint categories

## Who Is This For?

This SDK was written for the kind of developer who:

- Builds trading dashboards, research tools, or backtesting pipelines that need on-chain data
- Cares about dependency hygiene — no surprise transitive packages, no `go.sum` bloat
- Wants errors they can actually branch on with `errors.Is` / `errors.As` instead of grepping strings
- Runs services in production and needs concurrency-safe clients that won't fall over under load
- Likes their libraries to do the boring stuff (retries, redaction, context propagation) so they can focus on the
  interesting stuff

If that sounds like you, welcome — you're in the right place.

## Why This SDK?

- **Zero external dependencies** — uses only the Go standard library, so it won't bloat your `go.sum` or create
  transitive dependency conflicts that surface at the worst possible moment
- **Idiomatic Go** — functional options, `context.Context` propagation, strongly-typed structs, and `errors.Is` /
  `errors.As` support throughout
- **Concurrency-safe** — the `Client` is safe for concurrent use across goroutines, perfect for high-throughput services
  and fan-out workloads
- **Discover metrics at runtime** — not sure which parameters a metric accepts? Ask the metadata API before you make
  data calls and avoid wasting credits on bad requests
- **Survive rate limits gracefully** — automatic retry on 429 with server-aware backoff, so you don't have to wrap every
  call in your own retry loop
- **Keep your API key safe** — header-based auth by default, and the SDK redacts the key in all URLs and error messages
  so it never leaks into your logs
- **Only pay for what you use** — bulk endpoints with explicit asset lists put you in control of credit consumption, no
  accidental wildcard blow-ups

## Features

- **Zero external dependencies** — uses only the Go standard library, nothing else. Your `go.sum` stays as short as the
  day you started the project
- **25 category services** with typed convenience methods covering every documented endpoint category
- **Generic metric API** for every valid metric path — present or future, even ones not yet wrapped by a typed method
- **Metadata-first** — discover assets, metric paths, and parameter capabilities at runtime before making data calls
- **Bulk endpoint support** with repeated query parameters for multi-asset requests
- **Point-in-Time metrics** with `computed_at` timestamp preservation for historically accurate analysis and backtesting
- **Header authentication** (`X-Api-Key`) by default; query-string opt-in for the rare environments that genuinely need
  it
- **Automatic retry** on HTTP 429 with `x-rate-limit-reset` support and exponential backoff as a fallback
- **Exported, inspectable error types** with `errors.Is` / `errors.As` support — no string matching required
- **API key redaction** in all URLs and error messages — your key never leaks into logs, panic messages, or stack traces
- **Concurrency-safe** `Client` suitable for goroutine use without additional locking
- **Configurable HTTP transport** for testing and custom `http.RoundTripper` implementations (proxies, middleware,
  recording)
- **Functional options** — configure only what you need, sensible defaults for everything else

## Installation

```bash
go get github.com/tigusigalpa/glassnode-go
```

That's it. The module has zero external dependencies, so your `go.sum` stays clean. Just import it and start using it —
no `go mod tidy` surprises, no version conflicts to untangle.

## Quick Start

Here's a complete, runnable example — create a client, fetch BTC price data, and print it. Drop it into a `main.go` and
you're off:

```go
package main

import (
    "context"
    "fmt"
    "log"

    glassnode "github.com/tigusigalpa/glassnode-go"
)

func main() {
    // From environment variable GLASSNODE_API_KEY
    client, err := glassnode.NewClientFromEnv()
    if err != nil {
        log.Fatal(err)
    }

    // Or pass an explicit key if you prefer
    // client := glassnode.NewClient("YOUR_API_KEY")

    ctx := context.Background()

    // Fetch BTC price with 24h resolution
    price, err := client.Market.Price(ctx, &glassnode.MetricQuery{
        Asset:      "BTC",
        Resolution: glassnode.Resolution24h,
    })
    if err != nil {
        log.Fatal(err)
    }
    for _, p := range price {
        fmt.Printf("BTC price at %d: $%.2f\n", p.T, p.V)
    }
}
```

Want to explore what's available before pulling data? The metadata API is your friend — it lets you window-shop metrics
without spending credits on guesses:

```go
// List all supported assets
assets, err := client.Metadata.Assets(ctx)
if err != nil {
    log.Fatal(err)
}
for _, a := range assets {
    fmt.Printf("%s (%s)\n", a.Name, a.Symbol)
}

// Inspect a specific metric's parameters
metricInfo, err := client.Metadata.Metric(ctx, "/market/price_usd")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Metric: %s\n", metricInfo.Description)
fmt.Printf("Accepts resolution: %v\n", metricInfo.Resolutions)
```

### Common Use Cases

A few real-world patterns to get you oriented. These are intentionally small — copy, paste, and adapt.

**Building a market dashboard:**

```go
ctx := context.Background()

// Get OHLC candles for charting
ohlc, err := client.Market.PriceOHLC(ctx, &glassnode.MetricQuery{
    Asset:      "BTC",
    Resolution: glassnode.Resolution1h,
})

// Market cap and realized cap for valuation analysis
mcap, err := client.Market.MarketCap(ctx, &glassnode.MetricQuery{Asset: "BTC"})
rcap, err := client.Market.RealizedCap(ctx, &glassnode.MetricQuery{Asset: "BTC"})

// MVRV ratio — a classic cycle indicator
mvrv, err := client.Indicators.MVRV(ctx, &glassnode.MetricQuery{Asset: "BTC"})
```

**Monitoring network health:**

```go
// Active addresses — measures network usage
active, err := client.Addresses.ActiveCount(ctx, &glassnode.MetricQuery{
    Asset:      "BTC",
    Resolution: glassnode.Resolution24h,
})

// Hash rate — mining security
hashrate, err := client.Mining.HashRateMean(ctx, &glassnode.MetricQuery{Asset: "BTC"})

// Total supply — track inflation
supply, err := client.Supply.Current(ctx, &glassnode.MetricQuery{Asset: "BTC"})
```

**Concurrent fetches with goroutines:**

The `Client` is concurrency-safe, so fan-out patterns just work. Here's the textbook pattern for pulling several metrics
in parallel:

```go
var (
    price   []glassnode.TimePoint
    sopr    []glassnode.TimePoint
    active  []glassnode.TimePoint
)

var wg sync.WaitGroup
wg.Add(3)

query := &glassnode.MetricQuery{Asset: "BTC", Resolution: glassnode.Resolution24h}

go func() {
    defer wg.Done()
    price, _ = client.Market.Price(ctx, query)
}()
go func() {
    defer wg.Done()
    sopr, _ = client.Indicators.SOPR(ctx, query)
}()
go func() {
    defer wg.Done()
    active, _ = client.Addresses.ActiveCount(ctx, query)
}()

wg.Wait()
// The Client is concurrency-safe — no additional locking needed
```

> **Tip:** In production code, propagate errors back from each goroutine through a channel or shared `error` slice
> instead of discarding them with `_`. The example above keeps things short; real services should never swallow errors
> silently.

**Tracking credit usage over time:**

```go
usage, err := client.User.APIUsage(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Credits used today: %d / %d\n", usage.Used, usage.Limit)
```

## API Key Security

Your Glassnode API key is the gateway to your account and its data credits. Treat it like a password — because it
effectively is one.

- **Never hardcode** your API key in source code or commit it to version control. Secrets in git history are forever
- Use `NewClientFromEnv()` to read from the `GLASSNODE_API_KEY` environment variable — this is the recommended path for
  almost every deployment
- Or pass the key at runtime: `glassnode.NewClient(os.Getenv("GLASSNODE_API_KEY"))`
- The SDK **redacts** the API key in all URLs, error messages, and response metadata — your key never leaks into logs or
  panic messages, even when something goes wrong
- **Header mode** (`X-Api-Key`) is the default and recommended; the key is sent in an HTTP header and never appears in
  URLs, access logs, or proxy traces
- **Query-string mode** (`api_key`) is opt-in via `WithAuthMode(AuthModeQuery)` for environments where headers aren't
  supported, but it's less secure — the key ends up in URLs that get logged everywhere they pass through
- If you accidentally expose a key, rotate it immediately
  in [Glassnode Studio settings](https://studio.glassnode.com/settings/api). Don't wait, don't hope nobody noticed

## Authentication Modes

The SDK supports two authentication modes. Header mode is the default and recommended for all production usage — there's
rarely a good reason to switch:

```go
// Header mode (default, recommended)
// Sends the key as an X-Api-Key header — invisible in URLs and logs
client := glassnode.NewClient("key")

// Query-string mode (opt-in, less secure)
// Appends ?api_key=... to every request URL
client := glassnode.NewClient("key", glassnode.WithAuthMode(glassnode.AuthModeQuery))
```

## Configuration

The SDK uses the functional options pattern — configure only what you need, and everything else gets sensible defaults.
This keeps the common case a one-liner while still letting you tune behavior when you need to:

```go
client := glassnode.NewClient("key",
    glassnode.WithBaseURL("https://api.glassnode.com"),  // API base URL (rarely needs changing)
    glassnode.WithTimeout(15*time.Second),               // HTTP timeout
    glassnode.WithRetry(3, time.Second),                 // Max retries + base delay for 429 backoff
    glassnode.WithHTTPClient(customHTTPClient),          // Custom *http.Client for testing or proxies
    glassnode.WithUserAgent("my-app/1.0"),               // User-Agent header
    glassnode.WithAppID("trading-bot"),                  // Optional app identifier for Glassnode analytics
)
```

### Available Options

| Option                | Default                     | Description                                                  |
|-----------------------|-----------------------------|--------------------------------------------------------------|
| `WithBaseURL(url)`    | `https://api.glassnode.com` | API base URL                                                 |
| `WithTimeout(d)`      | `30s`                       | HTTP request timeout                                         |
| `WithRetry(n, delay)` | `3, 1s`                     | Max retry attempts on 429 + base backoff delay               |
| `WithHTTPClient(c)`   | `http.DefaultClient`        | Custom HTTP client (for testing, proxies, custom transports) |
| `WithUserAgent(ua)`   | `glassnode-go/1.0`          | User-Agent header                                            |
| `WithAppID(id)`       | `""`                        | Optional app identifier                                      |
| `WithAuthMode(mode)`  | `AuthModeHeader`            | `AuthModeHeader` or `AuthModeQuery`                          |

### A Note on Timeouts

The default 30s timeout is generous on purpose — some historical metric pulls return large payloads and take a moment.
If you're building a latency-sensitive service (say, a live trading dashboard), dial it down with `WithTimeout`. If
you're pulling years of daily data in a batch job, you may want to dial it up. The right value is the one that matches
your workload, not a magic number.

## Category Services

The 25 category services map 1:1 to Glassnode's documented endpoint categories. Each one exposes typed convenience
methods so you get autocomplete, compile-time checks, and self-documenting code:

| Service       | Client Field          | Documentation                                                               |
|---------------|-----------------------|-----------------------------------------------------------------------------|
| Addresses     | `client.Addresses`    | [Addresses](https://docs.glassnode.com/basic-api/endpoints/addresses)       |
| Bridges       | `client.Bridges`      | [Bridges](https://docs.glassnode.com/basic-api/endpoints/bridges)           |
| Blockchain    | `client.Blockchain`   | [Blockchain](https://docs.glassnode.com/basic-api/endpoints/blockchain)     |
| Breakdowns    | `client.Breakdowns`   | [Breakdowns](https://docs.glassnode.com/basic-api/endpoints/breakdowns)     |
| DeFi          | `client.DeFi`         | [DeFi](https://docs.glassnode.com/basic-api/endpoints/defi)                 |
| Derivatives   | `client.Derivatives`  | [Derivatives](https://docs.glassnode.com/basic-api/endpoints/derivatives)   |
| Distribution  | `client.Distribution` | [Distribution](https://docs.glassnode.com/basic-api/endpoints/distribution) |
| Entities      | `client.Entities`     | [Entities](https://docs.glassnode.com/basic-api/endpoints/entities)         |
| ETH 2.0       | `client.ETH2`         | [ETH 2.0](https://docs.glassnode.com/basic-api/endpoints/eth2)              |
| Fees          | `client.Fees`         | [Fees](https://docs.glassnode.com/basic-api/endpoints/fees)                 |
| Global        | `client.Global`       | [Global](https://docs.glassnode.com/basic-api/endpoints/global)             |
| Indicators    | `client.Indicators`   | [Indicators](https://docs.glassnode.com/basic-api/endpoints/indicators)     |
| Institutions  | `client.Institutions` | [Institutions](https://docs.glassnode.com/basic-api/endpoints/institutions) |
| Lightning     | `client.Lightning`    | [Lightning](https://docs.glassnode.com/basic-api/endpoints/lightning)       |
| Macro         | `client.Macro`        | [Macro](https://docs.glassnode.com/basic-api/endpoints/macro)               |
| Market        | `client.Market`       | [Market](https://docs.glassnode.com/basic-api/endpoints/market)             |
| Mempool       | `client.Mempool`      | [Mempool](https://docs.glassnode.com/basic-api/endpoints/mempool)           |
| Mining        | `client.Mining`       | [Mining](https://docs.glassnode.com/basic-api/endpoints/mining)             |
| Options       | `client.Options`      | [Options](https://docs.glassnode.com/basic-api/endpoints/options)           |
| Point-In-Time | `client.PointInTime`  | [PIT](https://docs.glassnode.com/basic-api/endpoints/pit)                   |
| Protocols     | `client.Protocols`    | [Protocols](https://docs.glassnode.com/basic-api/endpoints/protocols)       |
| Signals       | `client.Signals`      | [Signals](https://docs.glassnode.com/basic-api/endpoints/signals)           |
| Supply        | `client.Supply`       | [Supply](https://docs.glassnode.com/basic-api/endpoints/supply)             |
| Transactions  | `client.Transactions` | [Transactions](https://docs.glassnode.com/basic-api/endpoints/transactions) |
| Treasuries    | `client.Treasuries`   | [Treasuries](https://docs.glassnode.com/basic-api/endpoints/treasuries)     |

**Full endpoint coverage:** [docs/endpoint-coverage.md](docs/endpoint-coverage.md)

## Generic Metric API

The 25 category services cover every documented endpoint category, but Glassnode occasionally adds new metrics before a
typed wrapper exists. The generic `MetricsService` ensures you're never blocked waiting for an SDK release — it works
with **any** valid metric path, present or future:

```go
// Raw JSON — works with any metric path, returns []byte
// Use this when you want full control over decoding, or when the response shape is unusual
raw, err := client.Metrics.Get(ctx, "/addresses/sending_count", &glassnode.MetricQuery{
    Asset: "BTC",
})

// Typed scalar time-series — parses into []TimePoint
// Use this when the metric returns {t, v} pairs
data, err := client.Metrics.GetTimePoints(ctx, "/indicators/sopr", &glassnode.MetricQuery{
    Asset:      "BTC",
    Resolution: glassnode.Resolution24h,
})

// Typed object time-series — parses into []ObjectPoint
// Use this when the metric returns {t, o: {...}} pairs (e.g. OHLC)
ohlc, err := client.Metrics.GetObjectPoints(ctx, "/market/price_ohlc", &glassnode.MetricQuery{
    Asset:      "BTC",
    Resolution: glassnode.Resolution1h,
})
```

### MetricQuery Parameters

The `MetricQuery` struct covers all common query parameters. Only `Asset` is required — the rest are optional and only
sent when set:

```go
type MetricQuery struct {
    Asset           string   // Required — e.g. "BTC", "ETH"
    Resolution      string   // Optional — "10m", "1h", "24h", "1w", "1month"
    Since           *int64   // Optional — Unix timestamp
    Until           *int64   // Optional — Unix timestamp
    Currency        string   // Optional — for market metrics ("USD", "EUR")
    Format          string   // Optional — "Json" (default) or "CSV"
    TimestampFormat string   // Optional — "unix", "humanized", "humanized-30d"
}
```

Use `metadata/metric` to discover which parameters a specific metric accepts — not every metric honors every field, and
sending unsupported parameters can cause 400s.

### Convenience Resolution Constants

Instead of remembering string literals, use the exported constants for the common resolutions:

```go
glassnode.Resolution10m     // "10m"
glassnode.Resolution1h      // "1h"
glassnode.Resolution24h     // "24h"
glassnode.Resolution1w      // "1w"
glassnode.Resolution1month  // "1month"
```

## Bulk Metrics

Bulk endpoints let you fetch data for multiple assets in a single request, saving on rate-limit budget and round-trips.
However, **credits are still consumed per asset** — so a bulk call for 5 assets costs the same as 5 individual calls.
Bulk is about throughput, not credit savings:

```go
since := time.Now().AddDate(0, 0, -7).Unix()
resp, err := client.Metrics.GetBulk(ctx, "/market/mvrv", &glassnode.BulkQuery{
    Assets:     []string{"BTC", "ETH"},
    Since:      &since,
    Resolution: glassnode.Resolution24h,
})
```

### Bulk Best Practices

- **Always specify assets explicitly** — never use wildcards, as this can consume unexpected credits and burn through
  your quota before you notice
- **Check `bulk_supported`** via `metadata/metric` before calling bulk endpoints — not all metrics support it, and a 400
  is the polite way to find out
- **Batch in reasonable sizes** — 5–10 assets per call is a good balance between throughput and credit visibility;
  larger batches make it harder to attribute credit spikes
- **Monitor your credit usage** with `client.User.APIUsage()` to catch unexpected consumption early, before it becomes a
  bill

## Error Handling

The SDK exports sentinel error variables and a structured `APIError` type, so you can handle errors idiomatically with
`errors.Is` and `errors.As` — no string matching, no fragile `strings.Contains` checks that break when the API changes
its wording:

```go
_, err := client.Indicators.SOPR(ctx, &glassnode.MetricQuery{Asset: "BTC"})
if err != nil {
    switch {
    case errors.Is(err, glassnode.ErrUnauthorized):
        // 401 — API key is missing, invalid, or expired
        log.Println("Invalid API key — check your configuration")

    case errors.Is(err, glassnode.ErrRateLimited):
        // 429 — rate limit hit after all retries exhausted
        var apiErr *glassnode.APIError
        if errors.As(err, &apiErr) {
            log.Printf("Rate limited after %d retries, reset in %ds\n",
                apiErr.Retried, apiErr.RateLimitReset)
        }

    case errors.Is(err, glassnode.ErrBadRequest):
        // 400 — invalid parameters, unsupported asset, etc.
        log.Printf("Bad request: %v", err)

    case errors.Is(err, glassnode.ErrNotFound):
        // 404 — the metric path doesn't exist
        log.Println("Metric not found — check the path with metadata/metric")

    default:
        // Network errors, context cancellation, unexpected responses
        log.Printf("Unexpected error: %v", err)
    }
}
```

### Error Types

The SDK defines sentinel errors for each HTTP status code and an `APIError` struct that carries additional context —
retry count, rate-limit reset, response body — so you can make informed decisions about what to do next:

| Sentinel            | HTTP Status | Description                              |
|---------------------|-------------|------------------------------------------|
| `ErrBadRequest`     | 400         | Invalid parameters or unsupported asset  |
| `ErrUnauthorized`   | 401         | API key missing, invalid, or expired     |
| `ErrNotFound`       | 404         | Metric path not found                    |
| `ErrRateLimited`    | 429         | Rate limit hit (after retries exhausted) |
| `ErrInternalServer` | 500         | Glassnode server error                   |

The `APIError` struct includes the HTTP status code, response body, retry count, and `RateLimitReset` value (when
available). Use `errors.As` to unwrap it and inspect the details.

### A Practical Recovery Strategy

For long-running services, a simple but effective pattern is: on `ErrRateLimited`, pause the worker for `RateLimitReset`
seconds (or a sensible fallback) before retrying the queue; on `ErrUnauthorized`, fail fast and alert — there's no point
retrying a bad key. On `ErrInternalServer`, retry with backoff. Everything else is probably a bug in your request — log
it and move on.

## Retry Behavior

Nobody likes getting rate-limited, and nobody likes writing retry loops around every API call. The SDK handles 429
responses automatically so you can keep your call sites clean:

- **Retries only idempotent `GET` requests** on HTTP 429 — non-idempotent methods are never retried, because retrying
  them could double-charge credits
- **Honors `x-rate-limit-reset`** header when present — waits exactly as long as the server tells us to, no guessing
- **Falls back to exponential backoff** when the header is absent: `baseDelay * 2^attempt`
- **Never retries** 400, 401, or 404 — these are client errors that won't resolve by retrying, and retrying them just
  wastes time
- **Configurable** via `WithRetry(maxAttempts, baseDelay)` (defaults: 3 attempts, 1s base delay)

If all retry attempts are exhausted, an `APIError` wrapping `ErrRateLimited` is returned with the retry count and
`RateLimitReset` value (if available) so your application can decide how to handle it — queue the request for later,
alert the user, or back off further. The SDK does the retrying; you decide the policy on top.

## Rate Limits

Rate limits are governed by Glassnode's servers and depend on your subscription tier. The API returns these headers on
every response, so you can monitor your usage proactively rather than discovering the limit by hitting it:

| Header                   | Description                                                 |
|--------------------------|-------------------------------------------------------------|
| `x-rate-limit-limit`     | Total request limit per minute (e.g. 600 for standard tier) |
| `x-rate-limit-remaining` | Requests remaining in the current window                    |
| `x-rate-limit-reset`     | Seconds until the limit resets                              |

Metadata endpoints are separately limited to 1200 req/min, so you can discover metrics freely without worrying about
impacting your data request budget. That separation is by design — exploration shouldn't cost you throughput on the
calls that matter.

### Tips for Staying Within Limits

- **Cache metadata responses** — assets and metric definitions rarely change; there's no reason to re-fetch them on
  every request
- **Use appropriate resolutions** — don't fetch 10-minute data when you only need daily aggregates; smaller resolutions
  mean more data points and more frequent refresh needs
- **Batch with bulk endpoints where supported** — one HTTP call instead of many reduces your request count without
  changing your credit cost
- **Monitor `x-rate-limit-remaining`** and back off before hitting zero, not after — proactive throttling is cheaper
  than reactive retrying

## Data Credits

Glassnode charges data credits per request, not per data point. Understanding the credit model helps you avoid surprises
on your bill:

- **BTC**: 1 credit per request
- **All other assets**: 2 credits per request
- **Bulk endpoints**: credits = sum of individual calls (e.g., 5 assets = 5× credits)
- **Monitor usage** via `client.User.APIUsage()` or [Studio settings](https://studio.glassnode.com/settings/api)

If you're building a service that pulls many metrics, consider caching results and refreshing on a schedule rather than
polling continuously. A daily refresh job will almost always cost less than a live polling loop, and for most on-chain
analysis the difference in freshness is irrelevant.

## Point-in-Time Data

For backtesting and historically accurate analysis, the Point-in-Time service preserves the `computed_at` timestamp —
the moment a metric was actually calculated, not just the data point's timestamp. This matters because Glassnode
sometimes revises historical data, and a backtest that uses today's values for past dates will lie to you about how your
strategy would have performed:

```go
pit, err := client.PointInTime./* metric method */(ctx, &glassnode.MetricQuery{
    Asset: "BTC",
})
// Each point carries both the data timestamp (t) and the computed_at timestamp,
// so you can reconstruct exactly what was known at any moment in the past.
```

See the [Point-in-Time endpoint docs](https://docs.glassnode.com/basic-api/endpoints/pit) for the full list of supported
metrics.

## Testing

The test suite uses mocked HTTP transports — no API key or live requests are required, so you can run the full suite
anywhere: CI, air-gapped laptops, your phone if you really wanted to:

```bash
# Run all tests with verbose output
go test ./... -v

# Check formatting (no output = clean)
gofmt -d .

# Run the linter
go vet ./...
```

The suite includes 39 tests covering all services, error handling, retry logic, configuration, and the generic metrics
API. If you're contributing, please add tests for any new functionality — all tests must pass with mocked transports (no
live API calls), so the suite stays hermetic and deterministic.

### Writing Tests Against the SDK

For your own application tests, inject a custom `*http.Client` via `WithHTTPClient` and back it with an
`http.RoundTripper` that returns canned responses. This keeps your tests fast, deterministic, and free of credit
consumption:

```go
rt := &mockRoundTripper{response: cannedResponse}
client := glassnode.NewClient("test-key", glassnode.WithHTTPClient(&http.Client{Transport: rt}))
```

## Examples

The `examples/` directory contains ready-to-run programs you can adapt for your own projects. Each one is self-contained
and prints useful output so you can verify it works before wiring it into anything serious:

- [Basic usage](examples/basic/main.go) — price, indicators, assets, API usage
- [Metadata](examples/metadata/main.go) — list metrics, inspect parameters at runtime
- [Bulk metrics](examples/bulk/main.go) — multi-asset bulk requests with credit awareness
- [Error handling](examples/error_handling/main.go) — error types, `errors.Is`/`errors.As`, and recovery strategies

## Compatibility

- **Go 1.21+** (uses `log/slog`, enhanced `errors` support, and modern stdlib features)
- **No external dependencies** — uses only the Go standard library, so it builds cleanly in any environment that has Go
- Works with custom `http.RoundTripper` implementations for testing, proxies, or middleware — nothing in the SDK assumes
  a particular transport

## FAQ

<details>
<summary><strong>Do I need a Glassnode account to use this SDK?</strong></summary>

Yes. You need a Glassnode account with an API key. Sign up at [studio.glassnode.com](https://studio.glassnode.com/) —
there's a free tier with limited credits to get you started, which is plenty for experimentation.
</details>

<details>
<summary><strong>Is this an official Glassnode product?</strong></summary>

No. This is an unofficial, community-built SDK. It's not affiliated with or endorsed by Glassnode. The official API
documentation is at [docs.glassnode.com](https://docs.glassnode.com/) and should always be your source of truth for
endpoint behavior.
</details>

<details>
<summary><strong>Why zero dependencies? Isn't that overkill?</strong></summary>

Not really. The Go standard library already provides everything needed — `net/http` for transport, `encoding/json` for
parsing, `errors` for error handling. Avoiding external dependencies means no transitive dependency conflicts, no
security advisories from third-party packages to track, and a smaller binary. It's a feature, not a limitation, and it
means the SDK will keep building cleanly for years without dependency maintenance.
</details>

<details>
<summary><strong>What happens when Glassnode adds new metrics?</strong></summary>

The generic `MetricsService` works with any valid metric path, so you can use new metrics immediately even before a
typed wrapper is added. Check `metadata/metrics` to discover new paths at runtime. If you'd like a typed wrapper for a
new metric, open an issue or submit a PR — they're straightforward to add.
</details>

<details>
<summary><strong>Is the Client safe for concurrent use?</strong></summary>

Yes. The `Client` is safe for concurrent use across goroutines without additional locking. The underlying `http.Client`
is also goroutine-safe. Just pass a `context.Context` with a timeout or cancellation signal per request, and you're good
to fan out as wide as your rate limit allows.
</details>

<details>
<summary><strong>How do I get CSV format instead of JSON?</strong></summary>

Set `Format: "CSV"` in the `MetricQuery` struct. The SDK returns the raw CSV bytes as-is — it doesn't parse them,
because CSV shapes vary too much across metrics to model cleanly. JSON (the default) is automatically decoded into typed
structs.
</details>

<details>
<summary><strong>Does the SDK work with the Glassnode Advanced API?</strong></summary>

No. This SDK targets the [Basic API](https://docs.glassnode.com/basic-api/) only. The Advanced API has a different
surface and auth model; supporting it would be a separate effort. If there's enough interest, that could happen — open
an issue to signal demand.
</details>

<details>
<summary><strong>Can I use this in a commercial product?</strong></summary>

Yes. The SDK is MIT-licensed — do what you want with it, including using it in closed-source commercial products. Just
don't blame me if something breaks, and keep the license notice around as the license requires.
</details>

## Contributing

Contributions are welcome, and they're a big part of what keeps community projects like this alive. Whether it's a bug
fix, a new example, improved documentation, or a feature — here's how to get started:

1. **Fork** the repository and create your branch from `main`
2. **Run tests** to make sure everything passes before you start: `go test ./...`
3. **Make your changes** — keep code style consistent with the existing codebase (run `gofmt` and `go vet`)
4. **Add tests** for any new functionality — all tests must pass with mocked transports (no live API calls)
5. **Submit a pull request** with a clear description of what and why. The "why" matters more than the "what"

### Reporting Issues

Found a bug or have a feature request? Please open an issue on GitHub with:

- A clear description of the problem or request
- Steps to reproduce (for bugs) — ideally a minimal code snippet
- Expected vs. actual behavior
- Go version and any relevant environment details

The more reproducible your report, the faster it gets fixed.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and breaking changes. Breaking changes will be signaled clearly in
the changelog and bumped in the major version number — no silent surprises.

## Author

**Igor
Sazonov** — [sovletig@gmail.com](mailto:sovletig@gmail.com) — [github.com/tigusigalpa](https://github.com/tigusigalpa)

Bug reports, feature requests, and pull requests are all welcome.

## License

[MIT](LICENSE) — do whatever you want, just don't blame me if something breaks.

## Official Links

- [Glassnode Documentation](https://docs.glassnode.com/)
- [Glassnode Basic API](https://docs.glassnode.com/basic-api/)
- [Glassnode Studio](https://studio.glassnode.com/)
- [Glassnode Academy](https://academy.glassnode.com/)
