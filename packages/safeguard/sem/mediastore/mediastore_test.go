package mediastore

import (
	"testing"

	pb "github.com/divyag9/encryptmedia/packages/safeguard/sem/protobuf"
	"github.com/golang/protobuf/proto"
)

var mediaCases = []struct {
	data          []byte
	expectedFname string
}{
	{[]byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116,
		101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114, 3, 10, 20, 30, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116,
		101, 115, 116, 146, 1, 4, 116, 101, 115, 116}, "test.pack"},
	//{[]byte{5, 4, 7}, "error.pack"},
}

func TestDecodeMedia(t *testing.T) {
	for _, media := range mediaCases {
		fileName := DecodeMedia(media.data)
		if fileName != media.expectedFname {
			t.Errorf("FileName returned:%q, expectedFileName:%q", fileName, media.expectedFname)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	var data = []byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116,
		101, 115, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114, 3, 10, 20, 30, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116,
		101, 115, 116, 146, 1, 4, 116, 101, 115, 116}
	media := &pb.Media{}
	for n := 0; n < b.N; n++ {
		proto.Unmarshal(data, media)
	}
}
