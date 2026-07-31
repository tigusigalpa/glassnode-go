package glassnode

import (
	"context"
	"encoding/json"
)

// DistributionService provides access to the Distribution category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/distribution
type DistributionService struct {
	client *Client
}

func (s *DistributionService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/distribution/"+name), q.ToValues(), out)
	return err
}

// ExchangeBalanceTotal retrieves total exchange balance.
// Path: /distribution/balance_exchanges
func (s *DistributionService) ExchangeBalanceTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "balance_exchanges", q, &out)
}

// ExchangeBalancePercent retrieves exchange balance as a percentage.
// Path: /distribution/balance_exchanges_percent
func (s *DistributionService) ExchangeBalancePercent(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "balance_exchanges_percent", q, &out)
}

// ExchangeNetPositionChange retrieves the exchange net position change.
// Path: /distribution/exchange_net_position_change
func (s *DistributionService) ExchangeNetPositionChange(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "exchange_net_position_change", q, &out)
}

// MinerBalance retrieves the miner balance.
// Path: /distribution/miner_balance
func (s *DistributionService) MinerBalance(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "miner_balance", q, &out)
}

// MinerNetPositionChange retrieves the miner net position change.
// Path: /distribution/miner_net_position_change
func (s *DistributionService) MinerNetPositionChange(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "miner_net_position_change", q, &out)
}

// GiniCoefficient retrieves the Gini coefficient.
// Path: /distribution/gini
func (s *DistributionService) GiniCoefficient(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "gini", q, &out)
}

// HerfindahlIndex retrieves the Herfindahl index.
// Path: /distribution/herfindahl
func (s *DistributionService) HerfindahlIndex(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "herfindahl", q, &out)
}

// SupplyInSmartContracts retrieves supply held in smart contracts.
// Path: /distribution/supply_contracts
func (s *DistributionService) SupplyInSmartContracts(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "supply_contracts", q, &out)
}

// GetRaw retrieves any distribution metric as raw JSON.
func (s *DistributionService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
