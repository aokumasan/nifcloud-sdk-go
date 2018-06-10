// +build integration

// Package integration performs initialization and validation for integration
// tests.
package integration

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/session"
)

// Session is a shared session for all integration tests to use.
var Session = session.Must(session.NewSession())

func init() {
	logLevel := Session.Config.LogLevel
	if os.Getenv("DEBUG") != "" {
		logLevel = nifcloud.LogLevel(nifcloud.LogDebug)
	}
	if os.Getenv("DEBUG_SIGNING") != "" {
		logLevel = nifcloud.LogLevel(nifcloud.LogDebugWithSigning)
	}
	if os.Getenv("DEBUG_BODY") != "" {
		logLevel = nifcloud.LogLevel(nifcloud.LogDebugWithSigning | nifcloud.LogDebugWithHTTPBody)
	}
	Session.Config.LogLevel = logLevel

	if nifcloud.StringValue(Session.Config.Region) == "" {
		panic("AWS_REGION must be configured to run integration tests")
	}
}

// UniqueID returns a unique UUID-like identifier for use in generating
// resources for integration tests.
func UniqueID() string {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)
	return fmt.Sprintf("%x", uuid)
}
