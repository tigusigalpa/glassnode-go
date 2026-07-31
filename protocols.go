package glassnode

import (
	"context"
	"encoding/json"
)

// ProtocolsService provides access to the Protocols category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/protocols
type ProtocolsService struct {
	client *Client
}

func (s *ProtocolsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/protocols/"+name), q.ToValues(), out)
	return err
}

// AaveV3AvailableLiquidity retrieves Aave V3 available liquidity by token.
// Path: /protocols/aave-v3/tvl_by_token
func (s *ProtocolsService) AaveV3AvailableLiquidity(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "aave-v3/tvl_by_token", q, &out)
}

// UniswapTVLPerFeeTier retrieves Uniswap TVL per fee tier.
// Path: /protocols/uniswap/tvl_per_fee_tier
func (s *ProtocolsService) UniswapTVLPerFeeTier(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "uniswap/tvl_per_fee_tier", q, &out)
}

// GetRaw retrieves any protocols metric as raw JSON.
func (s *ProtocolsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
