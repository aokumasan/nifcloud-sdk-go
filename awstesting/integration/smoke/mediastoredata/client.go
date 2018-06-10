// +build integration

//Package mediastoredata provides gucumber integration tests support.
package mediastoredata

import (
	"fmt"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/mediastore"
	"github.com/alice02/nifcloud-sdk-go/service/mediastoredata"
	"github.com/gucumber/gucumber"
)

func init() {
	const containerName = "awsgosdkteamintegcontainer"

	gucumber.Before("@mediastoredata", func() {
		mediastoreSvc := mediastore.New(smoke.Session)

		resp, err := mediastoreSvc.DescribeContainer(&mediastore.DescribeContainerInput{
			ContainerName: nifcloud.String(containerName),
		})
		if err != nil {
			gucumber.World["error"] = fmt.Errorf("failed to get mediastore container endpoint for test, %v", err)
			return
		}

		gucumber.World["client"] = mediastoredata.New(smoke.Session, &nifcloud.Config{
			Endpoint: resp.Container.Endpoint,
		})
	})
}
