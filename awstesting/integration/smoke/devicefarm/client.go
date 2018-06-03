// +build integration

//Package devicefarm provides gucumber integration tests support.
package devicefarm

import (
	"github.com/alice02/nifcloud-sdk-go/aws"
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/devicefarm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@devicefarm", func() {
		// FIXME remove custom region
		gucumber.World["client"] = devicefarm.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
