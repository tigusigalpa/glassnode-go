package glassnode

import (
	"context"
	"encoding/json"
)

// SupplyService provides access to the Supply category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/supply
type SupplyService struct {
	client *Client
}

func (s *SupplyService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/supply/"+name), q.ToValues(), out)
	return err
}

// CirculatingSupply retrieves the circulating supply.
// Path: /supply/current
func (s *SupplyService) CirculatingSupply(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "current", q, &out)
}

// ActiveSupply retrieves the active supply.
// Path: /supply/active
func (s *SupplyService) ActiveSupply(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "active", q, &out)
}

// LongTermHolderSupply retrieves the long-term holder supply.
// Path: /supply/lth
func (s *SupplyService) LongTermHolderSupply(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "lth", q, &out)
}

// ShortTermHolderSupply retrieves the short-term holder supply.
// Path: /supply/sth
func (s *SupplyService) ShortTermHolderSupply(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "sth", q, &out)
}

// InflationRate retrieves the inflation rate.
// Path: /supply/inflation_rate
func (s *SupplyService) InflationRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "inflation_rate", q, &out)
}

// ProfitableSupply retrieves the supply in profit.
// Path: /supply/profit_relative
func (s *SupplyService) ProfitableSupply(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "profit_relative", q, &out)
}

// HODLWaves retrieves the HODL Waves distribution.
// Path: /supply/hodl_waves
func (s *SupplyService) HODLWaves(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "hodl_waves", q, &out)
}

// GetRaw retrieves any supply metric as raw JSON.
func (s *SupplyService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
