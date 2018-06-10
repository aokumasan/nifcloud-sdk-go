// +build bench

package restjson_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alice02/nifcloud-sdk-go/nifcloud"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/credentials"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/endpoints"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/request"
	"github.com/alice02/nifcloud-sdk-go/nifcloud/session"
	"github.com/alice02/nifcloud-sdk-go/private/protocol/restjson"
	"github.com/alice02/nifcloud-sdk-go/service/elastictranscoder"
)

var (
	elastictranscoderSvc *elastictranscoder.ElasticTranscoder
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
	elastictranscoderSvc = elastictranscoder.New(sess)

	c := m.Run()
	server.Close()
	os.Exit(c)
}

func BenchmarkRESTJSONBuild_Complex_ETCCreateJob(b *testing.B) {
	params := elastictranscoderCreateJobInput()

	benchRESTJSONBuild(b, func() *request.Request {
		req, _ := elastictranscoderSvc.CreateJobRequest(params)
		return req
	})
}

func BenchmarkRESTJSONBuild_Simple_ETCListJobsByPipeline(b *testing.B) {
	params := elastictranscoderListJobsByPipeline()

	benchRESTJSONBuild(b, func() *request.Request {
		req, _ := elastictranscoderSvc.ListJobsByPipelineRequest(params)
		return req
	})
}

func BenchmarkRESTJSONRequest_Complex_CFCreateJob(b *testing.B) {
	benchRESTJSONRequest(b, func() *request.Request {
		req, _ := elastictranscoderSvc.CreateJobRequest(elastictranscoderCreateJobInput())
		return req
	})
}

func BenchmarkRESTJSONRequest_Simple_ETCListJobsByPipeline(b *testing.B) {
	benchRESTJSONRequest(b, func() *request.Request {
		req, _ := elastictranscoderSvc.ListJobsByPipelineRequest(elastictranscoderListJobsByPipeline())
		return req
	})
}

func benchRESTJSONBuild(b *testing.B, reqFn func() *request.Request) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := reqFn()
		restjson.Build(req)
		if req.Error != nil {
			b.Fatal("Unexpected error", req.Error)
		}
	}
}

func benchRESTJSONRequest(b *testing.B, reqFn func() *request.Request) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := reqFn().Send()
		if err != nil {
			b.Fatal("Unexpected error", err)
		}
	}
}

func elastictranscoderListJobsByPipeline() *elastictranscoder.ListJobsByPipelineInput {
	return &elastictranscoder.ListJobsByPipelineInput{
		PipelineId: nifcloud.String("Id"), // Required
		Ascending:  nifcloud.String("Ascending"),
		PageToken:  nifcloud.String("Id"),
	}
}

