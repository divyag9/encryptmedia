package mediastore

import "testing"

var expectedFname = "test.sem"
var data = []byte{8, 1, 18, 4, 116, 101, 115, 116, 26, 4, 116, 101, 115, 116, 34, 4, 116, 101, 115, 116, 42, 4, 116, 101, 115, 116, 50, 4, 116, 101, 115, 116, 61, 0, 0, 128, 63, 69, 0, 0, 0, 64, 74, 4, 116, 101, 11, 5, 116, 82, 4, 116, 101, 115, 116, 90, 4, 116, 101, 115, 116, 98, 4, 116, 101, 115, 116, 106, 4, 116, 101, 115, 116, 114, 3, 10, 20, 30, 122, 4, 116, 101, 115, 116, 130, 1, 4, 116, 101, 115, 116, 138, 1, 4, 116, 101, 115, 116, 146, 1, 4, 116, 101, 115, 116}

func TestDecodeMedia(t *testing.T) {
	fileName := DecodeMedia(data)
	if fileName != expectedFname {
		t.Errorf("FileName returned:%q, expectedFileName:%q", fileName, expectedFname)
	}
}

// func BenchmarkDecodeMedia(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		data := []byte(strconv.Itoa(n))
// 		DecodeMedia(data)
// 	}
// }
