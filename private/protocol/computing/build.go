// Package computing provides serialization of AWS COMPUTING requests and responses.
package computing

//go:generate go run -tags codegen ../../../models/protocol_tests/generate.go ../../../models/protocol_tests/input/computing.json build_test.go

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/alice02/nifcloud-sdk-go/nifcloud/awserr"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/request"
	"github.com/alice02/nifcloud-sdk-go/private/protocol/query/queryutil"
)

// BuildHandler is a named request handler for building computing protocol requests
var BuildHandler = request.NamedHandler{Name: "awssdk.computing.Build", Fn: Build}

// Build builds a request for the COMPUTING protocol.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}
	if err := queryutil.Parse(body, r.Params, true); err != nil {
		r.Error = awserr.New("SerializationError", "failed encoding COMPUTING Query request", err)
	}

	// Fix request parameters of DescribeLoadBalancers for NIFCLOUD
	if r.Operation.Name == "DescribeLoadBalancers" {
		parameterLength := reflect.ValueOf(r.Params).Elem().FieldByName("LoadBalancerNames").Len()
		prefix := "LoadBalancerNames"
		for i := 1; i <= parameterLength; i++ {
			body[fmt.Sprintf("%s.LoadBalancerPort.%d", prefix, i)] = body[fmt.Sprintf("%s.member.%d.LoadBalancerPort", prefix, i)]
			body[fmt.Sprintf("%s.InstancePort.%d", prefix, i)] = body[fmt.Sprintf("%s.member.%d.InstancePort", prefix, i)]
			body[fmt.Sprintf("%s.member.%d", prefix, i)] = body[fmt.Sprintf("%s.member.%d.LoadBalancerName", prefix, i)]
			delete(body, fmt.Sprintf("%s.member.%d.LoadBalancerPort", prefix, i))
			delete(body, fmt.Sprintf("%s.member.%d.InstancePort", prefix, i))
			delete(body, fmt.Sprintf("%s.member.%d.LoadBalancerName", prefix, i))
		}
	}

	if !r.IsPresigned() {
		r.HTTPRequest.Method = "POST"
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else { // This is a pre-signed request
		r.HTTPRequest.Method = "GET"
		r.HTTPRequest.URL.RawQuery = body.Encode()
	}
}
