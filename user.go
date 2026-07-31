package glassnode

import "context"

// UserService provides access to the Glassnode user endpoints.
type UserService struct {
	client *Client
}

// APIUsage retrieves the current month's API credit usage. The response
// includes a creditsUsed field indicating how many data credits have
// been consumed.
//
// Documentation: https://docs.glassnode.com/basic-api/api-credits
func (s *UserService) APIUsage(ctx context.Context) (*APIUsage, error) {
	var out APIUsage
	_, err := s.client.get(ctx, "/v1/user/api_usage", nil, &out)
	return &out, err
}
