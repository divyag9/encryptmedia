package asymmetric

import (
	"bytes"
	"testing"
)

var encryptCases = []struct {
	label   []byte
	message []byte
}{
	{[]byte("test"),
		[]byte{116, 101, 115, 116},
	},
}

func TestEncrypt(t *testing.T) {
	for _, encrypt := range encryptCases {
		privateKey, publicKey, _ := GenerateKeys()
		encryptedMessage, err := Encrypt(publicKey, encrypt.message, encrypt.label)
		if err != nil {
			t.Fatalf("%v", err)
		}

		decryptedMessage, err := Decrypt(privateKey, encryptedMessage, encrypt.label)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !bytes.Equal(encrypt.message, decryptedMessage) {
			t.Errorf("decrypted message returned:%q, expected:%q", decryptedMessage, encrypt.message)
		}
	}
}

func TestGenerateKeys(t *testing.T) {
	_, _, err := GenerateKeys()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
