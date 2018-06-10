package request

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
)

func TestResetBody_WithBodyContents(t *testing.T) {
	r := Request{
		HTTPRequest: &http.Request{},
	}

	reader := strings.NewReader("abc")
	r.Body = reader

	r.ResetBody()

	if v, ok := r.HTTPRequest.Body.(*offsetReader); !ok || v == nil {
		t.Errorf("expected request body to be set to reader, got %#v",
			r.HTTPRequest.Body)
	}
}

type mockReader struct{}

func (mockReader) Read([]byte) (int, error) {
	return 0, io.EOF
}

func TestResetBody_ExcludeEmptyUnseekableBodyByMethod(t *testing.T) {
	cases := []struct {
		Method   string
		Body     io.ReadSeeker
		IsNoBody bool
	}{
		{
			Method:   "GET",
			IsNoBody: true,
			Body:     nifcloud.ReadSeekCloser(mockReader{}),
		},
		{
			Method:   "HEAD",
			IsNoBody: true,
			Body:     nifcloud.ReadSeekCloser(mockReader{}),
		},
		{
			Method:   "DELETE",
			IsNoBody: true,
			Body:     nifcloud.ReadSeekCloser(mockReader{}),
		},
		{
			Method:   "PUT",
			IsNoBody: false,
			Body:     nifcloud.ReadSeekCloser(mockReader{}),
		},
		{
			Method:   "PATCH",
			IsNoBody: false,
			Body:     nifcloud.ReadSeekCloser(mockReader{}),
		},
		{
			Method:   "POST",
			IsNoBody: false,
			Body:     nifcloud.ReadSeekCloser(mockReader{}),
		},
		{
			Method:   "GET",
			IsNoBody: false,
			Body:     nifcloud.ReadSeekCloser(bytes.NewBuffer([]byte("abc"))),
		},
		{
			Method:   "GET",
			IsNoBody: true,
			Body:     nifcloud.ReadSeekCloser(bytes.NewBuffer(nil)),
		},
		{
			Method:   "POST",
			IsNoBody: false,
			Body:     nifcloud.ReadSeekCloser(bytes.NewBuffer([]byte("abc"))),
		},
		{
			Method:   "POST",
			IsNoBody: true,
			Body:     nifcloud.ReadSeekCloser(bytes.NewBuffer(nil)),
		},
	}

	for i, c := range cases {
		r := Request{
			HTTPRequest: &http.Request{},
			Operation: &Operation{
				HTTPMethod: c.Method,
			},
		}
		r.SetReaderBody(c.Body)

		if a, e := r.HTTPRequest.Body == NoBody, c.IsNoBody; a != e {
			t.Errorf("%d, expect body to be set to noBody(%t), but was %t", i, e, a)
		}
	}

}
