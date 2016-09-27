package encrypt

import (
	"bytes"
	"testing"
)

var encryptCases = []struct {
	key     []byte
	message []byte
}{
	{[]byte{72, 64, 58, 173, 137, 74, 234, 185, 160, 212, 135, 134, 140, 113, 151, 157, 235, 251, 91, 65, 23, 173, 74, 73, 151, 120, 19, 39, 65, 254, 94, 36},
		[]byte{116, 101, 115, 116},
	},
}

func TestEncrypt(t *testing.T) {
	for _, encrypt := range encryptCases {
		encryptedMessage, err := Encrypt(encrypt.key, encrypt.message)
		if err != nil {
			t.Fatalf("%v", err)
		}

		decryptedMessage, err := Decrypt(encrypt.key, encryptedMessage)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !bytes.Equal(encrypt.message, decryptedMessage) {
			t.Errorf("decrypted message returned:%q, expected:%q", decryptedMessage, encrypt.message)
		}
	}
}

func TestGenerateKey(t *testing.T) {
	_, err := GenerateKey()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
