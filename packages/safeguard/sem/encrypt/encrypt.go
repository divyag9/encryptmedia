package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const (
	keySize   = 32
	nonceSize = 12
)

// GenerateKey generates a new AES-256 key.
func GenerateKey() ([]byte, error) {
	key := make([]byte, keySize)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}

	return key, nil
}

// GenerateNonce generates a new AES-GCM nonce.
func GenerateNonce() ([]byte, error) {
	nonce := make([]byte, nonceSize)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, err
	}

	return nonce, nil
}

// Encrypt secures a message using AES-GCM.
func Encrypt(key, message []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce, err := GenerateNonce()
	if err != nil {
		return nil, err
	}

	out := gcm.Seal(nonce, nonce, message, nil)
	return out, nil
}
