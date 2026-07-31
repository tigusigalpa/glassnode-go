package glassnode

import (
	"context"
	"encoding/json"
)

// TreasuriesService provides access to the Treasuries category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/treasuries
type TreasuriesService struct {
	client *Client
}

func (s *TreasuriesService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/treasuries/"+name), q.ToValues(), out)
	return err
}

// TreasuryBalancesCompanies retrieves treasury balances for companies.
// Path: /treasuries/balance_companies
func (s *TreasuriesService) TreasuryBalancesCompanies(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "balance_companies", q, &out)
}

// TreasuryBalancesGovernments retrieves treasury balances for governments.
// Path: /treasuries/balance_governments
func (s *TreasuriesService) TreasuryBalancesGovernments(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "balance_governments", q, &out)
}

// TreasuryNetFlowsCompanies retrieves treasury net flows for companies.
// Path: /treasuries/flows_companies
func (s *TreasuriesService) TreasuryNetFlowsCompanies(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "flows_companies", q, &out)
}

// TreasuryNetFlowsGovernments retrieves treasury net flows for governments.
// Path: /treasuries/flows_governments
func (s *TreasuriesService) TreasuryNetFlowsGovernments(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "flows_governments", q, &out)
}

// TreasuryCountCompanies retrieves the treasury count for companies.
// Path: /treasuries/count_companies
func (s *TreasuriesService) TreasuryCountCompanies(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "count_companies", q, &out)
}

// GetRaw retrieves any treasuries metric as raw JSON.
func (s *TreasuriesService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
