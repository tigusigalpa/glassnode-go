// Package glassnode provides a Go SDK for the Glassnode Basic API
// (https://api.glassnode.com). It exposes service structs grouped by
// resource category accessible as fields on Client.
//
// The SDK supports both header (X-Api-Key) and query-string (api_key)
// authentication, with header mode as the secure default. All
// network-facing methods accept context.Context as their first parameter.
package glassnode

import (
	"net/http"
	"time"
)

// DefaultBaseURL is the default Glassnode API base URL used by
// NewClient unless overridden with WithBaseURL.
const DefaultBaseURL = "https://api.glassnode.com"

// DefaultTimeout is the default HTTP client timeout used unless
// overridden with WithTimeout.
const DefaultTimeout = 30 * time.Second

// DefaultMaxAttempts is the default number of attempts (including the
// initial request) made for a request before giving up when receiving
// HTTP 429 responses, unless overridden with WithRetry.
const DefaultMaxAttempts = 1

// DefaultBaseDelay is the default initial backoff delay used for retry
// attempts, unless overridden with WithRetry.
const DefaultBaseDelay = 500 * time.Millisecond

// DefaultUserAgent is the default User-Agent header value sent with
// every request.
const DefaultUserAgent = "glassnode-go/1.0.0"

// AuthMode controls how the API key is transmitted.
type AuthMode int

const (
	// AuthModeHeader sends the API key via the X-Api-Key header.
	// This is the default and recommended mode.
	AuthModeHeader AuthMode = iota
	// AuthModeQuery sends the API key via the api_key query parameter.
	// This is less secure because the key may appear in logs and URLs.
	AuthModeQuery
)

// Client is the entry point of the Glassnode Go SDK. It holds the HTTP
// configuration and exposes one service struct per API resource group.
// A Client is safe for concurrent use by multiple goroutines.
type Client struct {
	apiKey      string
	authMode    AuthMode
	baseURL     string
	httpClient  *http.Client
	maxAttempts int
	baseDelay   time.Duration
	userAgent   string
	appID       string

	// Metadata provides access to the metadata endpoints.
	Metadata *MetadataService
	// User provides access to the user endpoints.
	User *UserService
	// Metrics provides access to the generic metric endpoints.
	Metrics *MetricsService
	// Addresses provides access to the addresses category endpoints.
	Addresses *AddressesService
	// Bridges provides access to the bridges category endpoints.
	Bridges *BridgesService
	// Blockchain provides access to the blockchain category endpoints.
	Blockchain *BlockchainService
	// Breakdowns provides access to the breakdowns category endpoints.
	Breakdowns *BreakdownsService
	// DeFi provides access to the DeFi category endpoints.
	DeFi *DeFiService
	// Derivatives provides access to the derivatives category endpoints.
	Derivatives *DerivativesService
	// Distribution provides access to the distribution category endpoints.
	Distribution *DistributionService
	// Entities provides access to the entities category endpoints.
	Entities *EntitiesService
	// ETH2 provides access to the ETH 2.0 category endpoints.
	ETH2 *ETH2Service
	// Fees provides access to the fees category endpoints.
	Fees *FeesService
	// Global provides access to the global category endpoints.
	Global *GlobalService
	// Indicators provides access to the indicators category endpoints.
	Indicators *IndicatorsService
	// Institutions provides access to the institutions category endpoints.
	Institutions *InstitutionsService
	// Lightning provides access to the lightning category endpoints.
	Lightning *LightningService
	// Macro provides access to the macro category endpoints.
	Macro *MacroService
	// Market provides access to the market category endpoints.
	Market *MarketService
	// Mempool provides access to the mempool category endpoints.
	Mempool *MempoolService
	// Mining provides access to the mining category endpoints.
	Mining *MiningService
	// Options provides access to the options category endpoints.
	Options *OptionsService
	// PointInTime provides access to the point-in-time category endpoints.
	PointInTime *PointInTimeService
	// Protocols provides access to the protocols category endpoints.
	Protocols *ProtocolsService
	// Signals provides access to the signals category endpoints.
	Signals *SignalsService
	// Supply provides access to the supply category endpoints.
	Supply *SupplyService
	// Transactions provides access to the transactions category endpoints.
	Transactions *TransactionsService
	// Treasuries provides access to the treasuries category endpoints.
	Treasuries *TreasuriesService
}

// NewClient creates a new Glassnode API client authenticated with the
// given apiKey. The apiKey is sent in the X-Api-Key header on every
// request by default. Behaviour can be customized via Option values such
// as WithBaseURL, WithHTTPClient, WithTimeout, WithRetry, and
// WithAuthMode.
//
// Example:
//
//	client := glassnode.NewClient("YOUR_API_KEY",
//	    glassnode.WithTimeout(15*time.Second),
//	    glassnode.WithRetry(3, time.Second),
//	)
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:      apiKey,
		authMode:    AuthModeHeader,
		baseURL:     DefaultBaseURL,
		httpClient:  &http.Client{Timeout: DefaultTimeout},
		maxAttempts: DefaultMaxAttempts,
		baseDelay:   DefaultBaseDelay,
		userAgent:   DefaultUserAgent,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.initServices()

	return c
}

// NewClientFromEnv reads the API key from the GLASSNODE_API_KEY
// environment variable and returns a new Glassnode client. Additional
// Option values can be passed to override timeouts, retries, etc.
func NewClientFromEnv(opts ...Option) (*Client, error) {
	apiKey := ""
	if v, ok := lookupEnv("GLASSNODE_API_KEY"); ok {
		apiKey = v
	}
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	return NewClient(apiKey, opts...), nil
}

// initServices creates all service structs and wires them to the client.
func (c *Client) initServices() {
	c.Metadata = &MetadataService{client: c}
	c.User = &UserService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.Addresses = &AddressesService{client: c}
	c.Bridges = &BridgesService{client: c}
	c.Blockchain = &BlockchainService{client: c}
	c.Breakdowns = &BreakdownsService{client: c}
	c.DeFi = &DeFiService{client: c}
	c.Derivatives = &DerivativesService{client: c}
	c.Distribution = &DistributionService{client: c}
	c.Entities = &EntitiesService{client: c}
	c.ETH2 = &ETH2Service{client: c}
	c.Fees = &FeesService{client: c}
	c.Global = &GlobalService{client: c}
	c.Indicators = &IndicatorsService{client: c}
	c.Institutions = &InstitutionsService{client: c}
	c.Lightning = &LightningService{client: c}
	c.Macro = &MacroService{client: c}
	c.Market = &MarketService{client: c}
	c.Mempool = &MempoolService{client: c}
	c.Mining = &MiningService{client: c}
	c.Options = &OptionsService{client: c}
	c.PointInTime = &PointInTimeService{client: c}
	c.Protocols = &ProtocolsService{client: c}
	c.Signals = &SignalsService{client: c}
	c.Supply = &SupplyService{client: c}
	c.Transactions = &TransactionsService{client: c}
	c.Treasuries = &TreasuriesService{client: c}
}

// lookupEnv wraps os.LookupEnv so it can be swapped in tests if needed.
func lookupEnv(key string) (string, bool) {
	return defaultLookupEnv(key)
}

var defaultLookupEnv = func(key string) (string, bool) {
	return "", false
}
