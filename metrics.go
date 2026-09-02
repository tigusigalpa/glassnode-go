package glassnode

import (
	"context"
	"encoding/json"
)

// MetricsService provides generic access to any Glassnode metric
// endpoint. It guarantees coverage of every documented or future metric
// path without waiting for a typed wrapper.
//
// Use Metadata.Metric to discover whether a metric supports bulk
// endpoints (bulk_supported field) before calling GetBulk.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints
type MetricsService struct {
	client *Client
}

// Get retrieves a single metric as a raw JSON response. This method
// safely calls every valid metric path, even if no convenience wrapper
// exists. The metricPath should be the metric portion only (e.g.
// "/indicators/sopr"), without the "/v1/metrics" prefix.
//
// The returned json.RawMessage can be unmarshaled into []TimePoint,
// []ObjectPoint, []PITTimePoint, or []PITObjectPoint depending on the
// metric's response shape. When query.Format is FormatCSV, it instead
// contains the unmodified CSV response bytes. Use the metadata endpoint
// to determine the expected shape.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints
func (s *MetricsService) Get(ctx context.Context, metricPath string, query *MetricQuery) (json.RawMessage, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out json.RawMessage
	_, err := s.client.get(ctx, metricsPath(metricPath), query.ToValues(), &out)
	return out, err
}

// GetTimePoints retrieves a metric and decodes it as a slice of
// TimePoint (scalar time-series: {"t":..., "v":...}).
func (s *MetricsService) GetTimePoints(ctx context.Context, metricPath string, query *MetricQuery) ([]TimePoint, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out []TimePoint
	_, err := s.client.get(ctx, metricsPath(metricPath), query.ToValues(), &out)
	return out, err
}

// GetObjectPoints retrieves a metric and decodes it as a slice of
// ObjectPoint (object time-series: {"t":..., "o":...}).
func (s *MetricsService) GetObjectPoints(ctx context.Context, metricPath string, query *MetricQuery) ([]ObjectPoint, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out []ObjectPoint
	_, err := s.client.get(ctx, metricsPath(metricPath), query.ToValues(), &out)
	return out, err
}

// GetBulk retrieves a bulk metric response. The metricPath should be
// the regular metric's path (e.g. "/market/mvrv"); the "/bulk" suffix
// is appended automatically. Always specify assets explicitly in the
// BulkQuery to control credit consumption.
//
// Documentation: https://docs.glassnode.com/basic-api/bulk-metrics
func (s *MetricsService) GetBulk(ctx context.Context, metricPath string, query *BulkQuery) (*BulkResponse, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out BulkResponse
	_, err := s.client.get(ctx, bulkMetricsPath(metricPath), query.ToValues(), &out)
	return &out, err
}

// GetBulkRaw retrieves a bulk metric response as raw JSON, for
// endpoints whose bulk shape cannot be safely represented with
// BulkResponse.
func (s *MetricsService) GetBulkRaw(ctx context.Context, metricPath string, query *BulkQuery) (json.RawMessage, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out json.RawMessage
	_, err := s.client.get(ctx, bulkMetricsPath(metricPath), query.ToValues(), &out)
	return out, err
}
