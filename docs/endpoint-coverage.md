# Endpoint Coverage

This document maps every Glassnode Basic API endpoint category to the SDK methods that cover it. It is auto-maintained and updated with each release.

## Legend

- **Typed wrapper** — a typed convenience method exists on the category service
- **Generic client** — available through `client.Metrics.Get()` / `GetTimePoints()` / `GetObjectPoints()`
- **Bulk wrapper** — `client.Metrics.GetBulk()` supports this metric path
- **PIT wrapper** — `client.PointInTime.GetPITTimePoints()` / `GetPITObjectPoints()` supports this metric
- **Not available** — not available for the tested subscription plan

## Core Endpoints

| API Path | SDK Method | Parameters | Response Model | Bulk | Status |
|---|---|---|---|---|---|
| `GET /v1/metadata/assets` | `client.Metadata.Assets()` | `filter` | `AssetsResponse` | N/A | Typed wrapper |
| `GET /v1/metadata/metrics` | `client.Metadata.Metrics()` | `MetricQuery` | `[]string` | N/A | Typed wrapper |
| `GET /v1/metadata/metric` | `client.Metadata.Metric()` | `path`, `MetricQuery` | `MetricMetadata` | N/A | Typed wrapper |
| `GET /v1/user/api_usage` | `client.User.APIUsage()` | none | `APIUsage` | N/A | Typed wrapper |
| `GET /v1/metrics/<path>` | `client.Metrics.Get()` / `GetTimePoints()` / `GetObjectPoints()` | `MetricQuery` | `json.RawMessage` / `[]TimePoint` / `[]ObjectPoint` | N/A | Generic client |
| `GET /v1/metrics/<path>/bulk` | `client.Metrics.GetBulk()` | `BulkQuery` | `BulkResponse` | Yes | Bulk wrapper |

## Category Services

