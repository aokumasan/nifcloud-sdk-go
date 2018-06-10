// +build integration

//Package iotdataplane provides gucumber integration tests support.
package iotdataplane

import (
	"fmt"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration/smoke"
	"github.com/alice02/nifcloud-sdk-go/service/iot"
	"github.com/alice02/nifcloud-sdk-go/service/iotdataplane"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@iotdataplane", func() {
		svc := iot.New(smoke.Session)
		result, err := svc.DescribeEndpoint(&iot.DescribeEndpointInput{})
		if err != nil {
			gucumber.World["error"] = err
			return
		}

		fmt.Println("IOT Data endpoint:", *result.EndpointAddress)
		gucumber.World["client"] = iotdataplane.New(smoke.Session, nifcloud.NewConfig().
			WithEndpoint(*result.EndpointAddress))
	})
}
