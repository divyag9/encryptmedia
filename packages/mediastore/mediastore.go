package mediastore

import (
	"io/ioutil"
	"log"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/encrypt"
	"github.com/divyag9/encryptmedia/packages/protobuf"
)

// GetMediaEncryptedBytes returns the bytes of MediaEncrypted struct
func GetMediaEncryptedBytes(media *encryptMedia.Media, mediaEncrypted *encryptMedia.MediaEncrypted) ([]byte, error) {
	//Generate key for encryption of media bytes
	key, err := encrypt.GenerateKey()
	if err != nil {
		log.Fatalln("Error generating AES-256 key ", err)
		return nil, err
	}
	// Encrypt the media bytes
	encryptedBytes, errEncrypt := encrypt.Encrypt(key, media.Bytes)
	if errEncrypt != nil {
		log.Fatalln("Error encrypting media bytes ", errEncrypt)
		return nil, errEncrypt
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
	mediaEncrypted.SymmetricKey = key
	mediaEncrypted.EncryptedBytes = encryptedBytes

	// Marshal MediaEncrypted
	mediaEncryptedBytes, err := protobuf.MarshalMediaEncrypted(mediaEncrypted)
	if err != nil {
		return nil, err
	}

	return mediaEncryptedBytes, nil
}

//SaveMediaEncrypted saves the encrypted media bytes to file on disk
func SaveMediaEncrypted(mediaEncryptedBytes []byte, fileName string) error {
	ioutil.WriteFile(fileName, mediaEncryptedBytes, 0644)

	return nil
}
