package glassnode

import (
	"context"
	"encoding/json"
)

// BlockchainService provides access to the Blockchain category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/blockchain
type BlockchainService struct {
	client *Client
}

func (s *BlockchainService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/blockchain/"+name), q.ToValues(), out)
	return err
}

// BlockHeight retrieves the current block height.
// Path: /blockchain/block_height
func (s *BlockchainService) BlockHeight(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "block_height", q, &out)
}

// BlockIntervalMean retrieves the mean block interval.
// Path: /blockchain/block_interval_mean
func (s *BlockchainService) BlockIntervalMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "block_interval_mean", q, &out)
}

// BlockIntervalMedian retrieves the median block interval.
// Path: /blockchain/block_interval_median
func (s *BlockchainService) BlockIntervalMedian(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "block_interval_median", q, &out)
}

// BlockSizeMean retrieves the mean block size.
// Path: /blockchain/block_size_mean
func (s *BlockchainService) BlockSizeMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "block_size_mean", q, &out)
}

// BlockSizeTotal retrieves the total block size.
// Path: /blockchain/block_size_total
func (s *BlockchainService) BlockSizeTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "block_size_total", q, &out)
}

// BlocksMined retrieves the number of blocks mined.
// Path: /blockchain/blocks_mined
func (s *BlockchainService) BlocksMined(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "blocks_mined", q, &out)
}

// UTXOsTotal retrieves the total number of UTXOs.
// Path: /blockchain/utxo_count
func (s *BlockchainService) UTXOsTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "utxo_count", q, &out)
}

// UTXOsCreated retrieves the number of UTXOs created.
// Path: /blockchain/utxo_created_count
func (s *BlockchainService) UTXOsCreated(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "utxo_created_count", q, &out)
}

// UTXOsSpent retrieves the number of UTXOs spent.
// Path: /blockchain/utxo_spent_count
func (s *BlockchainService) UTXOsSpent(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "utxo_spent_count", q, &out)
}

// GetRaw retrieves any blockchain metric as raw JSON.
func (s *BlockchainService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
