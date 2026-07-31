package glassnode

import (
	"context"
	"encoding/json"
)

// MacroService provides access to the Macro category endpoints.
//
// Documentation: https://docs.glassnode.com/basic-api/endpoints/macro
type MacroService struct {
	client *Client
}

func (s *MacroService) getMetric(ctx context.Context, name string, q *MetricQuery, out interface{}) error {
	_, err := s.client.get(ctx, metricsPath("/macro/"+name), q.ToValues(), out)
	return err
}

// USInterestRate retrieves the US interest rate.
// Path: /macro/usa_interest_rate
func (s *MacroService) USInterestRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "usa_interest_rate", q, &out)
}

// USCentralBankBalanceSheet retrieves the US central bank balance sheet.
// Path: /macro/usa_central_bank_balance_sheet
func (s *MacroService) USCentralBankBalanceSheet(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "usa_central_bank_balance_sheet", q, &out)
}

// USGDP retrieves the US GDP.
// Path: /macro/usa_gdp
func (s *MacroService) USGDP(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "usa_gdp", q, &out)
}

// USCoreInflationRate retrieves the US core inflation rate.
// Path: /macro/usa_core_inflation_rate
func (s *MacroService) USCoreInflationRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "usa_core_inflation_rate", q, &out)
}

// USMoneySupplyM2 retrieves the US M2 money supply.
// Path: /macro/usa_money_supply_m2
func (s *MacroService) USMoneySupplyM2(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "usa_money_supply_m2", q, &out)
}

// EuroAreaInterestRate retrieves the Euro area interest rate.
// Path: /macro/eu_interest_rate
func (s *MacroService) EuroAreaInterestRate(ctx context.Context, q *MetricQuery) ([]TimePoint, error) {
	var out []TimePoint
	return out, s.getMetric(ctx, "eu_interest_rate", q, &out)
}

// GetRaw retrieves any macro metric as raw JSON.
func (s *MacroService) GetRaw(ctx context.Context, metricName string, q *MetricQuery) (json.RawMessage, error) {
	var out json.RawMessage
	return out, s.getMetric(ctx, metricName, q, &out)
}
