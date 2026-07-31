package glassnode

import (
	"context"
	"encoding/json"
)

// BridgesService provides access to the Bridges category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/bridges
type BridgesService struct {
	client *Client
}

func (s *BridgesService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/bridges/"+name), q.ToValues(), out)
	return err
}

// TVL retrieves the total value locked in bridges.
// Path: /bridges/tvl
func (s *BridgesService) TVL(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "tvl", q, &out)
}

// TVLRelative retrieves the relative TVL across bridges.
// Path: /bridges/tvl_relative
func (s *BridgesService) TVLRelative(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "tvl_relative", q, &out)
}

// DepositsByChain retrieves bridge deposits by chain.
// Path: /bridges/deposits_by_chain
func (s *BridgesService) DepositsByChain(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "deposits_by_chain", q, &out)
}

// NetFlowByChain retrieves bridge net flow by chain.
// Path: /bridges/net_flow_by_chain
func (s *BridgesService) NetFlowByChain(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "net_flow_by_chain", q, &out)
}

// WithdrawalsByChain retrieves bridge withdrawals by chain.
// Path: /bridges/withdrawals_by_chain
func (s *BridgesService) WithdrawalsByChain(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "withdrawals_by_chain", q, &out)
}

// GetRaw retrieves any bridges metric as raw JSON.
func (s *BridgesService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
