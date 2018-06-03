// +build integration

//Package route53domains provides gucumber integration tests support.
package route53domains

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/route53domains"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@route53domains", func() {
		gucumber.World["client"] = route53domains.New(smoke.Session)
	})
}
