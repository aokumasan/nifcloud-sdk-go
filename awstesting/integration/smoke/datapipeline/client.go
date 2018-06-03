// +build integration

//Package datapipeline provides gucumber integration tests support.
package datapipeline

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/datapipeline"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@datapipeline", func() {
		gucumber.World["client"] = datapipeline.New(smoke.Session)
	})
}
