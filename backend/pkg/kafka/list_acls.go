package kafka

import (
	"context"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// ListACLs sends a DescribeACL request for one or more specific filters
//
// Kafka Request documentation:
// DescribeACLsRequest describes ACLs. Describing ACLs works on a filter basis:
// anything that matches the filter is described. Note that there are two
// "types" of filters in this request: the resource filter and the entry
// filter, with entries corresponding to users. The first three fields form the
// resource filter, the last four the entry filter.
func (s *Service) ListACLs(ctx context.Context, req kmsg.DescribeACLsRequest) (*kmsg.DescribeACLsResponse, error) {
	return req.RequestWith(ctx, s.KafkaClient)
}
