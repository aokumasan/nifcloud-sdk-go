package endpoints_test

import (
	"fmt"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/endpoints"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/session"
	"github.com/alice02/nifcloud-sdk-go/service/s3"
	"github.com/alice02/nifcloud-sdk-go/service/sqs"
)

func ExampleEnumPartitions() {
	resolver := endpoints.DefaultResolver()
	partitions := resolver.(endpoints.EnumPartitions).Partitions()

	for _, p := range partitions {
		fmt.Println("Regions for", p.ID())
		for id := range p.Regions() {
			fmt.Println("*", id)
		}

		fmt.Println("Services for", p.ID())
		for id := range p.Services() {
			fmt.Println("*", id)
		}
	}
}

func ExampleResolverFunc() {
	myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				URL:           "s3.custom.endpoint.com",
				SigningRegion: "custom-signing-region",
			}, nil
		}

		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	sess := session.Must(session.NewSession(&nifcloud.Config{
		Region:           nifcloud.String("us-west-2"),
		EndpointResolver: endpoints.ResolverFunc(myCustomResolver),
	}))

	// Create the S3 service client with the shared session. This will
	// automatically use the S3 custom endpoint configured in the custom
	// endpoint resolver wrapping the default endpoint resolver.
	s3Svc := s3.New(sess)
	// Operation calls will be made to the custom endpoint.
	s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: nifcloud.String("myBucket"),
		Key:    nifcloud.String("myObjectKey"),
	})

	// Create the SQS service client with the shared session. This will
	// fallback to the default endpoint resolver because the customization
	// passes any non S3 service endpoint resolve to the default resolver.
	sqsSvc := sqs.New(sess)
	// Operation calls will be made to the default endpoint for SQS for the
	// region configured.
	sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: nifcloud.String("my-queue-url"),
	})
}
