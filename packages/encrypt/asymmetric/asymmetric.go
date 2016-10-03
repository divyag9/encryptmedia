package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// GenerateKeys generates the private and public key for encryption/decryption
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	//Generate Private Key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println("privateKey", privateKey)
	// Precompute Calculations that speed up private key operations in the future
	privateKey.Precompute()
	//Validate Private Key
	err = privateKey.Validate()
	//Public key address (of an RSA key)
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, err
}

//Encrypt message using the public key
func Encrypt(publicKey *rsa.PublicKey, message []byte, label []byte) ([]byte, error) {
	hash := sha256.New()
	encryptedBytes, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)

	return encryptedBytes, err
}

//Decrypt the encryptedMessage using privateKey
func Decrypt(privateKey *rsa.PrivateKey, encryptedMessage []byte, label []byte) ([]byte, error) {
	hash := sha256.New()
	decryptedBytes, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, encryptedMessage, label)

	return decryptedBytes, err
}