### Addresses

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/addresses

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/addresses/accumulation_count` | `client.Addresses.AccumulationCount()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/accumulation_balance` | `client.Addresses.AccumulationBalance()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/active_count` | `client.Addresses.ActiveCount()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/active_count_with_contracts` | `client.Addresses.ActiveCountWithContracts()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/count` | `client.Addresses.Count()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/new_entities` | `client.Addresses.NewAddresses()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/sending_count` | `client.Addresses.SendingAddresses()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/receiving_count` | `client.Addresses.ReceivingAddresses()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/addresses/active_count_pit` | `client.Addresses.ActiveCountPIT()` | `[]PITTimePoint` | N/A | Typed wrapper |
| `/addresses/accumulation_count_pit` | `client.Addresses.AccumulationCountPIT()` | `[]PITTimePoint` | N/A | Typed wrapper |
| Any other `/addresses/*` | `client.Addresses.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Bridges

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/bridges

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/bridges/tvl` | `client.Bridges.TVL()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/bridges/tvl_relative` | `client.Bridges.TVLRelative()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/bridges/deposits_by_chain` | `client.Bridges.DepositsByChain()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/bridges/net_flow_by_chain` | `client.Bridges.NetFlowByChain()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/bridges/withdrawals_by_chain` | `client.Bridges.WithdrawalsByChain()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| Any other `/bridges/*` | `client.Bridges.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Blockchain

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/blockchain

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/blockchain/block_height` | `client.Blockchain.BlockHeight()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/block_interval_mean` | `client.Blockchain.BlockIntervalMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/block_interval_median` | `client.Blockchain.BlockIntervalMedian()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/block_size_mean` | `client.Blockchain.BlockSizeMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/block_size_total` | `client.Blockchain.BlockSizeTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/blocks_mined` | `client.Blockchain.BlocksMined()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/utxo_count` | `client.Blockchain.UTXOsTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/utxo_created_count` | `client.Blockchain.UTXOsCreated()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/blockchain/utxo_spent_count` | `client.Blockchain.UTXOsSpent()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/blockchain/*` | `client.Blockchain.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Breakdowns

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/breakdowns

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/breakdowns/supply_by_age` | `client.Breakdowns.SupplyByAge()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/breakdowns/supply_by_wallet_size` | `client.Breakdowns.SupplyByWalletSize()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/breakdowns/supply_by_pnl` | `client.Breakdowns.SupplyByPnL()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/breakdowns/sopr_by_age` | `client.Breakdowns.SOPRByAge()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/breakdowns/sopr_by_wallet_size` | `client.Breakdowns.SOPRByWalletSize()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/breakdowns/mvrv_by_age` | `client.Breakdowns.MVRVByAge()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/breakdowns/realized_cap_by_age` | `client.Breakdowns.RealizedCapByAge()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| Any other `/breakdowns/*` | `client.Breakdowns.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### DeFi

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/defi

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/defi/total_value_locked` | `client.DeFi.TVL()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/defi/*` | `client.DeFi.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Derivatives

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/derivatives

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/derivatives/futures_open_interest` | `client.Derivatives.FuturesOpenInterest()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/futures_volume` | `client.Derivatives.FuturesVolume()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/futures_perpetual_funding_rate` | `client.Derivatives.FuturesPerpetualFundingRate()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/futures_estimated_leverage_ratio` | `client.Derivatives.FuturesEstimatedLeverageRatio()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/options_open_interest` | `client.Derivatives.OptionsOpenInterest()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/options_volume_24h` | `client.Derivatives.OptionsVolume24h()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/futures_long_liquidations_total` | `client.Derivatives.FuturesLiquidationsLongTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/derivatives/futures_short_liquidations_total` | `client.Derivatives.FuturesLiquidationsShortTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/derivatives/*` | `client.Derivatives.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Distribution

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/distribution

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/distribution/balance_exchanges` | `client.Distribution.ExchangeBalanceTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/balance_exchanges_percent` | `client.Distribution.ExchangeBalancePercent()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/exchange_net_position_change` | `client.Distribution.ExchangeNetPositionChange()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/miner_balance` | `client.Distribution.MinerBalance()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/miner_net_position_change` | `client.Distribution.MinerNetPositionChange()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/gini` | `client.Distribution.GiniCoefficient()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/herfindahl` | `client.Distribution.HerfindahlIndex()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/distribution/supply_contracts` | `client.Distribution.SupplyInSmartContracts()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/distribution/*` | `client.Distribution.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Entities

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/entities

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/entities/active_count` | `client.Entities.ActiveEntities()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/entities/new_count` | `client.Entities.NewEntities()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/entities/sending_count` | `client.Entities.SendingEntities()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/entities/receiving_count` | `client.Entities.ReceivingEntities()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/entities/net_growth_count` | `client.Entities.EntitiesNetGrowth()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/entities/count` | `client.Entities.NumberWhales()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/entities/*` | `client.Entities.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### ETH 2.0

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/eth2

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/eth2/validators_active` | `client.ETH2.ActiveValidators()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/eth2/staked_total` | `client.ETH2.TotalValueStaked()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/eth2/validators_count` | `client.ETH2.TotalValidators()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/eth2/validator_balance_mean` | `client.ETH2.AverageValidatorBalance()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/eth2/participation_rate` | `client.ETH2.ParticipationRate()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/eth2/issuance_annualized` | `client.ETH2.EstimatedAnnualIssuance()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/eth2/*` | `client.ETH2.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Fees

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/fees

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/fees/volume_sum` | `client.Fees.FeesTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/fees/volume_mean` | `client.Fees.FeesMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/fees/volume_median` | `client.Fees.FeesMedian()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/fees/gas_price_mean` | `client.Fees.GasPriceMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/fees/gas_price_median` | `client.Fees.GasPriceMedian()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/fees/gas_used_sum` | `client.Fees.GasUsedTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/fees/fee_ratio_multiple` | `client.Fees.FeeRatioMultiple()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/fees/*` | `client.Fees.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Global

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/global

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/global/marketcap` | `client.Global.TotalMarketCap()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/global/marketcap_realized` | `client.Global.TotalRealizedCap()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/global/open_interest` | `client.Global.TotalOpenInterest()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/global/sopr_median` | `client.Global.MedianSOPR()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/global/*` | `client.Global.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Indicators

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/indicators

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/indicators/sopr` | `client.Indicators.SOPR()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/cdd` | `client.Indicators.CDD()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/accumulation_trend_score` | `client.Indicators.AccumulationTrendScore()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/nvt` | `client.Indicators.NVT()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/mvrv` | `client.Indicators.MVRV()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/nupl` | `client.Indicators.NUPL()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/realized_profit` | `client.Indicators.RealizedProfit()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/realized_loss` | `client.Indicators.RealizedLoss()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/reserve_risk` | `client.Indicators.ReserveRisk()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/puell_multiple` | `client.Indicators.PuellMultiple()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/s2f` | `client.Indicators.StockToFlow()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/rhodl_ratio` | `client.Indicators.RHODL()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/fear_greed_index` | `client.Indicators.FearGreed()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/liveliness` | `client.Indicators.Liveliness()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/indicators/dormancy` | `client.Indicators.Dormancy()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/indicators/*` | `client.Indicators.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Institutions

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/institutions

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/institutions/holdings_us_spot_etf` | `client.Institutions.USSpotETFBalances()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/institutions/flows_us_spot_etf` | `client.Institutions.USSpotETFNetFlows()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/institutions/volume_us_spot_etf` | `client.Institutions.USSpotETFTradingVolume()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/institutions/price_us_spot_etf` | `client.Institutions.USSpotETFPrice()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/institutions/holdings_purpose_etf` | `client.Institutions.PurposeBitcoinETFHoldings()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/institutions/*` | `client.Institutions.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Lightning

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/lightning

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/lightning/capacity` | `client.Lightning.Capacity()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/lightning/node_count` | `client.Lightning.NumberOfNodes()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/lightning/channel_count` | `client.Lightning.NumberOfChannels()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/lightning/channel_size_mean` | `client.Lightning.ChannelSizeMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/lightning/channel_size_median` | `client.Lightning.ChannelSizeMedian()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/lightning/base_fee_median` | `client.Lightning.BaseFeeMedian()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/lightning/fee_rate_median` | `client.Lightning.FeeRateMedian()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/lightning/*` | `client.Lightning.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Macro

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/macro

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/macro/us_interest_rate` | `client.Macro.USInterestRate()` | `[]TimePoint` | N/A | Typed wrapper |
| `/macro/us_central_bank_balance_sheet` | `client.Macro.USCentralBankBalanceSheet()` | `[]TimePoint` | N/A | Typed wrapper |
| `/macro/us_gdp` | `client.Macro.USGDP()` | `[]TimePoint` | N/A | Typed wrapper |
| `/macro/us_core_inflation_rate` | `client.Macro.USCoreInflationRate()` | `[]TimePoint` | N/A | Typed wrapper |
| `/macro/us_money_supply_m2` | `client.Macro.USMoneySupplyM2()` | `[]TimePoint` | N/A | Typed wrapper |
| `/macro/euro_interest_rate` | `client.Macro.EuroAreaInterestRate()` | `[]TimePoint` | N/A | Typed wrapper |
| Any other `/macro/*` | `client.Macro.GetRaw()` | `json.RawMessage` | N/A | Generic client |

### Market

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/market

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/market/price_usd_close` | `client.Market.Price()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/price_usd_ohlc` | `client.Market.PriceOHLC()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/market/marketcap_usd` | `client.Market.MarketCap()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/mvrv` | `client.Market.MVRV()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/marketcap_realized_usd` | `client.Market.RealizedCap()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/price_realized_usd` | `client.Market.RealizedPrice()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/btc_dominance` | `client.Market.BTCDominance()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/volume` | `client.Market.SpotVolume()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/volume_usd_24h` | `client.Market.SpotVolume24h()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/price_drawdown_relative` | `client.Market.PriceDrawdownFromATH()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/realized_volatility_1m` | `client.Market.RealizedVolatility1M()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/market/beta_btc_7d` | `client.Market.BetaBtc7D()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/market/*` | `client.Market.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Mempool

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/mempool

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/mempool/count` | `client.Mempool.TransactionCount()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mempool/fees_sum` | `client.Mempool.TotalAmountOfFees()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mempool/value_sum` | `client.Mempool.TotalAmountOfCoins()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mempool/size_sum` | `client.Mempool.TransactionsSize()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mempool/fees_relative_mean` | `client.Mempool.AverageRelativeFee()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/mempool/*` | `client.Mempool.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Mining

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/mining

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/mining/hash_rate_mean` | `client.Mining.HashRate()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mining/difficulty_latest` | `client.Mining.Difficulty()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mining/revenue_sum` | `client.Mining.MinerRevenueTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mining/revenue_from_rewards` | `client.Mining.MinerRevenueBlockRewards()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mining/revenue_from_fees` | `client.Mining.MinerRevenueFees()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/mining/thermocap` | `client.Mining.Thermocap()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/mining/*` | `client.Mining.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Options

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/options

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/options/open_interest_sum` | `client.Options.OpenInterest()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/options/volume_sum` | `client.Options.Volume()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/options/volume_put_call_ratio` | `client.Options.PutCallRatio()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/options/max_pain` | `client.Options.MaxPain()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/options/iv_atm` | `client.Options.ImpliedVolatility()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/options/*` | `client.Options.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Point-In-Time

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/pit

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| Any `/indicators/*_pit` | `client.PointInTime.GetPITTimePoints()` | `[]PITTimePoint` | N/A | PIT wrapper |
| Any `/market/*_pit` | `client.PointInTime.GetPITTimePoints()` | `[]PITTimePoint` | N/A | PIT wrapper |
| Any `*_pit` (object) | `client.PointInTime.GetPITObjectPoints()` | `[]PITObjectPoint` | N/A | PIT wrapper |
| Any `*_pit` (raw) | `client.PointInTime.GetRaw()` | `json.RawMessage` | N/A | Generic client |

### Protocols

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/protocols

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/protocols/aave-v3/tvl_by_token` | `client.Protocols.AaveV3AvailableLiquidity()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| `/protocols/uniswap/tvl_per_fee_tier` | `client.Protocols.UniswapTVLPerFeeTier()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| Any other `/protocols/*` | `client.Protocols.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Signals

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/signals

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/signals/altcoin_cycle_signal` | `client.Signals.AltcoinCycleSignal()` | `[]TimePoint` | N/A | Typed wrapper |
| `/signals/bitcoin_risk_signal` | `client.Signals.BitcoinRiskSignal()` | `[]TimePoint` | N/A | Typed wrapper |
| `/signals/bitcoin_sharpe_signal` | `client.Signals.BitcoinSharpeSignal()` | `[]TimePoint` | N/A | Typed wrapper |
| `/signals/bss_goldilocks_signal` | `client.Signals.BSSGoldilocksSignal()` | `[]TimePoint` | N/A | Typed wrapper |
| `/signals/ecosystem_momentum_signal` | `client.Signals.EcosystemMomentumSignal()` | `[]TimePoint` | N/A | Typed wrapper |
| Any other `/signals/*` | `client.Signals.GetRaw()` | `json.RawMessage` | N/A | Generic client |

### Supply

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/supply

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/supply/current` | `client.Supply.CirculatingSupply()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/supply/active` | `client.Supply.ActiveSupply()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/supply/lth` | `client.Supply.LongTermHolderSupply()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/supply/sth` | `client.Supply.ShortTermHolderSupply()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/supply/inflation_rate` | `client.Supply.InflationRate()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/supply/profit_relative` | `client.Supply.ProfitableSupply()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/supply/hodl_waves` | `client.Supply.HODLWaves()` | `[]ObjectPoint` | Check metadata | Typed wrapper |
| Any other `/supply/*` | `client.Supply.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Transactions

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/transactions

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/transactions/count` | `client.Transactions.Count()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/rate` | `client.Transactions.Rate()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/size_mean` | `client.Transactions.SizeMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/size_sum` | `client.Transactions.SizeTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/volume_mean` | `client.Transactions.TransferVolumeMean()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/volume_sum` | `client.Transactions.TransferVolumeTotal()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/deposits_count` | `client.Transactions.ExchangeDeposits()` | `[]TimePoint` | Check metadata | Typed wrapper |
| `/transactions/withdrawals_count` | `client.Transactions.ExchangeWithdrawals()` | `[]TimePoint` | Check metadata | Typed wrapper |
| Any other `/transactions/*` | `client.Transactions.GetRaw()` | `json.RawMessage` | Check metadata | Generic client |

### Treasuries

**Documentation:** https://docs.glassnode.com/basic-api/endpoints/treasuries

| Metric Path | SDK Method | Response Model | Bulk | Status |
|---|---|---|---|---|
| `/treasuries/balance_companies` | `client.Treasuries.TreasuryBalancesCompanies()` | `[]TimePoint` | N/A | Typed wrapper |
| `/treasuries/balance_governments` | `client.Treasuries.TreasuryBalancesGovernments()` | `[]TimePoint` | N/A | Typed wrapper |
| `/treasuries/flows_companies` | `client.Treasuries.TreasuryNetFlowsCompanies()` | `[]TimePoint` | N/A | Typed wrapper |
| `/treasuries/flows_governments` | `client.Treasuries.TreasuryNetFlowsGovernments()` | `[]TimePoint` | N/A | Typed wrapper |
| `/treasuries/count_companies` | `client.Treasuries.TreasuryCountCompanies()` | `[]TimePoint` | N/A | Typed wrapper |
| Any other `/treasuries/*` | `client.Treasuries.GetRaw()` | `json.RawMessage` | N/A | Generic client |

## Bulk Support

Bulk endpoints are available at `/v1/metrics/<path>/bulk` for metrics where `metadata/metric` reports `bulk_supported: true`. Use `client.Metrics.GetBulk()` with a `BulkQuery` to fetch data for multiple assets, exchanges, or networks in a single request.

Bulk endpoints are part of the Professional API and are not available on the Advanced/Light API tier.

**Timerange constraints for bulk requests:**
- `10m` and `1h` resolutions: 10 days
- `1d` (24h) resolution: 31 days
- `1w` and `1month` resolutions: 93 days

**Credit consumption:** Each bulk request costs as many credits as individual calls would. BTC costs 1 credit per request; all other assets cost 2 credits. Always specify assets explicitly.

## Point-In-Time Metrics

Point-in-Time (PIT) metrics include a `computed_at` timestamp indicating when the data was computed. Use `client.PointInTime.GetPITTimePoints()` or `GetPITObjectPoints()` to retrieve PIT data with the `computed_at` field preserved.
