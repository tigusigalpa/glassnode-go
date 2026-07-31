package glassnode

import (
	"context"
	"encoding/json"
)

// BreakdownsService provides access to the Breakdowns category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/breakdowns
type BreakdownsService struct {
	client *Client
}

func (s *BreakdownsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/breakdowns/"+name), q.ToValues(), out)
	return err
}

// SupplyByAge retrieves supply breakdown by age.
// Path: /breakdowns/supply_by_age
func (s *BreakdownsService) SupplyByAge(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "supply_by_age", q, &out)
}

// SupplyByWalletSize retrieves supply breakdown by wallet size.
// Path: /breakdowns/supply_by_wallet_size
func (s *BreakdownsService) SupplyByWalletSize(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "supply_by_wallet_size", q, &out)
}

// SupplyByPnL retrieves supply breakdown by profit/loss.
// Path: /breakdowns/supply_by_pnl
func (s *BreakdownsService) SupplyByPnL(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "supply_by_pnl", q, &out)
}

// SOPRByAge retrieves SOPR breakdown by age.
// Path: /breakdowns/sopr_by_age
func (s *BreakdownsService) SOPRByAge(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "sopr_by_age", q, &out)
}

// SOPRByWalletSize retrieves SOPR breakdown by wallet size.
// Path: /breakdowns/sopr_by_wallet_size
func (s *BreakdownsService) SOPRByWalletSize(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "sopr_by_wallet_size", q, &out)
}

// MVRVByAge retrieves MVRV breakdown by age.
// Path: /breakdowns/mvrv_by_age
func (s *BreakdownsService) MVRVByAge(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "mvrv_by_age", q, &out)
}

// RealizedCapByAge retrieves realized cap breakdown by age.
// Path: /breakdowns/realized_cap_by_age
func (s *BreakdownsService) RealizedCapByAge(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "realized_cap_by_age", q, &out)
}

// GetRaw retrieves any breakdowns metric as raw JSON.
func (s *BreakdownsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
