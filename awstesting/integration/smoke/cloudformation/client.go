// +build integration

//Package cloudformation provides gucumber integration tests support.
package cloudformation

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudformation"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudformation", func() {
		gucumber.World["client"] = cloudformation.New(smoke.Session)
	})
}
