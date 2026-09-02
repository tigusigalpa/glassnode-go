package glassnode

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"
)

// newTestClient spins up an httptest.Server using handler and returns a
// *Client configured to talk to it, along with the server for cleanup.
func newTestClient(t *testing.T, handler http.HandlerFunc, opts ...Option) (*Client, *httptest.Server) {
	t.Helper()
	srv := httptest.NewServer(handler)
	allOpts := append([]Option{WithBaseURL(srv.URL)}, opts...)
	c := NewClient("test-api-key", allOpts...)
	return c, srv
}

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient("key")
	if c.baseURL != DefaultBaseURL {
		t.Errorf("expected baseURL %s, got %s", c.baseURL, DefaultBaseURL)
	}
	if c.httpClient.Timeout != DefaultTimeout {
		t.Errorf("expected timeout %s, got %s", c.httpClient.Timeout, DefaultTimeout)
	}
	if c.maxAttempts != DefaultMaxAttempts {
		t.Errorf("expected maxAttempts %d, got %d", c.maxAttempts, DefaultMaxAttempts)
	}
	if c.authMode != AuthModeHeader {
		t.Errorf("expected AuthModeHeader, got %d", c.authMode)
	}
	if c.Metadata == nil || c.User == nil || c.Metrics == nil {
		t.Error("expected core services to be initialized")
	}
	if c.Addresses == nil || c.Market == nil || c.Indicators == nil {
		t.Error("expected category services to be initialized")
	}
}

func TestNewClient_Options(t *testing.T) {
	customClient := &http.Client{}
	c := NewClient("key",
		WithBaseURL("https://example.com"),
		WithHTTPClient(customClient),
		WithTimeout(5*time.Second),
		WithRetry(3, 100*time.Millisecond),
		WithAuthMode(AuthModeQuery),
		WithUserAgent("custom-ua"),
		WithAppID("my-app"),
	)
	if c.baseURL != "https://example.com" {
		t.Errorf("expected custom baseURL, got %s", c.baseURL)
	}
	if c.httpClient != customClient {
		t.Error("expected custom http client to be used")
	}
	if c.httpClient.Timeout != 5*time.Second {
		t.Errorf("expected timeout 5s, got %s", c.httpClient.Timeout)
	}
	if c.maxAttempts != 3 {
		t.Errorf("expected maxAttempts 3, got %d", c.maxAttempts)
	}
	if c.baseDelay != 100*time.Millisecond {
		t.Errorf("expected baseDelay 100ms, got %s", c.baseDelay)
	}
	if c.authMode != AuthModeQuery {
		t.Errorf("expected AuthModeQuery, got %d", c.authMode)
	}
	if c.userAgent != "custom-ua" {
		t.Errorf("expected userAgent 'custom-ua', got %s", c.userAgent)
	}
	if c.appID != "my-app" {
		t.Errorf("expected appID 'my-app', got %s", c.appID)
	}
}

func TestGet_Retry429(t *testing.T) {
	var attempts int
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.Header().Set("x-rate-limit-reset", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message":"rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":822077}]`))
	}, WithRetry(5, 10*time.Millisecond))
	defer srv.Close()

	data, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("expected no error after retries, got %v", err)
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
	if len(data) != 1 || data[0].V != 822077 {
		t.Errorf("unexpected data: %v", data)
	}
}

func TestGet_RateLimitedExhausted(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"message":"rate limited"}`))
	}, WithRetry(2, 5*time.Millisecond))
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrRateLimited) {
		t.Errorf("expected ErrRateLimited, got %v", err)
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Errorf("expected *APIError, got %T", err)
	} else if apiErr.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", apiErr.StatusCode)
	}
}

func TestGet_Unauthorized(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"invalid api key"}`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if !errors.Is(err, ErrUnauthorized) {
		t.Errorf("expected ErrUnauthorized, got %v", err)
	}
}

func TestGet_NotFound(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"not found"}`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestGet_BadRequest(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"invalid parameters"}`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if !errors.Is(err, ErrBadRequest) {
		t.Errorf("expected ErrBadRequest, got %v", err)
	}
}

