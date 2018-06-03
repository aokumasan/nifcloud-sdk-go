// +build integration

//Package cloudsearch provides gucumber integration tests support.
package cloudsearch

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudsearch"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudsearch", func() {
		gucumber.World["client"] = cloudsearch.New(smoke.Session)
	})
}
