// +build integration

//Package mediastore provides gucumber integration tests support.
package mediastore

import (
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/mediastore"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@mediastore", func() {
		gucumber.World["client"] = mediastore.New(smoke.Session)
	})
}