func TestGet_ContextCancellation(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	_, err := c.Metrics.GetTimePoints(ctx, "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err == nil {
		t.Fatal("expected context deadline error, got nil")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected context.DeadlineExceeded, got %v", err)
	}
}

func TestGet_AuthorizationHeader(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Api-Key"); got != "test-api-key" {
			t.Errorf("expected X-Api-Key header 'test-api-key', got %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGet_AuthModeQuery(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("api_key"); got != "test-api-key" {
			t.Errorf("expected api_key query param 'test-api-key', got %q", got)
		}
		if got := r.Header.Get("X-Api-Key"); got != "" {
			t.Errorf("expected no X-Api-Key header, got %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}, WithAuthMode(AuthModeQuery))
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGet_URLEncoding(t *testing.T) {
	var capturedPath string
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedPath != "/v1/metrics/indicators/sopr" {
		t.Errorf("expected path /v1/metrics/indicators/sopr, got %s", capturedPath)
	}
}

func TestGet_QueryParams(t *testing.T) {
	var capturedQuery string
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	since := int64(1609459200)
	until := int64(1609545600)
	q := &MetricQuery{
		Asset:      "BTC",
		Since:      &since,
		Until:      &until,
		Resolution: Resolution24h,
		Currency:   CurrencyUSD,
		Format:     FormatJSON,
	}
	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", q)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedQuery == "" {
		t.Error("expected non-empty query string")
	}
}

func TestGet_BulkRepeatedParams(t *testing.T) {
	var capturedQuery string
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":[]}`))
	})
	defer srv.Close()

	since := int64(1609459200)
	bq := &BulkQuery{
		Assets:     []string{"BTC", "ETH"},
		Exchanges:  []string{"binance", "coinbase"},
		Since:      &since,
		Resolution: Resolution24h,
	}
	_, err := c.Metrics.GetBulk(context.Background(), "/market/mvrv", bq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedQuery == "" {
		t.Error("expected non-empty query string")
	}
}

func TestGet_MalformedJSON(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{invalid json`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err == nil {
		t.Fatal("expected error for malformed JSON, got nil")
	}
}

func TestGet_ObjectTimeSeries(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"o":{"high":100,"low":90}}]`))
	})
	defer srv.Close()

	data, err := c.Metrics.GetObjectPoints(context.Background(), "/market/price_ohlc", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("expected 1 data point, got %d", len(data))
	}
	var ohlc struct {
		High float64 `json:"high"`
		Low  float64 `json:"low"`
	}
	if err := json.Unmarshal(data[0].O, &ohlc); err != nil {
		t.Fatalf("failed to unmarshal object: %v", err)
	}
	if ohlc.High != 100 || ohlc.Low != 90 {
		t.Errorf("unexpected OHLC values: %+v", ohlc)
	}
}

func TestGet_RawResponse(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":123.45}]`))
	})
	defer srv.Close()

	raw, err := c.Metrics.Get(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var points []TimePoint
	if err := json.Unmarshal(raw, &points); err != nil {
		t.Fatalf("failed to unmarshal raw: %v", err)
	}
	if len(points) != 1 || points[0].V != 123.45 {
		t.Errorf("unexpected data: %v", points)
	}
}

func TestGet_RawCSVResponse(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("f"); got != string(FormatCSV) {
			t.Errorf("expected format query parameter %q, got %q", FormatCSV, got)
		}
		w.Header().Set("Content-Type", "text/csv")
		_, _ = w.Write([]byte("t,v\n1726790400,123.45\n"))
	})
	defer srv.Close()

	raw, err := c.Metrics.Get(context.Background(), "/indicators/sopr", &MetricQuery{
		Asset:  "BTC",
		Format: FormatCSV,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got, want := string(raw), "t,v\n1726790400,123.45\n"; got != want {
		t.Errorf("expected CSV response %q, got %q", want, got)
	}
}

func TestGet_RawResponseRejectsMalformedJSON(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`not json`))
	})
	defer srv.Close()

	_, err := c.Metrics.Get(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err == nil {
		t.Fatal("expected malformed JSON error")
	}
}

