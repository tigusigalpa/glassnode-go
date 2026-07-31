package glassnode

import (
	"context"
	"encoding/json"
)

// DeFiService provides access to the DeFi category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/defi
type DeFiService struct {
	client *Client
}

func (s *DeFiService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/defi/"+name), q.ToValues(), out)
	return err
}

// TVL retrieves the total value locked in DeFi.
// Path: /defi/total_value_locked
func (s *DeFiService) TVL(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "total_value_locked", q, &out)
}

// GetRaw retrieves any DeFi metric as raw JSON.
func (s *DeFiService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
