package glassnode

import (
	"context"
	"encoding/json"
)

// OptionsService provides access to the Options category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/options
type OptionsService struct {
	client *Client
}

func (s *OptionsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/options/"+name), q.ToValues(), out)
	return err
}

// OpenInterest retrieves options open interest.
// Path: /options/open_interest_sum
func (s *OptionsService) OpenInterest(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "open_interest_sum", q, &out)
}

// Volume retrieves options volume.
// Path: /options/volume_sum
func (s *OptionsService) Volume(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_sum", q, &out)
}

// PutCallRatio retrieves the options put/call ratio.
// Path: /options/volume_put_call_ratio
func (s *OptionsService) PutCallRatio(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_put_call_ratio", q, &out)
}

// MaxPain retrieves the options max pain.
// Path: /options/max_pain
func (s *OptionsService) MaxPain(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "max_pain", q, &out)
}

// ImpliedVolatility retrieves ATM implied volatility.
// Path: /options/iv_atm
func (s *OptionsService) ImpliedVolatility(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "iv_atm", q, &out)
}

// GetRaw retrieves any options metric as raw JSON.
func (s *OptionsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
