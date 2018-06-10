// +build bench

package restxml_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"bytes"
	"encoding/xml"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/credentials"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/endpoints"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/request"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/session"
	"github.com/alice02/nifcloud-sdk-go/private/protocol/restxml"
	"github.com/alice02/nifcloud-sdk-go/service/cloudfront"
	"github.com/alice02/nifcloud-sdk-go/service/s3"
)

var (
	cloudfrontSvc *cloudfront.CloudFront
	s3Svc         *s3.S3
)

func TestMain(m *testing.M) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	sess := session.Must(session.NewSession(&nifcloud.Config{
		Credentials:      credentials.NewStaticCredentials("Key", "Secret", "Token"),
		Endpoint:         nifcloud.String(server.URL),
		S3ForcePathStyle: nifcloud.Bool(true),
		DisableSSL:       nifcloud.Bool(true),
		Region:           nifcloud.String(endpoints.UsWest2RegionID),
	}))
	cloudfrontSvc = cloudfront.New(sess)
	s3Svc = s3.New(sess)

	c := m.Run()
	server.Close()
	os.Exit(c)
}

func BenchmarkRESTXMLBuild_Complex_CFCreateDistro(b *testing.B) {
	params := cloudfrontCreateDistributionInput()

	benchRESTXMLBuild(b, func() *request.Request {
		req, _ := cloudfrontSvc.CreateDistributionRequest(params)
		return req
	})
}

func BenchmarkRESTXMLBuild_Simple_CFDeleteDistro(b *testing.B) {
	params := cloudfrontDeleteDistributionInput()

	benchRESTXMLBuild(b, func() *request.Request {
		req, _ := cloudfrontSvc.DeleteDistributionRequest(params)
		return req
	})
}

func BenchmarkRESTXMLBuild_REST_S3HeadObject(b *testing.B) {
	params := s3HeadObjectInput()

	benchRESTXMLBuild(b, func() *request.Request {
		req, _ := s3Svc.HeadObjectRequest(params)
		return req
	})
}

func BenchmarkRESTXMLBuild_XML_S3PutObjectAcl(b *testing.B) {
	params := s3PutObjectAclInput()

	benchRESTXMLBuild(b, func() *request.Request {
		req, _ := s3Svc.PutObjectAclRequest(params)
		return req
	})
}

func BenchmarkRESTXMLRequest_Complex_CFCreateDistro(b *testing.B) {
	benchRESTXMLRequest(b, func() *request.Request {
		req, _ := cloudfrontSvc.CreateDistributionRequest(cloudfrontCreateDistributionInput())
		return req
	})
}

func BenchmarkRESTXMLRequest_Simple_CFDeleteDistro(b *testing.B) {
	benchRESTXMLRequest(b, func() *request.Request {
		req, _ := cloudfrontSvc.DeleteDistributionRequest(cloudfrontDeleteDistributionInput())
		return req
	})
}

func BenchmarkRESTXMLRequest_REST_S3HeadObject(b *testing.B) {
	benchRESTXMLRequest(b, func() *request.Request {
		req, _ := s3Svc.HeadObjectRequest(s3HeadObjectInput())
		return req
	})
}

func BenchmarkRESTXMLRequest_XML_S3PutObjectAcl(b *testing.B) {
	benchRESTXMLRequest(b, func() *request.Request {
		req, _ := s3Svc.PutObjectAclRequest(s3PutObjectAclInput())
		return req
	})
}

func BenchmarkEncodingXML_Simple(b *testing.B) {
	params := cloudfrontDeleteDistributionInput()

	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		encoder := xml.NewEncoder(buf)
		if err := encoder.Encode(params); err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func benchRESTXMLBuild(b *testing.B, reqFn func() *request.Request) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := reqFn()
		restxml.Build(req)
		if req.Error != nil {
			b.Fatal("Unexpected error", req.Error)
		}
	}
}

