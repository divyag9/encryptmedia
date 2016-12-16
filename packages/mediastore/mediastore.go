package mediastore

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/encrypt/asymmetric"
	"github.com/divyag9/encryptmedia/packages/encrypt/symmetric"
	"github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/pkg/errors"
)

// GetMediaEncryptedBytes returns the bytes of MediaEncrypted struct
func GetMediaEncryptedBytes(media *encryptMedia.Media, mediaEncrypted *encryptMedia.MediaEncrypted, pemFilePath string) ([]byte, error) {
	//Generate key for encryption of media bytes
	key, err := symmetric.GenerateKey()
	if err != nil {
		return nil, errors.Wrap(err, "Error generating symmetric key ")
	}
	// Encrypt the media bytes
	encryptedBytes, err := symmetric.Encrypt(key, media.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "Error encrypting media bytes ")
	}
	// Get the public key for encryption
	publicKey, err := getPublicKey(pemFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "Error retrieving public key ")
	}
	// Encrypt the symmetric key used to encrypt the media bytes
	encryptedKey, err := asymmetric.Encrypt(publicKey, key, []byte(""))
	if err != nil {
		return nil, errors.Wrap(err, "Error encrypting symmetric key ")
	}
	// Convert public key to bytes
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling public key ")
	}

	mediaEncrypted.Version = media.Version
	mediaEncrypted.GUID = media.GUID
	mediaEncrypted.Client = media.Client
	mediaEncrypted.LoanType = media.LoanType
	mediaEncrypted.OrderNumber = media.OrderNumber
	mediaEncrypted.UserName = media.UserName
	mediaEncrypted.Latitude = media.Latitude
	mediaEncrypted.Longitude = media.Longitude
	mediaEncrypted.DateTaken = media.DateTaken
	mediaEncrypted.DeviceModel = media.DeviceModel
	mediaEncrypted.DeviceOS = media.DeviceOS
	mediaEncrypted.DeviceOSVersion = media.DeviceOSVersion
	mediaEncrypted.FileName = media.FileName
	mediaEncrypted.MimeType = media.MimeType
	mediaEncrypted.Application = media.Application
	mediaEncrypted.ApplicationID = media.ApplicationID
	mediaEncrypted.ApplicationVersion = media.ApplicationVersion
	mediaEncrypted.EncryptedBytes = encryptedBytes
	mediaEncrypted.EncryptedKey = encryptedKey
	mediaEncrypted.PublicKey = publicKeyBytes
	// Marshal MediaEncrypted
	mediaEncryptedBytes, err := protobuf.MarshalMediaEncrypted(mediaEncrypted)

	return mediaEncryptedBytes, err
}

// getPublicKey reads the public.pem file and returns the public key
func getPublicKey(pemFilePath string) (*rsa.PublicKey, error) {
	// Does the PEM file even exist?
	if _, err := os.Stat(pemFilePath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, fmt.Sprintf("Pem file %s does not exist", pemFilePath))
	}

	// Load PEM
	pemFile, err := os.Open(pemFilePath)
	// need to convert pemfile to []byte for decoding
	pemFileInfo, _ := pemFile.Stat()
	size := pemFileInfo.Size()
	pemBytes := make([]byte, size)
	// read pemfile content into pembytes
	buffer := bufio.NewReader(pemFile)
	_, err = buffer.Read(pemBytes)
	publicKeyData, _ := pem.Decode([]byte(pemBytes))
	// retrive the public key from the bytes
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyData.Bytes)
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	return publicKey, err
}
