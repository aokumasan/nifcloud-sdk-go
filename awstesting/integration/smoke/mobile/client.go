// +build integration

//Package mobile provides gucumber integration tests support.
package mobile

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/mobile"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@mobile", func() {
		gucumber.World["client"] = mobile.New(smoke.Session)
	})
}
