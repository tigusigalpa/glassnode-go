package glassnode

import (
	"context"
	"encoding/json"
)

// IndicatorsService provides access to the Indicators category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/indicators
type IndicatorsService struct {
	client *Client
}

func (s *IndicatorsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/indicators/"+name), q.ToValues(), out)
	return err
}

// SOPR retrieves the Spent Output Profit Ratio.
// Path: /indicators/sopr
func (s *IndicatorsService) SOPR(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "sopr", q, &out)
}

// CDD retrieves Coin Days Destroyed.
// Path: /indicators/cdd
func (s *IndicatorsService) CDD(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "cdd", q, &out)
}

// AccumulationTrendScore retrieves the Accumulation Trend Score.
// Path: /indicators/accumulation_trend_score
func (s *IndicatorsService) AccumulationTrendScore(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "accumulation_trend_score", q, &out)
}

// NVT retrieves the NVT Ratio.
// Path: /indicators/nvt
func (s *IndicatorsService) NVT(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "nvt", q, &out)
}

// MVRV retrieves the MVRV Ratio (Z-Score).
// Path: /indicators/mvrv
func (s *IndicatorsService) MVRV(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "mvrv", q, &out)
}

// NUPL retrieves Net Unrealized Profit/Loss.
// Path: /indicators/nupl
func (s *IndicatorsService) NUPL(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "nupl", q, &out)
}

// RealizedProfit retrieves realized profit.
// Path: /indicators/realized_profit
func (s *IndicatorsService) RealizedProfit(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "realized_profit", q, &out)
}

// RealizedLoss retrieves realized loss.
// Path: /indicators/realized_loss
func (s *IndicatorsService) RealizedLoss(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "realized_loss", q, &out)
}

// ReserveRisk retrieves the Reserve Risk indicator.
// Path: /indicators/reserve_risk
func (s *IndicatorsService) ReserveRisk(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "reserve_risk", q, &out)
}

// PuellMultiple retrieves the Puell Multiple.
// Path: /indicators/puell_multiple
func (s *IndicatorsService) PuellMultiple(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "puell_multiple", q, &out)
}

// StockToFlow retrieves the Stock-to-Flow Ratio.
// Path: /indicators/s2f
func (s *IndicatorsService) StockToFlow(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "s2f", q, &out)
}

// RHODL retrieves the RHODL Ratio.
// Path: /indicators/rhodl_ratio
func (s *IndicatorsService) RHODL(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "rhodl_ratio", q, &out)
}

// FearGreed retrieves the Fear & Greed Index.
// Path: /indicators/fear_greed_index
func (s *IndicatorsService) FearGreed(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "fear_greed_index", q, &out)
}

// Liveliness retrieves the Liveliness indicator.
// Path: /indicators/liveliness
func (s *IndicatorsService) Liveliness(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "liveliness", q, &out)
}

// Dormancy retrieves the Dormancy indicator.
// Path: /indicators/dormancy
func (s *IndicatorsService) Dormancy(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "dormancy", q, &out)
}

// GetRaw retrieves any indicators metric as raw JSON.
func (s *IndicatorsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
