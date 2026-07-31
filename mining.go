package glassnode

import (
	"context"
	"encoding/json"
)

// MiningService provides access to the Mining category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/mining
type MiningService struct {
	client *Client
}

func (s *MiningService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/mining/"+name), q.ToValues(), out)
	return err
}

// HashRate retrieves the network hash rate.
// Path: /mining/hash_rate_mean
func (s *MiningService) HashRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "hash_rate_mean", q, &out)
}

// Difficulty retrieves the mining difficulty.
// Path: /mining/difficulty_latest
func (s *MiningService) Difficulty(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "difficulty_latest", q, &out)
}

// MinerRevenueTotal retrieves total miner revenue.
// Path: /mining/revenue_sum
func (s *MiningService) MinerRevenueTotal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "revenue_sum", q, &out)
}

// MinerRevenueBlockRewards retrieves miner revenue from block rewards.
// Path: /mining/revenue_from_rewards
func (s *MiningService) MinerRevenueBlockRewards(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "revenue_from_rewards", q, &out)
}

// MinerRevenueFees retrieves miner revenue from fees.
// Path: /mining/revenue_from_fees
func (s *MiningService) MinerRevenueFees(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "revenue_from_fees", q, &out)
}

// Thermocap retrieves the Thermocap metric.
// Path: /mining/thermocap
func (s *MiningService) Thermocap(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "thermocap", q, &out)
}

// GetRaw retrieves any mining metric as raw JSON.
func (s *MiningService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
