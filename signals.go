package glassnode

import (
	"context"
	"encoding/json"
)

// SignalsService provides access to the Signals category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/signals
type SignalsService struct {
	client *Client
}

func (s *SignalsService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/signals/"+name), q.ToValues(), out)
	return err
}

// AltcoinCycleSignal retrieves the Altcoin Cycle Signal.
// Path: /signals/altcoin_cycle_signal
func (s *SignalsService) AltcoinCycleSignal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "altcoin_cycle_signal", q, &out)
}

// BitcoinRiskSignal retrieves the Bitcoin Risk Signal.
// Path: /signals/bitcoin_risk_signal
func (s *SignalsService) BitcoinRiskSignal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "bitcoin_risk_signal", q, &out)
}

// BitcoinSharpeSignal retrieves the Bitcoin Sharpe Signal.
// Path: /signals/bitcoin_sharpe_signal
func (s *SignalsService) BitcoinSharpeSignal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "bitcoin_sharpe_signal", q, &out)
}

// BSSGoldilocksSignal retrieves the BSS Goldilocks Signal.
// Path: /signals/bss_goldilocks_signal
func (s *SignalsService) BSSGoldilocksSignal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "bss_goldilocks_signal", q, &out)
}

// EcosystemMomentumSignal retrieves the Ecosystem Momentum Signal.
// Path: /signals/ecosystem_momentum_signal
func (s *SignalsService) EcosystemMomentumSignal(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "ecosystem_momentum_signal", q, &out)
}

// GetRaw retrieves any signals metric as raw JSON.
func (s *SignalsService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
