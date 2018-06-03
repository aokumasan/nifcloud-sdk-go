// +build codegen

package api

import (
	"bytes"
	"fmt"
)

type wafregionalExamplesBuilder struct {
	defaultExamplesBuilder
}

func (builder wafregionalExamplesBuilder) Imports(a *API) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`"fmt"
	"strings"
	"time"

	"github.com/alice02/nifcloud-sdk-go/aws"
	"github.com/alice02/nifcloud-sdk-go/aws/awserr"
	"github.com/alice02/nifcloud-sdk-go/aws/session"
	"github.com/alice02/nifcloud-sdk-go/service/waf"
	`)

	buf.WriteString(fmt.Sprintf("\"%s/%s\"", "github.com/alice02/nifcloud-sdk-go/service", a.PackageName()))
	return buf.String()
}
