package glassnode

import (
	"errors"
	"fmt"
)

// APIError represents an error response returned by the Glassnode API.
// It is returned by every service method when the API responds with a
// non-2xx HTTP status code. The error carries the status code, a safe
// response excerpt, rate-limit metadata, and the request ID when
// available, but never includes the API key.
type APIError struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int
	// Message is a human-readable error message extracted from the
	// response body or a generated fallback.
	Message string
	// RawBody contains a truncated, redacted copy of the raw response
	// body for debugging purposes (max 512 bytes).
	RawBody []byte
	// RequestID is the value of the X-Request-ID response header, if
	// the server provides one.
	RequestID string
	// RateLimitLimit is the value of the x-rate-limit-limit header, if
	// present.
	RateLimitLimit string
	// RateLimitRemaining is the value of the x-rate-limit-remaining
	// header, if present.
	RateLimitRemaining string
	// RateLimitReset is the value of the x-rate-limit-reset header, if
	// present.
	RateLimitReset string
	// Retried is the number of retry attempts made before this error
	// was returned.
	Retried int
}

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	return fmt.Sprintf("glassnode: api error: status=%d message=%s", e.StatusCode, e.Message)
}

// Sentinel errors that can be checked with errors.Is against errors
// returned from service methods. The underlying *APIError is always
// available via errors.As for accessing StatusCode, RawBody, etc.
var (
	// ErrMissingAPIKey is returned when no API key is provided.
	ErrMissingAPIKey = errors.New("glassnode: missing API key — set GLASSNODE_API_KEY or pass a key to NewClient")
	// ErrBadRequest is returned when the API responds with HTTP 400.
	ErrBadRequest = errors.New("glassnode: bad request — invalid parameters or unsupported asset")
	// ErrUnauthorized is returned when the API responds with HTTP 401.
	ErrUnauthorized = errors.New("glassnode: unauthorized — check your API key")
	// ErrNotFound is returned when the API responds with HTTP 404.
	ErrNotFound = errors.New("glassnode: resource not found")
	// ErrRateLimited is returned when the API responds with HTTP 429 and
	// all retry attempts have been exhausted.
	ErrRateLimited = errors.New("glassnode: rate limit exceeded")
	// ErrInvalidMetricPath is returned when an empty or malformed metric
	// path is provided.
	ErrInvalidMetricPath = errors.New("glassnode: invalid metric path")
)

// wrappedAPIError wraps an *APIError together with a sentinel error so
// that both errors.As(&APIError{}) and errors.Is(sentinel) work on the
// value returned to the caller.
type wrappedAPIError struct {
	*APIError
	sentinel error
}

// Unwrap allows errors.Is/errors.As to traverse to both the *APIError
// and the sentinel error. Go 1.20+ supports returning a slice for
// multi-error unwrapping.
func (w *wrappedAPIError) Unwrap() []error {
	return []error{w.APIError, w.sentinel}
}

// newAPIError builds an error value for the given status code, message
// and raw body, attaching the appropriate sentinel error when the status
// code is a well-known one (400, 401, 404, 429).
func newAPIError(statusCode int, message string, rawBody []byte, meta *ResponseMetadata, retried int) error {
	apiErr := &APIError{
		StatusCode: statusCode,
		Message:    message,
		RawBody:    truncateBody(rawBody),
		Retried:    retried,
	}
	if meta != nil {
		apiErr.RequestID = meta.RequestID
		apiErr.RateLimitLimit = meta.RateLimitLimit
		apiErr.RateLimitRemaining = meta.RateLimitRemaining
		apiErr.RateLimitReset = meta.RateLimitReset
	}

	switch statusCode {
	case 400:
		return &wrappedAPIError{APIError: apiErr, sentinel: ErrBadRequest}
	case 401:
		return &wrappedAPIError{APIError: apiErr, sentinel: ErrUnauthorized}
	case 404:
		return &wrappedAPIError{APIError: apiErr, sentinel: ErrNotFound}
	case 429:
		return &wrappedAPIError{APIError: apiErr, sentinel: ErrRateLimited}
	default:
		return apiErr
	}
}

// truncateBody limits the raw body to 512 bytes for safe storage in
// error values.
func truncateBody(body []byte) []byte {
	if len(body) > 512 {
		return body[:512]
	}
	return body
}
