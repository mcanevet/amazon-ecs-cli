// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/internal/protocol/restxml"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Route53 is a client for Route 53.
type Route53 struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new Route53 client.
func New(config *aws.Config) *Route53 {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "route53",
			APIVersion:  "2013-04-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restxml.Build)
	service.Handlers.Unmarshal.PushBack(restxml.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restxml.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restxml.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &Route53{service}
}

// newRequest creates a new request for a Route53 operation and runs any
// custom request initialization.
func (c *Route53) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
