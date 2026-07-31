package glassnode

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// Resolution represents a supported time interval for Glassnode
// metrics. The official documented resolutions include 10m, 1h, 24h,
// 1w, and 1month, subject to metric and subscription-plan availability.
type Resolution string

const (
	Resolution10m    Resolution = "10m"
	Resolution1h     Resolution = "1h"
	Resolution24h    Resolution = "24h"
	Resolution1w     Resolution = "1w"
	Resolution1month Resolution = "1month"
)

// Currency represents a supported currency for metric values.
type Currency string

const (
	CurrencyNative Currency = "native"
	CurrencyUSD    Currency = "usd"
)

// Format represents a supported response format.
type Format string

const (
	FormatJSON Format = "json"
	FormatCSV  Format = "csv"
)

// TimestampFormat represents the timestamp format in responses.
type TimestampFormat string

const (
	TimestampUnix      TimestampFormat = "unix"
	TimestampHumanized TimestampFormat = "humanized"
)

// MetricQuery holds the common query parameters supported by Glassnode
// metric endpoints. Not all parameters apply to every metric — use the
// metadata endpoint to discover which parameters are valid for a given
// metric path.
type MetricQuery struct {
	// Asset (a) — asset id, e.g. BTC, ETH.
	Asset string
	// Since (s) — unix timestamp for the start of the data range.
	Since *int64
	// Until (u) — unix timestamp for the end of the data range.
	Until *int64
	// Resolution (i) — time interval, e.g. 24h, 1h.
	Resolution Resolution
	// Currency (c) — currency for values, e.g. native, usd.
	Currency Currency
	// Format (f) — response format, e.g. json, csv.
	Format Format
	// TimestampFormat — unix or humanized (RFC 3339).
	TimestampFormat TimestampFormat
	// Exchange (e) — exchange name, e.g. binance, coinbase.
	Exchange string
	// FromExchange — source exchange for inter-exchange metrics.
	FromExchange string
	// ToExchange — destination exchange for inter-exchange metrics.
	ToExchange string
	// Miner — miner identifier for mining-related metrics.
	Miner string
	// Maturity — maturity period for derivatives metrics.
	Maturity string
	// Network — network/blockchain identifier for cross-chain metrics.
	Network string
	// Period — time period for aggregation.
	Period string
	// QuoteSymbol — quote currency symbol for trading pairs.
	QuoteSymbol string
	// ExtraParams holds endpoint-specific parameters not covered by the
	// typed fields above. Values are URL-encoded and repeated keys are
	// supported by providing multiple values in the slice.
	ExtraParams map[string][]string
}

// ToValues converts the MetricQuery into url.Values suitable for use
// with the HTTP transport. Zero/empty values are omitted.
func (q *MetricQuery) ToValues() url.Values {
	v := url.Values{}
	if q == nil {
		return v
	}
	if q.Asset != "" {
		v.Set("a", q.Asset)
	}
	if q.Since != nil {
		v.Set("s", strconv.FormatInt(*q.Since, 10))
	}
	if q.Until != nil {
		v.Set("u", strconv.FormatInt(*q.Until, 10))
	}
	if q.Resolution != "" {
		v.Set("i", string(q.Resolution))
	}
	if q.Currency != "" {
		v.Set("c", string(q.Currency))
	}
	if q.Format != "" {
		v.Set("f", string(q.Format))
	}
	if q.TimestampFormat != "" {
		v.Set("timestamp_format", string(q.TimestampFormat))
	}
	if q.Exchange != "" {
		v.Set("e", q.Exchange)
	}
	if q.FromExchange != "" {
		v.Set("from_exchange", q.FromExchange)
	}
	if q.ToExchange != "" {
		v.Set("to_exchange", q.ToExchange)
	}
	if q.Miner != "" {
		v.Set("miner", q.Miner)
	}
	if q.Maturity != "" {
		v.Set("maturity", q.Maturity)
	}
	if q.Network != "" {
		v.Set("network", q.Network)
	}
	if q.Period != "" {
		v.Set("period", q.Period)
	}
	if q.QuoteSymbol != "" {
		v.Set("quote_symbol", q.QuoteSymbol)
	}
	for key, vals := range q.ExtraParams {
		for _, val := range vals {
			v.Add(key, val)
		}
	}
	return v
}

// BulkQuery holds query parameters for bulk metric requests. Bulk
// endpoints support repeated parameter values for assets, exchanges,
// and networks. The since (s), until (u), resolution (i), currency (c),
// and format (f) parameters behave the same as in MetricQuery.
type BulkQuery struct {
	// Assets (a) — list of asset ids. Use "*" for all available assets.
	// Always specify assets explicitly to control credit consumption.
	Assets []string
	// Exchanges (e) — list of exchange names.
	Exchanges []string
	// Networks (network) — list of network/blockchain identifiers.
	Networks []string
	// Since (s) — unix timestamp for the start of the data range.
	// Required for bulk requests.
	Since *int64
	// Until (u) — unix timestamp for the end of the data range.
	Until *int64
	// Resolution (i) — time interval. Defaults to 24h.
	Resolution Resolution
	// Currency (c) — currency for values. Defaults to native.
	Currency Currency
	// Format (f) — response format. Only json is supported for bulk.
	Format Format
	// ExtraParams holds endpoint-specific parameters not covered above.
	ExtraParams map[string][]string
}

