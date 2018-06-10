// +build go1.6

package request_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/awserr"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/client"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/client/metadata"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/defaults"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/request"
)

// go version 1.4 and 1.5 do not return an error. Version 1.5 will url encode
// the uri while 1.4 will not
func TestRequestInvalidEndpoint(t *testing.T) {
	endpoint := "http://localhost:90 "

	r := request.New(
		nifcloud.Config{},
		metadata.ClientInfo{Endpoint: endpoint},
		defaults.Handlers(),
		client.DefaultRetryer{},
		&request.Operation{},
		nil,
		nil,
	)

	assert.Error(t, r.Error)
}

type timeoutErr struct {
	error
}

var errTimeout = awserr.New("foo", "bar", &timeoutErr{
	errors.New("net/http: request canceled"),
})

func (e *timeoutErr) Timeout() bool {
	return true
}

func (e *timeoutErr) Temporary() bool {
	return true
}
