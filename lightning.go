package glassnode

import (
	"context"
	"encoding/json"
)

// LightningService provides access to the Lightning Network category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/lightning
type LightningService struct {
	client *Client
}

func (s *LightningService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/lightning/"+name), q.ToValues(), out)
	return err
}

// Capacity retrieves the Lightning Network capacity.
// Path: /lightning/capacity
func (s *LightningService) Capacity(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "capacity", q, &out)
}

// NumberOfNodes retrieves the number of Lightning Network nodes.
// Path: /lightning/node_count
func (s *LightningService) NumberOfNodes(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "node_count", q, &out)
}

// NumberOfChannels retrieves the number of Lightning Network channels.
// Path: /lightning/channel_count
func (s *LightningService) NumberOfChannels(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "channel_count", q, &out)
}

// ChannelSizeMean retrieves the mean channel size.
// Path: /lightning/channel_size_mean
func (s *LightningService) ChannelSizeMean(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "channel_size_mean", q, &out)
}

// ChannelSizeMedian retrieves the median channel size.
// Path: /lightning/channel_size_median
func (s *LightningService) ChannelSizeMedian(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "channel_size_median", q, &out)
}

// BaseFeeMedian retrieves the median base fee.
// Path: /lightning/base_fee_median
func (s *LightningService) BaseFeeMedian(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "base_fee_median", q, &out)
}

// FeeRateMedian retrieves the median fee rate.
// Path: /lightning/fee_rate_median
func (s *LightningService) FeeRateMedian(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "fee_rate_median", q, &out)
}

// GetRaw retrieves any lightning metric as raw JSON.
func (s *LightningService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
