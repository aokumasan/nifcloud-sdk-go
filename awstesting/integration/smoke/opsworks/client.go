// +build integration

//Package opsworks provides gucumber integration tests support.
package opsworks

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/opsworks"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@opsworks", func() {
		gucumber.World["client"] = opsworks.New(smoke.Session)
	})
}
