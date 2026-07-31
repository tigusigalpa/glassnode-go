package glassnode

import (
	"context"
	"encoding/json"
)

// TransactionsService provides access to the Transactions category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/transactions
type TransactionsService struct {
	client *Client
}

func (s *TransactionsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/transactions/"+name), q.ToValues(), out)
	return err
}

// Count retrieves the transaction count.
// Path: /transactions/count
func (s *TransactionsService) Count(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "count", q, &out)
}

// Rate retrieves the transaction rate.
// Path: /transactions/rate
func (s *TransactionsService) Rate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "rate", q, &out)
}

// SizeMean retrieves the mean transaction size.
// Path: /transactions/size_mean
func (s *TransactionsService) SizeMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "size_mean", q, &out)
}

// SizeTotal retrieves the total transaction size.
// Path: /transactions/size_sum
func (s *TransactionsService) SizeTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "size_sum", q, &out)
}

// TransferVolumeMean retrieves the mean transfer volume.
// Path: /transactions/volume_mean
func (s *TransactionsService) TransferVolumeMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_mean", q, &out)
}

// TransferVolumeTotal retrieves the total transfer volume.
// Path: /transactions/volume_sum
func (s *TransactionsService) TransferVolumeTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "volume_sum", q, &out)
}

// ExchangeDeposits retrieves the number of exchange deposits.
// Path: /transactions/deposits_count
func (s *TransactionsService) ExchangeDeposits(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "deposits_count", q, &out)
}

// ExchangeWithdrawals retrieves the number of exchange withdrawals.
// Path: /transactions/withdrawals_count
func (s *TransactionsService) ExchangeWithdrawals(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "withdrawals_count", q, &out)
}

// GetRaw retrieves any transactions metric as raw JSON.
func (s *TransactionsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
