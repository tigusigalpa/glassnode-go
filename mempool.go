package glassnode

import (
	"context"
	"encoding/json"
)

// MempoolService provides access to the Mempool category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/mempool
type MempoolService struct {
	client *Client
}

func (s *MempoolService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/mempool/"+name), q.ToValues(), out)
	return err
}

// TransactionCount retrieves the mempool transaction count.
// Path: /mempool/count
func (s *MempoolService) TransactionCount(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "count", q, &out)
}

// TotalAmountOfFees retrieves the total amount of fees in the mempool.
// Path: /mempool/fees_sum
func (s *MempoolService) TotalAmountOfFees(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "fees_sum", q, &out)
}

// TotalAmountOfCoins retrieves the total amount of coins in the mempool.
// Path: /mempool/value_sum
func (s *MempoolService) TotalAmountOfCoins(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "value_sum", q, &out)
}

// TransactionsSize retrieves the total size of mempool transactions.
// Path: /mempool/size_sum
func (s *MempoolService) TransactionsSize(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "size_sum", q, &out)
}

// AverageRelativeFee retrieves the mempool average relative fee.
// Path: /mempool/fees_relative_mean
func (s *MempoolService) AverageRelativeFee(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "fees_relative_mean", q, &out)
}

// GetRaw retrieves any mempool metric as raw JSON.
func (s *MempoolService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
