package awstesting

import (
	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/client"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/client/metadata"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/defaults"
)

// NewClient creates and initializes a generic service client for testing.
func NewClient(cfgs ...*nifcloud.Config) *client.Client {
	info := metadata.ClientInfo{
		Endpoint:    "http://endpoint",
		SigningName: "",
	}
	def := defaults.Get()
	def.Config.MergeIn(cfgs...)

	if v := nifcloud.StringValue(def.Config.Endpoint); len(v) > 0 {
		info.Endpoint = v
	}

	return client.New(*def.Config, info, def.Handlers)
}
