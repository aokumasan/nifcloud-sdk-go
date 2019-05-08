// Code generated by nifcloud/endpoints/v3model_codegen.go. DO NOT EDIT.

package endpoints

import (
	"regexp"
)

// Partition identifiers
const (
	NifcloudPartitionID = "nifcloud" // NIFCLOUD Standard partition.
)

// NIFCLOUD Standard partition's regions.
const (
	JpEast1RegionID = "jp-east-1" // jp-east-1.
	JpEast2RegionID = "jp-east-2" // jp-east-2.
	JpEast3RegionID = "jp-east-3" // jp-east-3.
	JpEast4RegionID = "jp-east-4" // jp-east-4.
	JpWest1RegionID = "jp-west-1" // jp-west-1.
	UsEast1RegionID = "us-east-1" // us-east-1.
)

// Service identifiers
const (
	ComputingServiceID   = "computing"   // Computing.
	Ec2metadataServiceID = "ec2metadata" // Ec2metadata.
	NasServiceID         = "nas"         // Nas.
	RdbServiceID         = "rdb"         // Rdb.
	ScriptServiceID      = "script"      // Script.
)

// DefaultResolver returns an Endpoint resolver that will be able
// to resolve endpoints for: NIFCLOUD Standard.
//
// Use DefaultPartitions() to get the list of the default partitions.
func DefaultResolver() Resolver {
	return defaultPartitions
}

// DefaultPartitions returns a list of the partitions the SDK is bundled
// with. The available partitions are: NIFCLOUD Standard.
//
//    partitions := endpoints.DefaultPartitions
//    for _, p := range partitions {
//        // ... inspect partitions
//    }
func DefaultPartitions() []Partition {
	return defaultPartitions.Partitions()
}

var defaultPartitions = partitions{
	nifcloudPartition,
}

// NifcloudPartition returns the Resolver for NIFCLOUD Standard.
func NifcloudPartition() Partition {
	return nifcloudPartition.Partition()
}

var nifcloudPartition = partition{
	ID:        "nifcloud",
	Name:      "NIFCLOUD Standard",
	DNSSuffix: "api.nifcloud.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(us|jp)\\-\\w+\\-\\d+$")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{region}.{service}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"v4", "v3", "v2-computing"},
	},
	Regions: regions{
		"jp-east-1": region{
			Description: "jp-east-1",
		},
		"jp-east-2": region{
			Description: "jp-east-2",
		},
		"jp-east-3": region{
			Description: "jp-east-3",
		},
		"jp-east-4": region{
			Description: "jp-east-4",
		},
		"jp-west-1": region{
			Description: "jp-west-1",
		},
		"us-east-1": region{
			Description: "us-east-1",
		},
	},
	Services: services{
		"computing": service{
			Defaults: endpoint{
				SignatureVersions: []string{"v2-computing"},
			},
			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
				"jp-east-3": endpoint{},
				"jp-east-4": endpoint{},
				"jp-west-1": endpoint{},
				"us-east-1": endpoint{},
			},
		},
		"ec2metadata": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "169.254.169.254/latest",
					Protocols: []string{"http"},
				},
			},
		},
		"nas": service{

			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
				"jp-east-3": endpoint{},
				"jp-east-4": endpoint{},
				"jp-west-1": endpoint{},
			},
		},
		"rdb": service{

			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
				"jp-east-3": endpoint{},
				"jp-east-4": endpoint{},
				"jp-west-1": endpoint{},
			},
		},
		"script": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "script.api.nifcloud.com",
					CredentialScope: credentialScope{
						Region: "jp-east-1",
					},
				},
			},
		},
	},
}
