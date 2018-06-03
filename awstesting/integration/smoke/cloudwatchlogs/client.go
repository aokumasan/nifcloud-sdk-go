// +build integration

//Package cloudwatchlogs provides gucumber integration tests support.
package cloudwatchlogs

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudwatchlogs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudwatchlogs", func() {
		gucumber.World["client"] = cloudwatchlogs.New(smoke.Session)
	})
}
