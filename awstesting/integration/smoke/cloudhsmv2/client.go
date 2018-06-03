// +build integration

//Package cloudhsmv2 provides gucumber integration tests support.
package cloudhsmv2

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudhsmv2"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudhsmv2", func() {
		gucumber.World["client"] = cloudhsmv2.New(smoke.Session)
	})
}
