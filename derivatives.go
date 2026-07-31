package glassnode

import (
	"context"
	"encoding/json"
)

// DerivativesService provides access to the Derivatives category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/derivatives
type DerivativesService struct {
	client *Client
}

func (s *DerivativesService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/derivatives/"+name), q.ToValues(), out)
	return err
}

// FuturesOpenInterest retrieves futures open interest.
// Path: /derivatives/futures_open_interest
func (s *DerivativesService) FuturesOpenInterest(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "futures_open_interest", q, &out)
}

// FuturesVolume retrieves futures trading volume.
// Path: /derivatives/futures_volume
func (s *DerivativesService) FuturesVolume(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "futures_volume", q, &out)
}

// FuturesPerpetualFundingRate retrieves the perpetual funding rate.
// Path: /derivatives/futures_perpetual_funding_rate
func (s *DerivativesService) FuturesPerpetualFundingRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "futures_perpetual_funding_rate", q, &out)
}

// FuturesEstimatedLeverageRatio retrieves the estimated leverage ratio.
// Path: /derivatives/futures_estimated_leverage_ratio
func (s *DerivativesService) FuturesEstimatedLeverageRatio(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "futures_estimated_leverage_ratio", q, &out)
}

// OptionsOpenInterest retrieves options open interest.
// Path: /derivatives/options_open_interest
func (s *DerivativesService) OptionsOpenInterest(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "options_open_interest", q, &out)
}

// OptionsVolume24h retrieves 24h options volume.
// Path: /derivatives/options_volume_24h
func (s *DerivativesService) OptionsVolume24h(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "options_volume_24h", q, &out)
}

// FuturesLiquidationsLongTotal retrieves total long liquidations.
// Path: /derivatives/futures_long_liquidations_total
func (s *DerivativesService) FuturesLiquidationsLongTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "futures_long_liquidations_total", q, &out)
}

// FuturesLiquidationsShortTotal retrieves total short liquidations.
// Path: /derivatives/futures_short_liquidations_total
func (s *DerivativesService) FuturesLiquidationsShortTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "futures_short_liquidations_total", q, &out)
}

// GetRaw retrieves any derivatives metric as raw JSON.
func (s *DerivativesService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
