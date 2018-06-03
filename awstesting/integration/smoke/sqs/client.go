// +build integration

//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/sqs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sqs", func() {
		gucumber.World["client"] = sqs.New(smoke.Session)
	})
}
