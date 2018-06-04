// +build integration

//Package autoscalingplans provides gucumber integration tests support.
package autoscalingplans

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/autoscalingplans"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@autoscalingplans", func() {
		gucumber.World["client"] = autoscalingplans.New(smoke.Session)
	})
}
