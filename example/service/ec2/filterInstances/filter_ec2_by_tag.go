// +build example

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/session"
	"github.com/alice02/nifcloud-sdk-go/service/ec2"
)

// This example will list instances with a filter
//
// Usage:
// filter_ec2_by_tag <name_filter>
func main() {
	sess := session.Must(session.NewSession())

	nameFilter := os.Args[1]
	awsRegion := "us-east-1"
	svc := ec2.New(sess, &nifcloud.Config{Region: nifcloud.String(awsRegion)})
	fmt.Printf("listing instances with tag %v in: %v\n", nameFilter, awsRegion)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: nifcloud.String("tag:Name"),
				Values: []*string{
					nifcloud.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", awsRegion, err.Error())
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v\n", *resp)
}
