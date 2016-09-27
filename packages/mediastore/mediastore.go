package mediastore

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/divyag9/encryptmedia/packages/encrypt"
	pb "github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/golang/protobuf/proto"
)

// Media protobuf struct
type Media struct {
	*pb.Media
}

// MediaEncrypted protobuf struct
type MediaEncrypted struct {
	*pb.MediaEncrypted
}

// GetMedia takes the bytes recieved and decodes to Media protobuf
func GetMedia(data []byte) (*Media, error) {
	// Create the Media protobuf
	media := &Media{&pb.Media{}}
	if err := proto.Unmarshal(data, media.Media); err != nil {
		log.Fatalln("Failed to decode data:", err)
		return nil, err
	}
	return media, nil
}

// GetEncryptedMedia recieves the Media protobuf and creates the encrypted media
func (media *Media) GetEncryptedMedia() (*MediaEncrypted, error) {
	if media != nil {
		mediaEncrypted := &MediaEncrypted{&pb.MediaEncrypted{}}
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
		key, err := encrypt.GenerateKey()
		if err != nil {
			log.Fatalln("Error generating AES-256 key ", err)
			return nil, err
		}
		mediaEncrypted.SymmetricKey = key
		// Encrypt the media bytes
		encryptedBytes, errEncrypt := encrypt.Encrypt(key, media.Bytes)
		if errEncrypt != nil {
			log.Fatalln("Error encrypting media bytes ", errEncrypt)
			return nil, errEncrypt
		}
		mediaEncrypted.EncryptedBytes = encryptedBytes

		return mediaEncrypted, nil
	}

	return nil, errors.New("Media is nil")
}

// WriteEncryptedMediaToFile recieves the mediaEncrypted, encodes it and writes to file on disk
func (mediaEncrypted *MediaEncrypted) WriteEncryptedMediaToFile() (string, error) {
	if mediaEncrypted != nil {
		// Marshal the mediaEncrypted
		encryptedData, err := proto.Marshal(mediaEncrypted.MediaEncrypted)
		if err != nil {
			log.Fatalln("Error marshaling mediaEncrypted: ", err)
			return "", err
		}
		// Write encrypted media to disk
		fileName := fmt.Sprint(mediaEncrypted.GUID, ".sem")
		if errWrite := ioutil.WriteFile(fileName, encryptedData, 0644); errWrite != nil {
			log.Fatalln("Failed to write file to disk:", errWrite)
			return "", errWrite
		}
		return fileName, nil
	}
	return "", errors.New("MediaEncrypted is nil")
}
