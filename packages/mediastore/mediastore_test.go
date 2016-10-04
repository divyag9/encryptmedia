package mediastore

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/encrypt/asymmetric"
	"github.com/divyag9/encryptmedia/packages/encrypt/symmetric"
	"github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/golang/protobuf/proto"
)

var mediaCases = []struct {
	data []byte
}{
	{[]byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116,
		61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116, 101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116,
		106, 4, 116, 101, 115, 116, 114, 4, 116, 101, 115, 116, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116, 101, 115,
		116, 146, 1, 4, 1, 2, 3, 4},
	},
}

func TestSaveMediaEncrypted(t *testing.T) {
	for _, m := range mediaCases {
		media := &encryptMedia.Media{}
		protobuf.UnmarshalMedia(m.data, media)

		mediaEncrypted := &encryptMedia.MediaEncrypted{}
		mediaEncryptedBytes, _ := GetMediaEncryptedBytes(media, mediaEncrypted)
		SaveMediaEncrypted(mediaEncryptedBytes, fmt.Sprint(mediaEncrypted.GUID, ".sem"))

		//Read the contents of file and make sure the contents are same as original Media protobuf after decrypting
		bytesFile, _ := ioutil.ReadFile("test.sem")

		mediaEncryptedTest := &encryptMedia.MediaEncrypted{}
		protobuf.UnmarshalMediaEncrypted(bytesFile, mediaEncryptedTest)
		privateKeyTest, _ := x509.ParsePKCS1PrivateKey(mediaEncryptedTest.PrivateKey)
		keyTest, _ := asymmetric.Decrypt(privateKeyTest, mediaEncryptedTest.EncryptedKey, []byte(""))
		decryptedBytes, _ := symmetric.Decrypt(keyTest, mediaEncryptedTest.EncryptedBytes)

		mediaTest := &encryptMedia.Media{}
		mediaTest.Version = mediaEncryptedTest.Version
		mediaTest.GUID = mediaEncryptedTest.GUID
		mediaTest.Client = mediaEncryptedTest.Client
		mediaTest.LoanType = mediaEncryptedTest.LoanType
		mediaTest.OrderNumber = mediaEncryptedTest.OrderNumber
		mediaTest.UserName = mediaEncryptedTest.UserName
		mediaTest.Latitude = mediaEncryptedTest.Latitude
		mediaTest.Longitude = mediaEncryptedTest.Longitude
		mediaTest.DateTaken = mediaEncryptedTest.DateTaken
		mediaTest.DeviceModel = mediaEncryptedTest.DeviceModel
		mediaTest.DeviceOS = mediaEncryptedTest.DeviceOS
		mediaTest.DeviceOSVersion = mediaEncryptedTest.DeviceOSVersion
		mediaTest.FileName = mediaEncryptedTest.FileName
		mediaTest.MimeType = mediaEncryptedTest.MimeType
		mediaTest.Application = mediaEncryptedTest.Application
		mediaTest.ApplicationID = mediaEncryptedTest.ApplicationID
		mediaTest.ApplicationVersion = mediaEncryptedTest.ApplicationVersion
		mediaTest.Bytes = decryptedBytes

		mediaTestBytes, _ := protobuf.MarshalMedia(mediaTest)
		if !bytes.Equal(mediaTestBytes, m.data) {
			t.Errorf("MediaBytes returned:%v, expected:%v", mediaTestBytes, m.data)
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
