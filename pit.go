package glassnode

import (
	"context"
	"encoding/json"
)

// PointInTimeService provides access to the Point-In-Time category endpoints.
// Point-in-Time metrics include a computed_at timestamp indicating when
// the data point was computed, which is important for backtesting and
// historical analysis.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/pit
type PointInTimeService struct {
	client *Client
}

func (s *PointInTimeService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/pit/"+name), q.ToValues(), out)
	return err
}

// GetRaw retrieves any point-in-time metric as raw JSON. The metricName
// should be the full metric path (e.g. "/indicators/sopr_pit").
func (s *PointInTimeService) GetRaw(ctx context.Context, metricPath string, q *MetricQuery) (json.RawMessage, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out json.RawMessage
	_, err := s.client.get(ctx, metricsPath(metricPath), q.ToValues(), &out)
	return out, err
}

// GetPITTimePoints retrieves a PIT metric and decodes it as PITTimePoint.
func (s *PointInTimeService) GetPITTimePoints(ctx context.Context, metricPath string, q *MetricQuery) ([]PITTimePoint, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out []PITTimePoint
	_, err := s.client.get(ctx, metricsPath(metricPath), q.ToValues(), &out)
	return out, err
}

// GetPITObjectPoints retrieves a PIT metric and decodes it as PITObjectPoint.
func (s *PointInTimeService) GetPITObjectPoints(ctx context.Context, metricPath string, q *MetricQuery) ([]PITObjectPoint, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	var out []PITObjectPoint
	_, err := s.client.get(ctx, metricsPath(metricPath), q.ToValues(), &out)
	return out, err
}
