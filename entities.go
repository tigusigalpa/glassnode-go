package glassnode

import (
	"context"
	"encoding/json"
)

// EntitiesService provides access to the Entities category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/entities
type EntitiesService struct {
	client *Client
}

func (s *EntitiesService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/entities/"+name), q.ToValues(), out)
	return err
}

// ActiveEntities retrieves the number of active entities.
// Path: /entities/active_count
func (s *EntitiesService) ActiveEntities(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "active_count", q, &out)
}

// NewEntities retrieves the number of new entities.
// Path: /entities/new_count
func (s *EntitiesService) NewEntities(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "new_count", q, &out)
}

// SendingEntities retrieves the number of sending entities.
// Path: /entities/sending_count
func (s *EntitiesService) SendingEntities(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "sending_count", q, &out)
}

// ReceivingEntities retrieves the number of receiving entities.
// Path: /entities/receiving_count
func (s *EntitiesService) ReceivingEntities(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "receiving_count", q, &out)
}

// EntitiesNetGrowth retrieves the net growth in entities.
// Path: /entities/net_growth_count
func (s *EntitiesService) EntitiesNetGrowth(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "net_growth_count", q, &out)
}

// NumberWhales retrieves the number of whale entities.
// Path: /entities/count
func (s *EntitiesService) NumberWhales(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "count", q, &out)
}

// GetRaw retrieves any entities metric as raw JSON.
func (s *EntitiesService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