func TestGet_BaseURLWithTrailingSlash(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, "/v1/metrics/indicators/sopr"; got != want {
			t.Errorf("expected request path %q, got %q", want, got)
		}
		_, _ = w.Write([]byte(`[]`))
	}))
	defer srv.Close()

	c := NewClient("test-api-key", WithBaseURL(srv.URL+"/"))
	if _, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGet_RateLimitHeaders(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-rate-limit-limit", "600")
		w.Header().Set("x-rate-limit-remaining", "599")
		w.Header().Set("x-rate-limit-reset", "60")
		w.Header().Set("X-Request-ID", "req-123")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGet_EmptyMetricPath(t *testing.T) {
	c := NewClient("key")
	_, err := c.Metrics.Get(context.Background(), "", &MetricQuery{Asset: "BTC"})
	if !errors.Is(err, ErrInvalidMetricPath) {
		t.Errorf("expected ErrInvalidMetricPath, got %v", err)
	}
}

func TestAPIError_Error(t *testing.T) {
	err := &APIError{StatusCode: 500, Message: "server error"}
	want := "glassnode: api error: status=500 message=server error"
	if err.Error() != want {
		t.Errorf("expected %q, got %q", want, err.Error())
	}
}

func TestRedactURL(t *testing.T) {
	raw := "https://api.glassnode.com/v1/metrics/indicators/sopr?a=BTC&api_key=secret123"
	redacted := redactURL(raw)
	if redacted == raw {
		t.Error("expected URL to be redacted")
	}
	if contains(redacted, "secret123") {
		t.Errorf("expected API key to be redacted, got %s", redacted)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsStr(s, substr))
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestClient_ConcurrentUse(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":1}]`))
	})
	defer srv.Close()

	var wg sync.WaitGroup
	errCh := make(chan error, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
			if err != nil {
				errCh <- err
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Errorf("unexpected concurrent error: %v", err)
	}
}

func TestMetadata_Assets(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metadata/assets" {
			t.Errorf("expected path /v1/metadata/assets, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":[{"id":"BTC","symbol":"BTC","name":"Bitcoin","asset_type":"coin"}]}`))
	})
	defer srv.Close()

	out, err := c.Metadata.Assets(context.Background(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Data) != 1 || out.Data[0].ID != "BTC" {
		t.Errorf("unexpected assets: %+v", out)
	}
}

func TestMetadata_Metrics(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metadata/metrics" {
			t.Errorf("expected path /v1/metadata/metrics, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`["/addresses/count","/indicators/sopr"]`))
	})
	defer srv.Close()

	paths, err := c.Metadata.Metrics(context.Background(), &MetricQuery{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(paths) != 2 || paths[0] != "/addresses/count" {
		t.Errorf("unexpected paths: %v", paths)
	}
}

func TestMetadata_Metric(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metadata/metric" {
			t.Errorf("expected path /v1/metadata/metric, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("path") != "/indicators/sopr" {
			t.Errorf("expected path param /indicators/sopr, got %s", r.URL.Query().Get("path"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"path":"/indicators/sopr","tier":1,"is_pit":false,"parameters":{"a":["BTC","ETH"]}}`))
	})
	defer srv.Close()

	meta, err := c.Metadata.Metric(context.Background(), "/indicators/sopr", &MetricQuery{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meta.Path != "/indicators/sopr" {
		t.Errorf("unexpected path: %s", meta.Path)
	}
	if meta.Tier != 1 {
		t.Errorf("unexpected tier: %d", meta.Tier)
	}
}

func TestMetadata_Metric_EmptyPath(t *testing.T) {
	c := NewClient("key")
	_, err := c.Metadata.Metric(context.Background(), "", &MetricQuery{})
	if !errors.Is(err, ErrInvalidMetricPath) {
		t.Errorf("expected ErrInvalidMetricPath, got %v", err)
	}
}

func TestUser_APIUsage(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/user/api_usage" {
			t.Errorf("expected path /v1/user/api_usage, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"creditsUsed":42}`))
	})
	defer srv.Close()

	usage, err := c.User.APIUsage(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if usage.CreditsUsed != 42 {
		t.Errorf("expected creditsUsed 42, got %d", usage.CreditsUsed)
	}
}

func TestCategoryService_Addresses(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics/addresses/active_count" {
			t.Errorf("expected path /v1/metrics/addresses/active_count, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":1000000}]`))
	})
	defer srv.Close()

	data, err := c.Addresses.ActiveCount(context.Background(), &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 1 || data[0].V != 1000000 {
		t.Errorf("unexpected data: %v", data)
	}
}

func TestCategoryService_Market(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics/market/price_usd_close" {
			t.Errorf("expected path /v1/metrics/market/price_usd_close, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":45000.5}]`))
	})
	defer srv.Close()

	data, err := c.Market.Price(context.Background(), &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 1 || data[0].V != 45000.5 {
		t.Errorf("unexpected data: %v", data)
	}
}

