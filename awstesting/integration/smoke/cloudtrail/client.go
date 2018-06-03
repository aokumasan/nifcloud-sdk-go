// +build integration

//Package cloudtrail provides gucumber integration tests support.
package cloudtrail

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudtrail"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudtrail", func() {
		gucumber.World["client"] = cloudtrail.New(smoke.Session)
	})
}