// ToValues converts the BulkQuery into url.Values with repeated keys
// for multi-value parameters.
func (q *BulkQuery) ToValues() url.Values {
	v := url.Values{}
	if q == nil {
		return v
	}
	for _, a := range q.Assets {
		v.Add("a", a)
	}
	for _, e := range q.Exchanges {
		v.Add("e", e)
	}
	for _, n := range q.Networks {
		v.Add("network", n)
	}
	if q.Since != nil {
		v.Set("s", strconv.FormatInt(*q.Since, 10))
	}
	if q.Until != nil {
		v.Set("u", strconv.FormatInt(*q.Until, 10))
	}
	if q.Resolution != "" {
		v.Set("i", string(q.Resolution))
	}
	if q.Currency != "" {
		v.Set("c", string(q.Currency))
	}
	if q.Format != "" {
		v.Set("f", string(q.Format))
	}
	for key, vals := range q.ExtraParams {
		for _, val := range vals {
			v.Add(key, val)
		}
	}
	return v
}

// TimePoint represents a standard scalar time-series data point:
// {"t": <unix-timestamp>, "v": <value>}.
type TimePoint struct {
	T int64   `json:"t"`
	V float64 `json:"v"`
}

// ObjectPoint represents an object-valued time-series data point:
// {"t": <unix-timestamp>, "o": {...}}.
type ObjectPoint struct {
	T int64           `json:"t"`
	O json.RawMessage `json:"o"`
}

// PITTimePoint represents a point-in-time scalar data point that
// includes a computed_at timestamp:
// {"t": <unix-timestamp>, "v": <value>, "computed_at": <unix-timestamp>}.
type PITTimePoint struct {
	T          int64   `json:"t"`
	V          float64 `json:"v"`
	ComputedAt int64   `json:"computed_at,omitempty"`
}

// PITObjectPoint represents a point-in-time object-valued data point:
// {"t": <unix-timestamp>, "o": {...}, "computed_at": <unix-timestamp>}.
type PITObjectPoint struct {
	T          int64           `json:"t"`
	O          json.RawMessage `json:"o"`
	ComputedAt int64           `json:"computed_at,omitempty"`
}

// BulkPoint represents a single entry in a bulk metric response. The
// parameter fields and value/category depend on the metric type.
type BulkPoint struct {
	T        int64           `json:"t"`
	Bulk     json.RawMessage `json:"bulk"`
	Computed json.RawMessage `json:"computed,omitempty"`
}

// BulkResponse wraps the bulk metric response envelope:
// {"data": [...]}.
type BulkResponse struct {
	Data []BulkPoint `json:"data"`
}

// RawResponse wraps a raw JSON response for endpoints whose response
// shape cannot be safely represented with a stable struct. The Raw
// field contains the original decoded JSON.
type RawResponse struct {
	Raw json.RawMessage `json:"-"`
}

// AssetBlockchain represents a blockchain entry in asset metadata.
type AssetBlockchain struct {
	Blockchain     string `json:"blockchain"`
	Address        string `json:"address"`
	Decimals       int    `json:"decimals"`
	OnChainSupport bool   `json:"on_chain_support"`
}

// AssetExternalIDs holds references to external third-party IDs.
type AssetExternalIDs struct {
	CCData        string `json:"ccdata"`
	CoinMarketCap string `json:"coinmarketcap"`
	CoinGecko     string `json:"coingecko"`
}

// Asset represents a single asset in the metadata/assets response.
type Asset struct {
	ID           string            `json:"id"`
	ExternalIDs  AssetExternalIDs  `json:"external_ids"`
	Symbol       string            `json:"symbol"`
	Name         string            `json:"name"`
	AssetType    string            `json:"asset_type"`
	Blockchains  []AssetBlockchain `json:"blockchains"`
	LogoURL      string            `json:"logo_url"`
	SemanticTags []string          `json:"semantic_tags"`
}

// AssetsResponse wraps the metadata/assets response: {"data": [...]}.
type AssetsResponse struct {
	Data []Asset `json:"data"`
}

// MetricVariant holds references to related metric variants.
type MetricVariant struct {
	Base string `json:"base"`
	Bulk string `json:"bulk"`
	PIT  string `json:"pit"`
}

// MetricRefs holds links to the metric in Studio, docs, and variants.
type MetricRefs struct {
	Docs          string        `json:"docs"`
	Studio        string        `json:"studio"`
	MetricVariant MetricVariant `json:"metric_variant"`
}

// MetricDescriptors holds detailed information about a metric.
type MetricDescriptors struct {
	Name        string            `json:"name"`
	ShortName   string            `json:"short_name"`
	Group       string            `json:"group"`
	Tags        []string          `json:"tags"`
	Description map[string]string `json:"description"`
}

// Timerange holds the minimum and maximum available timestamps.
type Timerange struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

// MetricMetadata represents the response from metadata/metric.
type MetricMetadata struct {
	Path        string              `json:"path"`
	Tier        int                 `json:"tier"`
	IsPIT       bool                `json:"is_pit"`
	Parameters  map[string][]string `json:"parameters"`
	Queried     map[string]string   `json:"queried"`
	Timerange   *Timerange          `json:"timerange"`
	Modified    int64               `json:"modified"`
	Refs        *MetricRefs         `json:"refs"`
	Descriptors *MetricDescriptors  `json:"descriptors"`
}

// APIUsage represents the response from the user/api_usage endpoint.
type APIUsage struct {
	CreditsUsed int `json:"creditsUsed"`
}

// normalizeMetricPath ensures the metric path starts with "/" and does
// not include the "/v1/metrics" prefix (which is added by the caller).
func normalizeMetricPath(path string) string {
	if path == "" {
		return ""
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}

// metricsPath builds the full API path for a metric endpoint.
func metricsPath(metricPath string) string {
	return "/v1/metrics" + normalizeMetricPath(metricPath)
}

// bulkMetricsPath builds the full API path for a bulk metric endpoint.
func bulkMetricsPath(metricPath string) string {
	return "/v1/metrics" + normalizeMetricPath(metricPath) + "/bulk"
}
