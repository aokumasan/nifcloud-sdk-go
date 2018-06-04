// +build integration

//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/storagegateway"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@storagegateway", func() {
		gucumber.World["client"] = storagegateway.New(smoke.Session)
	})
}