func benchRESTXMLRequest(b *testing.B, reqFn func() *request.Request) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reqFn().Send()
		if err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func cloudfrontCreateDistributionInput() *cloudfront.CreateDistributionInput {
	return &cloudfront.CreateDistributionInput{
		DistributionConfig: &cloudfront.DistributionConfig{ // Required
			CallerReference: nifcloud.String("string"), // Required
			Comment:         nifcloud.String("string"), // Required
			DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{ // Required
				ForwardedValues: &cloudfront.ForwardedValues{ // Required
					Cookies: &cloudfront.CookiePreference{ // Required
						Forward: nifcloud.String("ItemSelection"), // Required
						WhitelistedNames: &cloudfront.CookieNames{
							Quantity: nifcloud.Int64(1), // Required
							Items: []*string{
								nifcloud.String("string"), // Required
								// More values...
							},
						},
					},
					QueryString: nifcloud.Bool(true), // Required
					Headers: &cloudfront.Headers{
						Quantity: nifcloud.Int64(1), // Required
						Items: []*string{
							nifcloud.String("string"), // Required
							// More values...
						},
					},
				},
				MinTTL:         nifcloud.Int64(1),         // Required
				TargetOriginId: nifcloud.String("string"), // Required
				TrustedSigners: &cloudfront.TrustedSigners{ // Required
					Enabled:  nifcloud.Bool(true), // Required
					Quantity: nifcloud.Int64(1),   // Required
					Items: []*string{
						nifcloud.String("string"), // Required
						// More values...
					},
				},
				ViewerProtocolPolicy: nifcloud.String("ViewerProtocolPolicy"), // Required
				AllowedMethods: &cloudfront.AllowedMethods{
					Items: []*string{ // Required
						nifcloud.String("Method"), // Required
						// More values...
					},
					Quantity: nifcloud.Int64(1), // Required
					CachedMethods: &cloudfront.CachedMethods{
						Items: []*string{ // Required
							nifcloud.String("Method"), // Required
							// More values...
						},
						Quantity: nifcloud.Int64(1), // Required
					},
				},
				DefaultTTL:      nifcloud.Int64(1),
				MaxTTL:          nifcloud.Int64(1),
				SmoothStreaming: nifcloud.Bool(true),
			},
			Enabled: nifcloud.Bool(true), // Required
			Origins: &cloudfront.Origins{ // Required
				Quantity: nifcloud.Int64(1), // Required
				Items: []*cloudfront.Origin{
					{ // Required
						DomainName: nifcloud.String("string"), // Required
						Id:         nifcloud.String("string"), // Required
						CustomOriginConfig: &cloudfront.CustomOriginConfig{
							HTTPPort:             nifcloud.Int64(1),                       // Required
							HTTPSPort:            nifcloud.Int64(1),                       // Required
							OriginProtocolPolicy: nifcloud.String("OriginProtocolPolicy"), // Required
						},
						OriginPath: nifcloud.String("string"),
						S3OriginConfig: &cloudfront.S3OriginConfig{
							OriginAccessIdentity: nifcloud.String("string"), // Required
						},
					},
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: nifcloud.Int64(1), // Required
				Items: []*string{
					nifcloud.String("string"), // Required
					// More values...
				},
			},
			CacheBehaviors: &cloudfront.CacheBehaviors{
				Quantity: nifcloud.Int64(1), // Required
				Items: []*cloudfront.CacheBehavior{
					{ // Required
						ForwardedValues: &cloudfront.ForwardedValues{ // Required
							Cookies: &cloudfront.CookiePreference{ // Required
								Forward: nifcloud.String("ItemSelection"), // Required
								WhitelistedNames: &cloudfront.CookieNames{
									Quantity: nifcloud.Int64(1), // Required
									Items: []*string{
										nifcloud.String("string"), // Required
										// More values...
									},
								},
							},
							QueryString: nifcloud.Bool(true), // Required
							Headers: &cloudfront.Headers{
								Quantity: nifcloud.Int64(1), // Required
								Items: []*string{
									nifcloud.String("string"), // Required
									// More values...
								},
							},
						},
						MinTTL:         nifcloud.Int64(1),         // Required
						PathPattern:    nifcloud.String("string"), // Required
						TargetOriginId: nifcloud.String("string"), // Required
						TrustedSigners: &cloudfront.TrustedSigners{ // Required
							Enabled:  nifcloud.Bool(true), // Required
							Quantity: nifcloud.Int64(1),   // Required
							Items: []*string{
								nifcloud.String("string"), // Required
								// More values...
							},
						},
						ViewerProtocolPolicy: nifcloud.String("ViewerProtocolPolicy"), // Required
						AllowedMethods: &cloudfront.AllowedMethods{
							Items: []*string{ // Required
								nifcloud.String("Method"), // Required
								// More values...
							},
							Quantity: nifcloud.Int64(1), // Required
							CachedMethods: &cloudfront.CachedMethods{
								Items: []*string{ // Required
									nifcloud.String("Method"), // Required
									// More values...
								},
								Quantity: nifcloud.Int64(1), // Required
							},
						},
						DefaultTTL:      nifcloud.Int64(1),
						MaxTTL:          nifcloud.Int64(1),
						SmoothStreaming: nifcloud.Bool(true),
					},
					// More values...
				},
			},
			CustomErrorResponses: &cloudfront.CustomErrorResponses{
				Quantity: nifcloud.Int64(1), // Required
				Items: []*cloudfront.CustomErrorResponse{
					{ // Required
						ErrorCode:          nifcloud.Int64(1), // Required
						ErrorCachingMinTTL: nifcloud.Int64(1),
						ResponseCode:       nifcloud.String("string"),
						ResponsePagePath:   nifcloud.String("string"),
					},
					// More values...
				},
			},
			DefaultRootObject: nifcloud.String("string"),
			Logging: &cloudfront.LoggingConfig{
				Bucket:         nifcloud.String("string"), // Required
				Enabled:        nifcloud.Bool(true),       // Required
				IncludeCookies: nifcloud.Bool(true),       // Required
				Prefix:         nifcloud.String("string"), // Required
			},
			PriceClass: nifcloud.String("PriceClass"),
			Restrictions: &cloudfront.Restrictions{
				GeoRestriction: &cloudfront.GeoRestriction{ // Required
					Quantity:        nifcloud.Int64(1),                     // Required
					RestrictionType: nifcloud.String("GeoRestrictionType"), // Required
					Items: []*string{
						nifcloud.String("string"), // Required
						// More values...
					},
				},
			},
			ViewerCertificate: &cloudfront.ViewerCertificate{
				CloudFrontDefaultCertificate: nifcloud.Bool(true),
				IAMCertificateId:             nifcloud.String("string"),
				MinimumProtocolVersion:       nifcloud.String("MinimumProtocolVersion"),
				SSLSupportMethod:             nifcloud.String("SSLSupportMethod"),
			},
		},
	}
}

func cloudfrontDeleteDistributionInput() *cloudfront.DeleteDistributionInput {
	return &cloudfront.DeleteDistributionInput{
		Id:      nifcloud.String("string"), // Required
		IfMatch: nifcloud.String("string"),
	}
}

func s3HeadObjectInput() *s3.HeadObjectInput {
	return &s3.HeadObjectInput{
		Bucket:    nifcloud.String("somebucketname"),
		Key:       nifcloud.String("keyname"),
		VersionId: nifcloud.String("someVersion"),
		IfMatch:   nifcloud.String("IfMatch"),
	}
}

func s3PutObjectAclInput() *s3.PutObjectAclInput {
	return &s3.PutObjectAclInput{
		Bucket: nifcloud.String("somebucketname"),
		Key:    nifcloud.String("keyname"),
		AccessControlPolicy: &s3.AccessControlPolicy{
			Grants: []*s3.Grant{
				{
					Grantee: &s3.Grantee{
						DisplayName:  nifcloud.String("someName"),
						EmailAddress: nifcloud.String("someAddr"),
						ID:           nifcloud.String("someID"),
						Type:         nifcloud.String(s3.TypeCanonicalUser),
						URI:          nifcloud.String("someURI"),
					},
					Permission: nifcloud.String(s3.PermissionWrite),
				},
			},
			Owner: &s3.Owner{
				DisplayName: nifcloud.String("howdy"),
				ID:          nifcloud.String("someID"),
			},
		},
	}
}
