// Package script provides serialization of AWS query requests, and responses.
package script

//go:generate go run -tags codegen ../../../models/protocol_tests/generate.go ../../../models/protocol_tests/input/script.json build_test.go

import (
	"net/url"

	"github.com/alice02/nifcloud-sdk-go/aws/awserr"
	"github.com/alice02/nifcloud-sdk-go/aws/request"
	"github.com/alice02/nifcloud-sdk-go/private/protocol/query/queryutil"
)

// BuildHandler is a named request handler for building script protocol requests
var BuildHandler = request.NamedHandler{Name: "awssdk.script.Build", Fn: Build}

// Build builds a request for an AWS Query service.
func Build(r *request.Request) {
	body := url.Values{}
	if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = awserr.New("SerializationError", "failed encoding Script request", err)
		return
	}

	if !r.IsPresigned() {
		r.HTTPRequest.Method = "POST"
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.HTTPRequest.Header.Set("X-Amz-Target", r.ClientInfo.APIVersion+"."+r.Operation.Name)
		r.SetBufferBody([]byte(body.Encode()))
	} else { // This is a pre-signed request
		r.HTTPRequest.Method = "GET"
		r.HTTPRequest.URL.RawQuery = body.Encode()
	}
}
