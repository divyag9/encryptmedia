package mediastore

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/golang/protobuf/proto"
)

var mediaCases = []struct {
	data                []byte
	mediaEncryptedBytes []byte
	pemFile             string
}{
	{[]byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116,
		61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116, 101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116,
		106, 4, 116, 101, 115, 116, 114, 4, 116, 101, 115, 116, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116, 101, 115,
		116, 146, 1, 4, 1, 2, 3, 4},
		[]byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101,
			115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116, 101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114,
			4, 116, 101, 115, 116, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116, 101, 115, 116, 146, 1, 32, 216, 8, 235, 91, 58, 159, 129, 23, 51, 184, 151, 38, 166, 8, 152, 7,
			6, 239, 103, 199, 201, 131, 213, 54, 48, 67, 226, 169, 144, 18, 37, 91, 160, 154, 1, 128, 2, 82, 15, 191, 205, 253, 215, 218, 30, 63, 195, 244, 223, 14, 147, 220, 157, 235, 225, 201, 94, 30,
			254, 48, 66, 159, 72, 193, 55, 139, 28, 15, 49, 73, 19, 233, 244, 90, 204, 254, 117, 153, 194, 248, 185, 5, 163, 237, 3, 170, 149, 170, 227, 79, 194, 75, 251, 111, 168, 229, 41, 193, 237, 18,
			6, 198, 235, 172, 183, 194, 182, 65, 193, 148, 184, 141, 173, 81, 5, 114, 118, 90, 16, 85, 41, 110, 255, 242, 193, 65, 73, 121, 86, 96, 165, 131, 53, 168, 115, 244, 132, 97, 202, 196, 243, 2,
			04, 90, 249, 90, 245, 160, 95, 234, 204, 27, 235, 214, 198, 168, 230, 45, 156, 228, 17, 61, 253, 77, 97, 160, 238, 102, 151, 20, 156, 43, 53, 149, 105, 154, 195, 46, 248, 204, 101, 56, 202,
			204, 53, 127, 205, 211, 45, 41, 160, 253, 112, 25, 151, 236, 112, 254, 176, 125, 171, 252, 191, 239, 45, 125, 15, 187, 253, 161, 127, 200, 65, 83, 202, 25, 55, 242, 111, 109, 202, 79, 122,
			12, 226, 25, 4, 63, 121, 98, 41, 206, 98, 185, 20, 213, 248, 33, 158, 61, 215, 248, 46, 32, 79, 216, 138, 218, 201, 196, 46, 249, 25, 69, 88, 210, 230, 155, 31, 60, 191, 64, 50, 250, 145, 112, 91,
			151, 15, 209, 251, 98, 100, 174, 125, 21, 132, 148, 169, 133, 236, 33, 246, 141, 81, 67, 8, 185, 3, 244, 187, 30, 72, 161, 203, 162, 1, 166, 2, 48, 130, 1, 34, 48, 13, 6, 9, 42, 134, 72, 1,
			34, 247, 13, 1, 1, 1, 5, 0, 3, 130, 1, 15, 0, 48, 130, 1, 10, 2, 130, 1, 1, 0, 185, 232, 119, 118, 146, 120, 150, 185, 139, 244, 248, 159, 2, 116, 101, 208, 141, 211, 23, 203, 118, 149, 247, 177, 231,
			216, 194, 195, 121, 14, 34, 123, 230, 62, 176, 176, 223, 94, 160, 216, 224, 130, 11, 178, 118, 50, 172, 90, 223, 239, 13, 3, 44, 167, 37, 140, 42, 106, 204, 199, 227, 11, 163, 242, 20,
			5, 100, 167, 219, 80, 193, 230, 241, 118, 254, 38, 185, 28, 183, 48, 229, 177, 117, 213, 37, 31, 97, 42, 149, 189, 207, 100, 177, 124, 222, 38, 214, 137, 26, 111, 57, 28, 96, 62, 192, 128, 1,
			06, 53, 18, 232, 10, 70, 222, 136, 85, 108, 51, 119, 122, 144, 172, 4, 141, 145, 181, 17, 104, 216, 100, 4, 226, 218, 182, 2, 248, 95, 52, 196, 4, 241, 243, 114, 223, 85, 23, 223, 231, 199, 2,
			50, 111, 126, 207, 17, 60, 186, 20, 168, 46, 79, 49, 29, 162, 153, 150, 144, 6, 212, 129, 78, 152, 89, 122, 162, 9, 36, 122, 171, 187, 221, 107, 77, 52, 194, 82, 65, 47, 250, 177, 71, 209, 53,
			242, 45, 55, 27, 171, 193, 199, 19, 221, 179, 174, 205, 46, 230, 87, 224, 42, 58, 248, 149, 36, 149, 20, 192, 101, 31, 178, 5, 167, 66, 244, 212, 158, 68, 154, 35, 99, 102, 254, 96, 151, 81,
			5, 97, 111, 80, 0, 209, 107, 78, 124, 117, 200, 64, 70, 237, 26, 250, 59, 159, 158, 27, 250, 88, 62, 151, 2, 3, 1, 0, 1},
		"../../cmd/server/public.pem",
	},
}

func TestSaveMediaEncrypted(t *testing.T) {
	for _, m := range mediaCases {
		media := &encryptMedia.Media{}
		protobuf.UnmarshalMedia(m.data, media)

		mediaEncrypted := &encryptMedia.MediaEncrypted{}
		mediaEncryptedBytes, _ := GetMediaEncryptedBytes(media, mediaEncrypted, m.pemFile)
		mediaEncryptedTest := &encryptMedia.MediaEncrypted{}
		protobuf.UnmarshalMediaEncrypted(mediaEncryptedBytes, mediaEncryptedTest)

		mediaEncryptedOrig := &encryptMedia.MediaEncrypted{}
		protobuf.UnmarshalMediaEncrypted(m.mediaEncryptedBytes, mediaEncryptedOrig)

		if mediaEncryptedTest.GUID != mediaEncryptedOrig.GUID {
			t.Errorf("MediaEncrypted GUID returned:%v, expected:%v", mediaEncryptedTest.GUID, mediaEncryptedOrig.GUID)
		}
		ems := EncryptedMediaService{}
		ems.SaveMediaEncrypted(mediaEncryptedBytes, fmt.Sprint(mediaEncrypted.GUID, ".sem"))

		//Read the contents of file and make sure the contents are same as original Media protobuf after decrypting
		bytesFile, _ := ioutil.ReadFile("test.sem")
		if !bytes.Equal(bytesFile, mediaEncryptedBytes) {
			t.Errorf("sem file bytes returned:%v, expected:%v", mediaEncryptedBytes, m.mediaEncryptedBytes)
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	media := &protobuf.Media{Version: 1,
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
	media := &protobuf.Media{}
	for n := 0; n < b.N; n++ {
		proto.Unmarshal(data, media)
	}
}
