package encrypt

import (
	"bytes"
	"testing"
)

var encryptCases = []struct {
	key              []byte
	message          []byte
	encryptedMessage []byte
}{
	{[]byte{72, 64, 58, 173, 137, 74, 234, 185, 160, 212, 135, 134, 140, 113, 151, 157, 235, 251, 91, 65, 23, 173, 74, 73, 151, 120, 19, 39, 65, 254, 94, 36},
		[]byte{116, 101, 115, 116},
		[]byte{123, 192, 40, 40, 129, 134, 82, 86, 98, 118, 232, 0, 148, 65, 48, 206, 236, 199, 195, 211, 77, 113, 92, 80, 238, 24, 230, 104, 205, 175, 43, 253},
	},
}

func TestEncrypt(t *testing.T) {
	for _, encrypt := range encryptCases {
		encryptedBytes, _ := Encrypt(encrypt.key, encrypt.message)
		if !bytes.Equal(encryptedBytes, encrypt.encryptedMessage) {
			t.Errorf("Encrypted bytes returned:%v, expected:%v", encryptedBytes, encrypt.encryptedMessage)
		}
	}
}
