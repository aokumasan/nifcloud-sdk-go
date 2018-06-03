// +build integration

//Package cloudfront provides gucumber integration tests support.
package cloudfront

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/cloudfront"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudfront", func() {
		gucumber.World["client"] = cloudfront.New(smoke.Session)
	})
}