func TestCategoryService_MarketOHLC(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics/market/price_usd_ohlc" {
			t.Errorf("expected path /v1/metrics/market/price_usd_ohlc, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"o":{"o":44000,"h":46000,"l":43000,"c":45000}}]`))
	})
	defer srv.Close()

	data, err := c.Market.PriceOHLC(context.Background(), &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("expected 1 data point, got %d", len(data))
	}
}

func TestMacroServicePaths(t *testing.T) {
	var paths []string
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	q := &MetricQuery{Asset: "BTC"}
	methods := []func(context.Context, *MetricQuery) ([]TimePoint, error){
		c.Macro.USInterestRate,
		c.Macro.USCentralBankBalanceSheet,
		c.Macro.USGDP,
		c.Macro.USCoreInflationRate,
		c.Macro.USMoneySupplyM2,
		c.Macro.EuroAreaInterestRate,
	}
	for _, method := range methods {
		if _, err := method(context.Background(), q); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	want := []string{
		"/v1/metrics/macro/us_interest_rate",
		"/v1/metrics/macro/us_central_bank_balance_sheet",
		"/v1/metrics/macro/us_gdp",
		"/v1/metrics/macro/us_core_inflation_rate",
		"/v1/metrics/macro/us_money_supply_m2",
		"/v1/metrics/macro/euro_interest_rate",
	}
	if !reflect.DeepEqual(paths, want) {
		t.Errorf("unexpected macro paths:\nwant %v\n got %v", want, paths)
	}
}

func TestMarketUSDServicePaths(t *testing.T) {
	var paths []string
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		_, _ = w.Write([]byte(`[]`))
	})
	defer srv.Close()

	q := &MetricQuery{Asset: "BTC"}
	methods := []func(context.Context, *MetricQuery) ([]TimePoint, error){
		c.Market.MarketCap,
		c.Market.RealizedCap,
		c.Market.RealizedPrice,
	}
	for _, method := range methods {
		if _, err := method(context.Background(), q); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	want := []string{
		"/v1/metrics/market/marketcap_usd",
		"/v1/metrics/market/marketcap_realized_usd",
		"/v1/metrics/market/price_realized_usd",
	}
	if !reflect.DeepEqual(paths, want) {
		t.Errorf("unexpected market paths:\nwant %v\n got %v", want, paths)
	}
}

func TestCategoryService_Indicators(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics/indicators/sopr" {
			t.Errorf("expected path /v1/metrics/indicators/sopr, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":1.05}]`))
	})
	defer srv.Close()

	data, err := c.Indicators.SOPR(context.Background(), &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 1 || data[0].V != 1.05 {
		t.Errorf("unexpected data: %v", data)
	}
}

func TestCategoryService_PIT(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics/indicators/sopr_pit" {
			t.Errorf("expected path /v1/metrics/indicators/sopr_pit, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"t":1726790400,"v":1.05,"computed_at":1726790500}]`))
	})
	defer srv.Close()

	data, err := c.PointInTime.GetPITTimePoints(context.Background(), "/indicators/sopr_pit", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(data) != 1 || data[0].ComputedAt != 1726790500 {
		t.Errorf("unexpected data: %v", data)
	}
}

