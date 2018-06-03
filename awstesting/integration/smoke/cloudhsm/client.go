// +build integration

//Package cloudhsm provides gucumber integration tests support.
package cloudhsm

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudhsm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudhsm", func() {
		gucumber.World["client"] = cloudhsm.New(smoke.Session)
	})
}
