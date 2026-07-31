package glassnode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ResponseMetadata holds non-body information from an HTTP response. It
// is returned alongside the decoded payload so applications can inspect
// rate-limit headers, request IDs, and timing.
type ResponseMetadata struct {
	// StatusCode is the HTTP status code.
	StatusCode int
	// RateLimitLimit is the value of the x-rate-limit-limit header.
	RateLimitLimit string
	// RateLimitRemaining is the value of the x-rate-limit-remaining
	// header.
	RateLimitRemaining string
	// RateLimitReset is the value of the x-rate-limit-reset header
	// (seconds until the limit resets).
	RateLimitReset string
	// RequestID is the value of the X-Request-ID header, if present.
	RequestID string
	// Duration is the time taken for the HTTP round-trip.
	Duration time.Duration
	// FinalURL is the redacted final URL of the request. If query-string
	// authentication is used, the api_key value is replaced with
	// "[REDACTED]".
	FinalURL string
}

// get performs an HTTP GET request against the Glassnode API and decodes
// the JSON response into out. It handles authentication, query string
// construction, error mapping, response metadata capture, and automatic
// retry with exponential backoff for HTTP 429 responses.
//
// path must be an absolute path beginning with "/" relative to the
// client's configured base URL, e.g. "/v1/metrics/indicators/sopr".
func (c *Client) get(ctx context.Context, path string, query url.Values, out interface{}) (*ResponseMetadata, error) {
	return c.doGet(ctx, path, query, out, 0)
}

// doGet is the internal retry loop for GET requests.
func (c *Client) doGet(ctx context.Context, path string, query url.Values, out interface{}, retried int) (*ResponseMetadata, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	fullURL := c.buildURL(path, query)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("glassnode: failed to build request: %w", err)
	}

	// Apply authentication.
	if c.authMode == AuthModeHeader {
		req.Header.Set("X-Api-Key", c.apiKey)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.buildUserAgent())

	start := time.Now()
	resp, err := c.httpClient.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		return nil, fmt.Errorf("glassnode: request failed: %w", err)
	}
	defer resp.Body.Close()

	meta := c.captureMetadata(resp, fullURL, elapsed)

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return meta, fmt.Errorf("glassnode: failed to read response body: %w", readErr)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if out != nil && len(body) > 0 {
			if err := json.Unmarshal(body, out); err != nil {
				return meta, fmt.Errorf("glassnode: failed to decode response: %w", err)
			}
		}
		return meta, nil
	}

	// Retry on 429 if attempts remain.
	if resp.StatusCode == http.StatusTooManyRequests && retried+1 < c.maxAttempts {
		delay := c.retryDelay(meta.RateLimitReset, retried)
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return meta, ctx.Err()
		case <-timer.C:
		}
		return c.doGet(ctx, path, query, out, retried+1)
	}

	return meta, newAPIError(resp.StatusCode, extractMessage(body), body, meta, retried)
}

// buildURL constructs the full request URL with query parameters.
// When AuthModeQuery is used, the api_key parameter is appended.
func (c *Client) buildURL(path string, query url.Values) string {
	fullURL := c.baseURL + path

	// Copy query values so we don't mutate the caller's map.
	q := url.Values{}
	for k, v := range query {
		q[k] = v
	}

	if c.authMode == AuthModeQuery {
		q.Set("api_key", c.apiKey)
	}

	if len(q) > 0 {
		fullURL += "?" + q.Encode()
	}

	return fullURL
}

// buildUserAgent returns the User-Agent header value, optionally
// appending the application identifier.
func (c *Client) buildUserAgent() string {
	ua := c.userAgent
	if c.appID != "" {
		ua += " " + c.appID
	}
	return ua
}

// captureMetadata builds a ResponseMetadata from the HTTP response.
func (c *Client) captureMetadata(resp *http.Response, fullURL string, elapsed time.Duration) *ResponseMetadata {
	meta := &ResponseMetadata{
		StatusCode:         resp.StatusCode,
		RateLimitLimit:     resp.Header.Get("x-rate-limit-limit"),
		RateLimitRemaining: resp.Header.Get("x-rate-limit-remaining"),
		RateLimitReset:     resp.Header.Get("x-rate-limit-reset"),
		RequestID:          resp.Header.Get("X-Request-ID"),
		Duration:           elapsed,
		FinalURL:           redactURL(fullURL),
	}
	return meta
}

// retryDelay computes the delay before the next retry attempt. It
// honors the x-rate-limit-reset header (seconds until reset) if present
// and valid, otherwise falls back to exponential backoff with jitter.
func (c *Client) retryDelay(rateLimitReset string, attempt int) time.Duration {
	if rateLimitReset != "" {
		if secs, err := strconv.Atoi(strings.TrimSpace(rateLimitReset)); err == nil && secs >= 0 {
			return time.Duration(secs) * time.Second
		}
	}
	baseDelay := c.baseDelay
	if baseDelay <= 0 {
		baseDelay = DefaultBaseDelay
	}
	multiplier := 1 << attempt
	return baseDelay * time.Duration(multiplier)
}

// redactURL replaces the api_key query parameter value with
// "[REDACTED]" to prevent credential leakage in logs and error
// messages.
func redactURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	q := u.Query()
	if q.Has("api_key") {
		q.Set("api_key", "[REDACTED]")
		u.RawQuery = q.Encode()
	}
	return u.String()
}

// extractMessage attempts to pull a human-readable message from a JSON
// error response body. If the body cannot be parsed, a generic message
// with a truncated snippet is returned.
func extractMessage(body []byte) string {
	var parsed struct {
		Message string `json:"message"`
		Msg     string `json:"msg"`
		Error   string `json:"error"`
	}
	if err := json.Unmarshal(body, &parsed); err == nil {
		if parsed.Message != "" {
			return parsed.Message
		}
		if parsed.Msg != "" {
			return parsed.Msg
		}
		if parsed.Error != "" {
			return parsed.Error
		}
	}
	if len(body) == 0 {
		return "unknown error"
	}
	snippet := string(body)
	if len(snippet) > 200 {
		snippet = snippet[:200]
	}
	return snippet
}
