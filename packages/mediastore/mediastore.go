package mediastore

import (
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/encrypt/asymmetric"
	"github.com/divyag9/encryptmedia/packages/encrypt/symmetric"
	"github.com/divyag9/encryptmedia/packages/protobuf"
)

// EncryptedMediaService struct manages MediaEncrypted
type EncryptedMediaService struct{}

// GetMediaEncryptedBytes returns the bytes of MediaEncrypted struct
func GetMediaEncryptedBytes(media *encryptMedia.Media, mediaEncrypted *encryptMedia.MediaEncrypted) (mediaEncryptedBytes []byte, err error) {
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
	// Generate public and private keys
	_, publicKey, err := asymmetric.GenerateKeys()
	if err != nil {
		log.Println("Error generating public and private keys ", err)
	}
	// Encrypt the key used to encrypt the media bytes
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
