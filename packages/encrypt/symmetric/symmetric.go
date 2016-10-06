package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
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

	return key, err
}

// GenerateNonce generates a new AES-GCM nonce.
func GenerateNonce() ([]byte, error) {
	nonce := make([]byte, nonceSize)
	_, err := io.ReadFull(rand.Reader, nonce[:])

	return nonce, err
}

// Encrypt secures a message using AES-GCM.
func Encrypt(key, message []byte) (encryptedBytes []byte, err error) {
	c, err := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(c)
	nonce, err := GenerateNonce()
	encryptedBytes = gcm.Seal(nonce, nonce, message, nil)

	return
}

// Decrypt recovers a message secured using AES-GCM.
func Decrypt(key, message []byte) (decryptedBytes []byte, err error) {
	if len(message) <= nonceSize {
		return nil, errors.New("Decryption failed: message length is less than nonce size")
	}
	c, err := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(c)
	nonce := make([]byte, nonceSize)
	copy(nonce, message)
	decryptedBytes, err = gcm.Open(nil, nonce, message[nonceSize:], nil)

	return
}
