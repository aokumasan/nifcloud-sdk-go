// +build bench

package ec2query_test

import (
	"testing"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/request"
	"github.com/alice02/nifcloud-sdk-go/awstesting"
	"github.com/alice02/nifcloud-sdk-go/private/protocol/ec2query"
	"github.com/alice02/nifcloud-sdk-go/service/ec2"
)

func BenchmarkEC2QueryBuild_Complex_ec2AuthorizeSecurityGroupEgress(b *testing.B) {
	params := &ec2.AuthorizeSecurityGroupEgressInput{
		GroupId:  nifcloud.String("String"), // Required
		CidrIp:   nifcloud.String("String"),
		DryRun:   nifcloud.Bool(true),
		FromPort: nifcloud.Int64(1),
		IpPermissions: []*ec2.IpPermission{
			{ // Required
				FromPort:   nifcloud.Int64(1),
				IpProtocol: nifcloud.String("String"),
				IpRanges: []*ec2.IpRange{
					{ // Required
						CidrIp: nifcloud.String("String"),
					},
					// More values...
				},
				PrefixListIds: []*ec2.PrefixListId{
					{ // Required
						PrefixListId: nifcloud.String("String"),
					},
					// More values...
				},
				ToPort: nifcloud.Int64(1),
				UserIdGroupPairs: []*ec2.UserIdGroupPair{
					{ // Required
						GroupId:   nifcloud.String("String"),
						GroupName: nifcloud.String("String"),
						UserId:    nifcloud.String("String"),
					},
					// More values...
				},
			},
			// More values...
		},
		IpProtocol:                 nifcloud.String("String"),
		SourceSecurityGroupName:    nifcloud.String("String"),
		SourceSecurityGroupOwnerId: nifcloud.String("String"),
		ToPort: nifcloud.Int64(1),
	}

	benchEC2QueryBuild(b, "AuthorizeSecurityGroupEgress", params)
}

func BenchmarkEC2QueryBuild_Simple_ec2AttachNetworkInterface(b *testing.B) {
	params := &ec2.AttachNetworkInterfaceInput{
		DeviceIndex:        nifcloud.Int64(1),         // Required
		InstanceId:         nifcloud.String("String"), // Required
		NetworkInterfaceId: nifcloud.String("String"), // Required
		DryRun:             nifcloud.Bool(true),
	}

	benchEC2QueryBuild(b, "AttachNetworkInterface", params)
}

func benchEC2QueryBuild(b *testing.B, opName string, params interface{}) {
	svc := awstesting.NewClient()
	svc.ServiceName = "ec2"
	svc.APIVersion = "2015-04-15"

	for i := 0; i < b.N; i++ {
		r := svc.NewRequest(&request.Operation{
			Name:       opName,
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}, params, nil)
		ec2query.Build(r)
		if r.Error != nil {
			b.Fatal("Unexpected error", r.Error)
		}
	}
}
