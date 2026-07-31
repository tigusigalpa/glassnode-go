package glassnode

import (
	"context"
	"net/url"
)

// MetadataService provides access to the Glassnode metadata endpoints.
// Metadata is a first-class feature: use it to list supported assets,
// enumerate metric paths, inspect parameter capabilities, and determine
// whether bulk endpoints are available for a given metric.
//
// Documentation: https://docs.glassnode.com/basic-api/metadata
type MetadataService struct {
	client *Client
}

// Assets lists all assets and relevant metadata. The optional filter
// parameter accepts a CEL (Common Expression Language) expression to
// narrow results.
//
// Documentation: https://docs.glassnode.com/basic-api/metadata#assets
func (s *MetadataService) Assets(ctx context.Context, filter string) (*AssetsResponse, error) {
	q := url.Values{}
	if filter != "" {
		q.Set("filter", filter)
	}
	var out AssetsResponse
	_, err := s.client.get(ctx, "/v1/metadata/assets", q, &out)
	return &out, err
}

// Metrics lists all available metric paths. The returned paths include
// only the metric portion (e.g. "/addresses/count"), not the full
// "/v1/metrics" prefix. The optional query parameters (a, e, i, etc.)
// can be used to filter the list.
//
// Documentation: https://docs.glassnode.com/basic-api/metadata#metrics
func (s *MetadataService) Metrics(ctx context.Context, query *MetricQuery) ([]string, error) {
	var out []string
	_, err := s.client.get(ctx, "/v1/metadata/metrics", query.ToValues(), &out)
	return out, err
}

// Metric retrieves metadata describing the available parameters and
// potential drill-down options for a specific metric. The path parameter
// is required (e.g. "/distribution/balance_exchanges"). Additional
// query parameters can be passed to refine the metadata query.
//
// Documentation: https://docs.glassnode.com/basic-api/metadata#metric
func (s *MetadataService) Metric(ctx context.Context, metricPath string, query *MetricQuery) (*MetricMetadata, error) {
	if metricPath == "" {
		return nil, ErrInvalidMetricPath
	}
	q := query.ToValues()
	q.Set("path", normalizeMetricPath(metricPath))
	var out MetricMetadata
	_, err := s.client.get(ctx, "/v1/metadata/metric", q, &out)
	return &out, err
}
