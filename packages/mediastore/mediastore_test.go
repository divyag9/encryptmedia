package mediastore

import (
	"testing"

	pb "github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/golang/protobuf/proto"
)

var mediaCases = []struct {
	data          []byte
	expectedFname string
}{
	{[]byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116,
		101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114, 3, 10, 20, 30, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116,
		101, 115, 116, 146, 1, 4, 116, 101, 115, 116},
		"test.sem",
	},
}

var failureCases = []struct {
	data []byte
}{
	{[]byte{5, 4, 7}},
}

func TestWriteEncryptedMediaToFile(t *testing.T) {
	for _, m := range mediaCases {
		media, _ := GetMedia(m.data)
		encryptedMedia, _ := media.GetEncryptedMedia()
		fileName, _ := encryptedMedia.WriteEncryptedMediaToFile()
		if fileName != m.expectedFname {
			t.Errorf("File returned:%q, expected:%q", fileName, m.expectedFname)
		}
	}
}

// func TestDecodeMediaFailures(t *testing.T) {
// 	for _, media := range failureCases {
// 		_, err := DecodeMedia(media.data)
// 		if err != nil {
// 			t.Errorf("Error expected, but media decoded with no error")
// 		}
// 	}
// }

func BenchmarkMarshal(b *testing.B) {
	media := &pb.Media{Version: 1,
		GUID:               "test",
		Client:             "test",
		LoanType:           "test",
		OrderNumber:        "test",
		UserName:           "test",
		Latitude:           1.0,
		Longitude:          2.0,
		DateTaken:          "test",
		DeviceModel:        "test",
		DeviceOS:           "test",
		DeviceOSVersion:    "test",
		FileName:           "test",
		Bytes:              []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0},
		MimeType:           "test",
		Application:        "test",
		ApplicationID:      "test",
		ApplicationVersion: "test"}
	for n := 0; n < b.N; n++ {
		proto.Marshal(media)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	data := []byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116,
		101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114, 3, 10, 20, 30, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116,
		101, 115, 116, 146, 1, 4, 116, 101, 115, 116}
	media := &pb.Media{}
	for n := 0; n < b.N; n++ {
		proto.Unmarshal(data, media)
	}
}
