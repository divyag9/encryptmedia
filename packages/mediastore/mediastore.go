package mediastore

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/encrypt/asymmetric"
	"github.com/divyag9/encryptmedia/packages/encrypt/symmetric"
	"github.com/divyag9/encryptmedia/packages/protobuf"
)

// EncryptedMediaService struct manages MediaEncrypted
type EncryptedMediaService struct{}

// GetMediaEncryptedBytes returns the bytes of MediaEncrypted struct
func GetMediaEncryptedBytes(media *encryptMedia.Media, mediaEncrypted *encryptMedia.MediaEncrypted, pemFilePath string) (mediaEncryptedBytes []byte, err error) {
	//Generate key for encryption of media bytes
	key, err := symmetric.GenerateKey()
	if err != nil {
		log.Println("Error generating symmetric key ", err)
	}
	// Encrypt the media bytes
	encryptedBytes, err := symmetric.Encrypt(key, media.Bytes)
	if err != nil {
		log.Println("Error encrypting media bytes ", err)
	}
	// Get the public key for encryption
	publicKey, err := getPublicKey(pemFilePath)
	if err != nil {
		log.Println("Error retrieving public key ", err)
		return nil, fmt.Errorf("Error retrieving public key: %s", err)
	}
	// Encrypt the symmetric key used to encrypt the media bytes
	encryptedKey, err := asymmetric.Encrypt(publicKey, key, []byte(""))
	if err != nil {
		log.Println("Error encrypting symmetric key ", err)
	}
	// Convert public key to bytes
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Println("Error marshalling public key ", err)
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
	mediaEncryptedBytes, err = protobuf.MarshalMediaEncrypted(mediaEncrypted)

	return
}

//SaveMediaEncrypted saves the encrypted media bytes to file on disk
func (ems EncryptedMediaService) SaveMediaEncrypted(mediaEncryptedBytes []byte, fileName string) (err error) {
	err = ioutil.WriteFile(fileName, mediaEncryptedBytes, 0644)

	return
}

// getPublicKey reads the public.pem file and returns the public key
func getPublicKey(pemFilePath string) (publicKey *rsa.PublicKey, err error) {
	// Does the PEM file even exist?
	if _, err = os.Stat(pemFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Pem file %s does not exist", pemFilePath)
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
	publicKey = publicKeyInterface.(*rsa.PublicKey)

	return
}
