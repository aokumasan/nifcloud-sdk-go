// +build integration

//Package redshift provides gucumber integration tests support.
package redshift

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/redshift"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@redshift", func() {
		gucumber.World["client"] = redshift.New(smoke.Session)
	})
}
