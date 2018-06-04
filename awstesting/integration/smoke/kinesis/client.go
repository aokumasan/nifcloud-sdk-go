// +build integration

//Package kinesis provides gucumber integration tests support.
package kinesis

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/kinesis"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@kinesis", func() {
		gucumber.World["client"] = kinesis.New(smoke.Session)
	})
}
