package glassnode

import (
	"context"
	"encoding/json"
)

// InstitutionsService provides access to the Institutions category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/institutions
type InstitutionsService struct {
	client *Client
}

func (s *InstitutionsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/institutions/"+name), q.ToValues(), out)
	return err
}

// USSpotETFBalances retrieves US Spot ETF balances.
// Path: /institutions/holdings_us_spot_etf
func (s *InstitutionsService) USSpotETFBalances(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "holdings_us_spot_etf", q, &out)
}

// USSpotETFNetFlows retrieves US Spot ETF net flows.
// Path: /institutions/flows_us_spot_etf
func (s *InstitutionsService) USSpotETFNetFlows(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "flows_us_spot_etf", q, &out)
}

// USSpotETFTradingVolume retrieves US Spot ETF trading volume.
// Path: /institutions/volume_us_spot_etf
func (s *InstitutionsService) USSpotETFTradingVolume(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_us_spot_etf", q, &out)
}

// USSpotETFPrice retrieves US Spot ETF price.
// Path: /institutions/price_us_spot_etf
func (s *InstitutionsService) USSpotETFPrice(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "price_us_spot_etf", q, &out)
}

// PurposeBitcoinETFHoldings retrieves Purpose Bitcoin ETF holdings.
// Path: /institutions/holdings_purpose_etf
func (s *InstitutionsService) PurposeBitcoinETFHoldings(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "holdings_purpose_etf", q, &out)
}

// GetRaw retrieves any institutions metric as raw JSON.
func (s *InstitutionsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
