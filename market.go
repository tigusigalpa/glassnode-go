package glassnode

import (
	"context"
	"encoding/json"
)

// MarketService provides access to the Market category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/market
type MarketService struct {
	client *Client
}

func (s *MarketService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/market/"+name), q.ToValues(), out)
	return err
}

// Price retrieves the market price.
// Path: /market/price_usd_close
func (s *MarketService) Price(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "price_usd_close", q, &out)
}

// PriceOHLC retrieves the OHLC price.
// Path: /market/price_usd_ohlc
func (s *MarketService) PriceOHLC(ctx context.Context, q *MetricQuery) ([]ObjectPoint, error) {
	var out []ObjectPoint
	return out, s.getMetric(ctx, "price_usd_ohlc", q, &out)
}

// MarketCap retrieves the market capitalization.
// Path: /market/marketcap_usd
func (s *MarketService) MarketCap(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "marketcap_usd", q, &out)
}

// MVRV retrieves the MVRV Ratio.
// Path: /market/mvrv
func (s *MarketService) MVRV(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "mvrv", q, &out)
}

// RealizedCap retrieves the realized capitalization.
// Path: /market/marketcap_realized_usd
func (s *MarketService) RealizedCap(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "marketcap_realized_usd", q, &out)
}

// RealizedPrice retrieves the realized price.
// Path: /market/price_realized_usd
func (s *MarketService) RealizedPrice(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "price_realized_usd", q, &out)
}

// BTCDominance retrieves BTC Dominance.
// Path: /market/btc_dominance
func (s *MarketService) BTCDominance(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "btc_dominance", q, &out)
}

// SpotVolume retrieves spot trading volume.
// Path: /market/volume
func (s *MarketService) SpotVolume(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume", q, &out)
}

// SpotVolume24h retrieves 24h spot trading volume.
// Path: /market/volume_usd_24h
func (s *MarketService) SpotVolume24h(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_usd_24h", q, &out)
}

// PriceDrawdownFromATH retrieves the price drawdown from all-time high.
// Path: /market/price_drawdown_relative
func (s *MarketService) PriceDrawdownFromATH(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "price_drawdown_relative", q, &out)
}

// RealizedVolatility1M retrieves 1-month realized volatility.
// Path: /market/realized_volatility_1m
func (s *MarketService) RealizedVolatility1M(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "realized_volatility_1m", q, &out)
}

// BetaBtc7D retrieves the 7-day beta to BTC.
// Path: /market/beta_btc_7d
func (s *MarketService) BetaBtc7D(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "beta_btc_7d", q, &out)
}

// GetRaw retrieves any market metric as raw JSON.
func (s *MarketService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