func elastictranscoderCreateJobInput() *elastictranscoder.CreateJobInput {
	return &elastictranscoder.CreateJobInput{
		Input: &elastictranscoder.JobInput{ // Required
			AspectRatio: nifcloud.String("AspectRatio"),
			Container:   nifcloud.String("JobContainer"),
			DetectedProperties: &elastictranscoder.DetectedProperties{
				DurationMillis: nifcloud.Int64(1),
				FileSize:       nifcloud.Int64(1),
				FrameRate:      nifcloud.String("FloatString"),
				Height:         nifcloud.Int64(1),
				Width:          nifcloud.Int64(1),
			},
			Encryption: &elastictranscoder.Encryption{
				InitializationVector: nifcloud.String("ZeroTo255String"),
				Key:                  nifcloud.String("Base64EncodedString"),
				KeyMd5:               nifcloud.String("Base64EncodedString"),
				Mode:                 nifcloud.String("EncryptionMode"),
			},
			FrameRate:  nifcloud.String("FrameRate"),
			Interlaced: nifcloud.String("Interlaced"),
			Key:        nifcloud.String("Key"),
			Resolution: nifcloud.String("Resolution"),
		},
		PipelineId: nifcloud.String("Id"), // Required
		Output: &elastictranscoder.CreateJobOutput{
			AlbumArt: &elastictranscoder.JobAlbumArt{
				Artwork: []*elastictranscoder.Artwork{
					{ // Required
						AlbumArtFormat: nifcloud.String("JpgOrPng"),
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: nifcloud.String("ZeroTo255String"),
							Key:                  nifcloud.String("Base64EncodedString"),
							KeyMd5:               nifcloud.String("Base64EncodedString"),
							Mode:                 nifcloud.String("EncryptionMode"),
						},
						InputKey:      nifcloud.String("WatermarkKey"),
						MaxHeight:     nifcloud.String("DigitsOrAuto"),
						MaxWidth:      nifcloud.String("DigitsOrAuto"),
						PaddingPolicy: nifcloud.String("PaddingPolicy"),
						SizingPolicy:  nifcloud.String("SizingPolicy"),
					},
					// More values...
				},
				MergePolicy: nifcloud.String("MergePolicy"),
			},
			Captions: &elastictranscoder.Captions{
				CaptionFormats: []*elastictranscoder.CaptionFormat{
					{ // Required
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: nifcloud.String("ZeroTo255String"),
							Key:                  nifcloud.String("Base64EncodedString"),
							KeyMd5:               nifcloud.String("Base64EncodedString"),
							Mode:                 nifcloud.String("EncryptionMode"),
						},
						Format:  nifcloud.String("CaptionFormatFormat"),
						Pattern: nifcloud.String("CaptionFormatPattern"),
					},
					// More values...
				},
				CaptionSources: []*elastictranscoder.CaptionSource{
					{ // Required
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: nifcloud.String("ZeroTo255String"),
							Key:                  nifcloud.String("Base64EncodedString"),
							KeyMd5:               nifcloud.String("Base64EncodedString"),
							Mode:                 nifcloud.String("EncryptionMode"),
						},
						Key:        nifcloud.String("Key"),
						Label:      nifcloud.String("Name"),
						Language:   nifcloud.String("Key"),
						TimeOffset: nifcloud.String("TimeOffset"),
					},
					// More values...
				},
				MergePolicy: nifcloud.String("CaptionMergePolicy"),
			},
			Composition: []*elastictranscoder.Clip{
				{ // Required
					TimeSpan: &elastictranscoder.TimeSpan{
						Duration:  nifcloud.String("Time"),
						StartTime: nifcloud.String("Time"),
					},
				},
				// More values...
			},
			Encryption: &elastictranscoder.Encryption{
				InitializationVector: nifcloud.String("ZeroTo255String"),
				Key:                  nifcloud.String("Base64EncodedString"),
				KeyMd5:               nifcloud.String("Base64EncodedString"),
				Mode:                 nifcloud.String("EncryptionMode"),
			},
			Key:             nifcloud.String("Key"),
			PresetId:        nifcloud.String("Id"),
			Rotate:          nifcloud.String("Rotate"),
			SegmentDuration: nifcloud.String("FloatString"),
			ThumbnailEncryption: &elastictranscoder.Encryption{
				InitializationVector: nifcloud.String("ZeroTo255String"),
				Key:                  nifcloud.String("Base64EncodedString"),
				KeyMd5:               nifcloud.String("Base64EncodedString"),
				Mode:                 nifcloud.String("EncryptionMode"),
			},
			ThumbnailPattern: nifcloud.String("ThumbnailPattern"),
			Watermarks: []*elastictranscoder.JobWatermark{
				{ // Required
					Encryption: &elastictranscoder.Encryption{
						InitializationVector: nifcloud.String("ZeroTo255String"),
						Key:                  nifcloud.String("Base64EncodedString"),
						KeyMd5:               nifcloud.String("Base64EncodedString"),
						Mode:                 nifcloud.String("EncryptionMode"),
					},
					InputKey:          nifcloud.String("WatermarkKey"),
					PresetWatermarkId: nifcloud.String("PresetWatermarkId"),
				},
				// More values...
			},
		},
		OutputKeyPrefix: nifcloud.String("Key"),
		Outputs: []*elastictranscoder.CreateJobOutput{
			{ // Required
				AlbumArt: &elastictranscoder.JobAlbumArt{
					Artwork: []*elastictranscoder.Artwork{
						{ // Required
							AlbumArtFormat: nifcloud.String("JpgOrPng"),
							Encryption: &elastictranscoder.Encryption{
								InitializationVector: nifcloud.String("ZeroTo255String"),
								Key:                  nifcloud.String("Base64EncodedString"),
								KeyMd5:               nifcloud.String("Base64EncodedString"),
								Mode:                 nifcloud.String("EncryptionMode"),
							},
							InputKey:      nifcloud.String("WatermarkKey"),
							MaxHeight:     nifcloud.String("DigitsOrAuto"),
							MaxWidth:      nifcloud.String("DigitsOrAuto"),
							PaddingPolicy: nifcloud.String("PaddingPolicy"),
							SizingPolicy:  nifcloud.String("SizingPolicy"),
						},
						// More values...
					},
					MergePolicy: nifcloud.String("MergePolicy"),
				},
				Captions: &elastictranscoder.Captions{
					CaptionFormats: []*elastictranscoder.CaptionFormat{
						{ // Required
							Encryption: &elastictranscoder.Encryption{
								InitializationVector: nifcloud.String("ZeroTo255String"),
								Key:                  nifcloud.String("Base64EncodedString"),
								KeyMd5:               nifcloud.String("Base64EncodedString"),
								Mode:                 nifcloud.String("EncryptionMode"),
							},
							Format:  nifcloud.String("CaptionFormatFormat"),
							Pattern: nifcloud.String("CaptionFormatPattern"),
						},
						// More values...
					},
					CaptionSources: []*elastictranscoder.CaptionSource{
						{ // Required
							Encryption: &elastictranscoder.Encryption{
								InitializationVector: nifcloud.String("ZeroTo255String"),
								Key:                  nifcloud.String("Base64EncodedString"),
								KeyMd5:               nifcloud.String("Base64EncodedString"),
								Mode:                 nifcloud.String("EncryptionMode"),
							},
							Key:        nifcloud.String("Key"),
							Label:      nifcloud.String("Name"),
							Language:   nifcloud.String("Key"),
							TimeOffset: nifcloud.String("TimeOffset"),
						},
						// More values...
					},
					MergePolicy: nifcloud.String("CaptionMergePolicy"),
				},
				Composition: []*elastictranscoder.Clip{
					{ // Required
						TimeSpan: &elastictranscoder.TimeSpan{
							Duration:  nifcloud.String("Time"),
							StartTime: nifcloud.String("Time"),
						},
					},
					// More values...
				},
				Encryption: &elastictranscoder.Encryption{
					InitializationVector: nifcloud.String("ZeroTo255String"),
					Key:                  nifcloud.String("Base64EncodedString"),
					KeyMd5:               nifcloud.String("Base64EncodedString"),
					Mode:                 nifcloud.String("EncryptionMode"),
				},
				Key:             nifcloud.String("Key"),
				PresetId:        nifcloud.String("Id"),
				Rotate:          nifcloud.String("Rotate"),
				SegmentDuration: nifcloud.String("FloatString"),
				ThumbnailEncryption: &elastictranscoder.Encryption{
					InitializationVector: nifcloud.String("ZeroTo255String"),
					Key:                  nifcloud.String("Base64EncodedString"),
					KeyMd5:               nifcloud.String("Base64EncodedString"),
					Mode:                 nifcloud.String("EncryptionMode"),
				},
				ThumbnailPattern: nifcloud.String("ThumbnailPattern"),
				Watermarks: []*elastictranscoder.JobWatermark{
					{ // Required
						Encryption: &elastictranscoder.Encryption{
							InitializationVector: nifcloud.String("ZeroTo255String"),
							Key:                  nifcloud.String("Base64EncodedString"),
							KeyMd5:               nifcloud.String("Base64EncodedString"),
							Mode:                 nifcloud.String("EncryptionMode"),
						},
						InputKey:          nifcloud.String("WatermarkKey"),
						PresetWatermarkId: nifcloud.String("PresetWatermarkId"),
					},
					// More values...
				},
			},
			// More values...
		},
		Playlists: []*elastictranscoder.CreateJobPlaylist{
			{ // Required
				Format: nifcloud.String("PlaylistFormat"),
				HlsContentProtection: &elastictranscoder.HlsContentProtection{
					InitializationVector:  nifcloud.String("ZeroTo255String"),
					Key:                   nifcloud.String("Base64EncodedString"),
					KeyMd5:                nifcloud.String("Base64EncodedString"),
					KeyStoragePolicy:      nifcloud.String("KeyStoragePolicy"),
					LicenseAcquisitionUrl: nifcloud.String("ZeroTo512String"),
					Method:                nifcloud.String("HlsContentProtectionMethod"),
				},
				Name: nifcloud.String("Filename"),
				OutputKeys: []*string{
					nifcloud.String("Key"), // Required
					// More values...
				},
				PlayReadyDrm: &elastictranscoder.PlayReadyDrm{
					Format:                nifcloud.String("PlayReadyDrmFormatString"),
					InitializationVector:  nifcloud.String("ZeroTo255String"),
					Key:                   nifcloud.String("NonEmptyBase64EncodedString"),
					KeyId:                 nifcloud.String("KeyIdGuid"),
					KeyMd5:                nifcloud.String("NonEmptyBase64EncodedString"),
					LicenseAcquisitionUrl: nifcloud.String("OneTo512String"),
				},
			},
			// More values...
		},
		UserMetadata: map[string]*string{
			"Key": nifcloud.String("String"), // Required
			// More values...
		},
	}
}