func TestBulkMetrics(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/metrics/market/mvrv/bulk" {
			t.Errorf("expected path /v1/metrics/market/mvrv/bulk, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":[{"t":1726790400,"bulk":[{"a":"BTC","v":2.1}]}]}`))
	})
	defer srv.Close()

	since := int64(1609459200)
	bq := &BulkQuery{
		Assets:     []string{"BTC"},
		Since:      &since,
		Resolution: Resolution24h,
	}
	resp, err := c.Metrics.GetBulk(context.Background(), "/market/mvrv", bq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 data point, got %d", len(resp.Data))
	}
}

func TestNewClientFromEnv_MissingKey(t *testing.T) {
	// Ensure env var is not set
	defaultLookupEnv = func(key string) (string, bool) { return "", false }
	defer func() {
		defaultLookupEnv = func(key string) (string, bool) { return "", false }
	}()

	_, err := NewClientFromEnv()
	if !errors.Is(err, ErrMissingAPIKey) {
		t.Errorf("expected ErrMissingAPIKey, got %v", err)
	}
}

func TestNewClientFromEnv_WithKey(t *testing.T) {
	defaultLookupEnv = func(key string) (string, bool) {
		if key == "GLASSNODE_API_KEY" {
			return "env-key", true
		}
		return "", false
	}
	defer func() {
		defaultLookupEnv = func(key string) (string, bool) { return "", false }
	}()

	c, err := NewClientFromEnv()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.apiKey != "env-key" {
		t.Errorf("expected apiKey 'env-key', got %s", c.apiKey)
	}
}

func TestRetryDelay_RateLimitReset(t *testing.T) {
	c := NewClient("key", WithRetry(3, 100*time.Millisecond))
	delay := c.retryDelay("5", 0)
	if delay != 5*time.Second {
		t.Errorf("expected 5s delay from rate-limit-reset, got %s", delay)
	}
}

func TestRetryDelay_ExponentialBackoff(t *testing.T) {
	c := NewClient("key", WithRetry(3, 100*time.Millisecond))
	delay := c.retryDelay("", 0)
	if delay != 100*time.Millisecond {
		t.Errorf("expected 100ms delay, got %s", delay)
	}
	delay1 := c.retryDelay("", 1)
	if delay1 != 200*time.Millisecond {
		t.Errorf("expected 200ms delay, got %s", delay1)
	}
}

func TestMetricQuery_ToValues(t *testing.T) {
	since := int64(1609459200)
	q := &MetricQuery{
		Asset:      "BTC",
		Since:      &since,
		Resolution: Resolution1h,
		Currency:   CurrencyUSD,
		Exchange:   "binance",
	}
	v := q.ToValues()
	if v.Get("a") != "BTC" {
		t.Errorf("expected a=BTC, got %s", v.Get("a"))
	}
	if v.Get("s") != "1609459200" {
		t.Errorf("expected s=1609459200, got %s", v.Get("s"))
	}
	if v.Get("i") != "1h" {
		t.Errorf("expected i=1h, got %s", v.Get("i"))
	}
	if v.Get("c") != "usd" {
		t.Errorf("expected c=usd, got %s", v.Get("c"))
	}
	if v.Get("e") != "binance" {
		t.Errorf("expected e=binance, got %s", v.Get("e"))
	}
}

func TestBulkQuery_ToValues(t *testing.T) {
	since := int64(1609459200)
	q := &BulkQuery{
		Assets:     []string{"BTC", "ETH"},
		Exchanges:  []string{"binance"},
		Since:      &since,
		Resolution: Resolution24h,
	}
	v := q.ToValues()
	if len(v["a"]) != 2 {
		t.Errorf("expected 2 assets, got %d", len(v["a"]))
	}
	if v["a"][0] != "BTC" || v["a"][1] != "ETH" {
		t.Errorf("expected [BTC, ETH], got %v", v["a"])
	}
	if len(v["e"]) != 1 || v["e"][0] != "binance" {
		t.Errorf("expected [binance], got %v", v["e"])
	}
}

func TestGet_UserAgent(t *testing.T) {
	c, srv := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		if ua == "" {
			t.Error("expected non-empty User-Agent header")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}, WithUserAgent("test-agent"), WithAppID("test-app"))
	defer srv.Close()

	_, err := c.Metrics.GetTimePoints(context.Background(), "/indicators/sopr", &MetricQuery{Asset: "BTC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
