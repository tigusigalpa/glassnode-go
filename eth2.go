package glassnode

import (
	"context"
	"encoding/json"
)

// ETH2Service provides access to the ETH 2.0 category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/eth2
type ETH2Service struct {
	client *Client
}

func (s *ETH2Service) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/eth2/"+name), q.ToValues(), out)
	return err
}

// ActiveValidators retrieves the number of active validators.
// Path: /eth2/validators_active
func (s *ETH2Service) ActiveValidators(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "validators_active", q, &out)
}

// TotalValueStaked retrieves the total value staked in ETH 2.0.
// Path: /eth2/staked_total
func (s *ETH2Service) TotalValueStaked(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "staked_total", q, &out)
}

// TotalValidators retrieves the total number of validators.
// Path: /eth2/validators_count
func (s *ETH2Service) TotalValidators(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "validators_count", q, &out)
}

// AverageValidatorBalance retrieves the average validator balance.
// Path: /eth2/validator_balance_mean
func (s *ETH2Service) AverageValidatorBalance(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "validator_balance_mean", q, &out)
}

// ParticipationRate retrieves the participation rate.
// Path: /eth2/participation_rate
func (s *ETH2Service) ParticipationRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "participation_rate", q, &out)
}

// EstimatedAnnualIssuance retrieves the estimated annual issuance.
// Path: /eth2/issuance_annualized
func (s *ETH2Service) EstimatedAnnualIssuance(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "issuance_annualized", q, &out)
}

// GetRaw retrieves any ETH2 metric as raw JSON.
func (s *ETH2Service) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
