package glassnode

import (
	"context"
	"encoding/json"
)

// AddressesService provides access to the Addresses category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/addresses
type AddressesService struct {
	client *Client
}

func (s *AddressesService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/addresses/"+name), q.ToValues(), out)
	return err
}

// AccumulationCount retrieves the number of unique accumulation addresses.
// Path: /addresses/accumulation_count
func (s *AddressesService) AccumulationCount(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "accumulation_count", q, &out)
}

// AccumulationBalance retrieves the total funds held in accumulation addresses.
// Path: /addresses/accumulation_balance
func (s *AddressesService) AccumulationBalance(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "accumulation_balance", q, &out)
}

// ActiveCount retrieves the number of unique active addresses.
// Path: /addresses/active_count
func (s *AddressesService) ActiveCount(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "active_count", q, &out)
}

// ActiveCountWithContracts retrieves active addresses including smart contract calls.
// Path: /addresses/active_count_with_contracts
func (s *AddressesService) ActiveCountWithContracts(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "active_count_with_contracts", q, &out)
}

// Count retrieves the total number of addresses with a non-zero balance.
// Path: /addresses/count
func (s *AddressesService) Count(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "count", q, &out)
}

// NewAddresses retrieves the number of new addresses.
// Path: /addresses/new_entities
func (s *AddressesService) NewAddresses(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "new_entities", q, &out)
}

// SendingAddresses retrieves the number of sending addresses.
// Path: /addresses/sending_count
func (s *AddressesService) SendingAddresses(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "sending_count", q, &out)
}

// ReceivingAddresses retrieves the number of receiving addresses.
// Path: /addresses/receiving_count
func (s *AddressesService) ReceivingAddresses(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "receiving_count", q, &out)
}

// ActiveCountPIT retrieves the PIT version of active addresses count.
// Path: /addresses/active_count_pit
func (s *AddressesService) ActiveCountPIT(ctx context.Context, q *MetricQuery) ([]PITTimePoint, error) {
	var out []PITTimePoint
	return out, s.getMetric(ctx, "active_count_pit", q, &out)
}

// AccumulationCountPIT retrieves the PIT version of accumulation count.
// Path: /addresses/accumulation_count_pit
func (s *AddressesService) AccumulationCountPIT(ctx context.Context, q *MetricQuery) ([]PITTimePoint, error) {
	var out []PITTimePoint
	return out, s.getMetric(ctx, "accumulation_count_pit", q, &out)
}

// GetRaw retrieves any addresses metric as raw JSON.
func (s *AddressesService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
