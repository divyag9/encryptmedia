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
func GetMediaEncryptedBytes(media *encryptMedia.Media, mediaEncrypted *encryptMedia.MediaEncrypted) ([]byte, error) {
	//Generate key for encryption of media bytes
	key, err := symmetric.GenerateKey()
	if err != nil {
		log.Fatalln("Error generating AES-256 key ", err)
		return nil, err
	}
	// Encrypt the media bytes
	encryptedBytes, err := symmetric.Encrypt(key, media.Bytes)
	if err != nil {
		log.Fatalln("Error encrypting media bytes ", err)
		return nil, err
	}
	// Generate public and private keys
	privateKey, publicKey, err := asymmetric.GenerateKeys()
	if err != nil {
		log.Fatalln("Error generating keys ", err)
		return nil, err
	}
	// Encrypt the key used to encrypt the media bytes
	encryptedKey, err := asymmetric.Encrypt(publicKey, key, []byte(""))
	if err != nil {
		log.Fatalln("Error encrypting key ", err)
		return nil, err
	}
	// Convert private key to bytes
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

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
	mediaEncrypted.PrivateKey = privateKeyBytes

	// Marshal MediaEncrypted
	mediaEncryptedBytes, err := protobuf.MarshalMediaEncrypted(mediaEncrypted)
	if err != nil {
		return nil, err
	}

	return mediaEncryptedBytes, nil
}

//SaveMediaEncrypted saves the encrypted media bytes to file on disk
func (ems EncryptedMediaService) SaveMediaEncrypted(mediaEncryptedBytes []byte, fileName string) error {
	ioutil.WriteFile(fileName, mediaEncryptedBytes, 0644)

	return nil
}
