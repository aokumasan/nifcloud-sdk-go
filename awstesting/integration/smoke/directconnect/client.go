// +build integration

//Package directconnect provides gucumber integration tests support.
package directconnect

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/directconnect"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@directconnect", func() {
		gucumber.World["client"] = directconnect.New(smoke.Session)
	})
}
