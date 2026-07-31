package glassnode

import (
	"context"
	"encoding/json"
)

// FeesService provides access to the Fees category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/fees
type FeesService struct {
	client *Client
}

func (s *FeesService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/fees/"+name), q.ToValues(), out)
	return err
}

// FeesTotal retrieves total fees.
// Path: /fees/volume_sum
func (s *FeesService) FeesTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_sum", q, &out)
}

// FeesMean retrieves mean fees per transaction.
// Path: /fees/volume_mean
func (s *FeesService) FeesMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_mean", q, &out)
}

// FeesMedian retrieves median fees per transaction.
// Path: /fees/volume_median
func (s *FeesService) FeesMedian(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_median", q, &out)
}

// GasPriceMean retrieves the mean gas price.
// Path: /fees/gas_price_mean
func (s *FeesService) GasPriceMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "gas_price_mean", q, &out)
}

// GasPriceMedian retrieves the median gas price.
// Path: /fees/gas_price_median
func (s *FeesService) GasPriceMedian(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "gas_price_median", q, &out)
}

// GasUsedTotal retrieves total gas used.
// Path: /fees/gas_used_sum
func (s *FeesService) GasUsedTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "gas_used_sum", q, &out)
}

// FeeRatioMultiple retrieves the Fee Ratio Multiple (FRM).
// Path: /fees/fee_ratio_multiple
func (s *FeesService) FeeRatioMultiple(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "fee_ratio_multiple", q, &out)
}

// GetRaw retrieves any fees metric as raw JSON.
func (s *FeesService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
