package glassnode

import (
	"context"
	"encoding/json"
)

// GlobalService provides access to the Global category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/global
type GlobalService struct {
	client *Client
}

func (s *GlobalService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/global/"+name), q.ToValues(), out)
	return err
}

// TotalMarketCap retrieves the total market capitalization.
// Path: /global/marketcap
func (s *GlobalService) TotalMarketCap(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "marketcap", q, &out)
}

// TotalRealizedCap retrieves the total realized capitalization.
// Path: /global/marketcap_realized
func (s *GlobalService) TotalRealizedCap(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "marketcap_realized", q, &out)
}

// TotalOpenInterest retrieves the total open interest.
// Path: /global/open_interest
func (s *GlobalService) TotalOpenInterest(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "open_interest", q, &out)
}

// MedianSOPR retrieves the median SOPR.
// Path: /global/sopr_median
func (s *GlobalService) MedianSOPR(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "sopr_median", q, &out)
}

// GetRaw retrieves any global metric as raw JSON.
func (s *GlobalService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
