// +build integration

//Package rds provides gucumber integration tests support.
package rds

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/rds"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@rds", func() {
		gucumber.World["client"] = rds.New(smoke.Session)
	})
}
