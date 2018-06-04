// +build integration

//Package sts provides gucumber integration tests support.
package sts

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/sts"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sts", func() {
		gucumber.World["client"] = sts.New(smoke.Session)
	})
}
