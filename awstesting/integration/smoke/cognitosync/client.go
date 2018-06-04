// +build integration

//Package cognitosync provides gucumber integration tests support.
package cognitosync

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cognitosync"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cognitosync", func() {
		gucumber.World["client"] = cognitosync.New(smoke.Session)
	})
}
