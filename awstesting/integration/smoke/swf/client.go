// +build integration

//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/swf"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@swf", func() {
		gucumber.World["client"] = swf.New(smoke.Session)
	})
}
