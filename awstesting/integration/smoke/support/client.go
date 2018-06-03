// +build integration

//Package support provides gucumber integration tests support.
package support

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/support"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@support", func() {
		gucumber.World["client"] = support.New(smoke.Session)
	})
}
