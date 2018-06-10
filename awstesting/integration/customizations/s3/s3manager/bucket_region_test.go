// +build integration

package s3manager

import (
	"testing"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/awstesting/integration"
	"github.com/alice02/nifcloud-sdk-go/service/s3/s3manager"
)

func TestGetBucketRegion(t *testing.T) {
	expectRegion := nifcloud.StringValue(integration.Session.Config.Region)

	ctx := nifcloud.BackgroundContext()
	region, err := s3manager.GetBucketRegion(ctx, integration.Session,
		nifcloud.StringValue(bucketName), expectRegion)

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := expectRegion, region; e != a {
		t.Errorf("expect %s bucket region, got %s", e, a)
	}
}
