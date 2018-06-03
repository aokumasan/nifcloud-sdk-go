// +build integration

//Package iam provides gucumber integration tests support.
package iam

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/iam"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@iam", func() {
		gucumber.World["client"] = iam.New(smoke.Session)
	})
}
