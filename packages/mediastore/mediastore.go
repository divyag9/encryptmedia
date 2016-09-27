package mediastore

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/divyag9/encryptmedia/packages/encrypt"
	pb "github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/golang/protobuf/proto"
)

// DecodeMedia takes the bytes recieved, decodes to Media protobuf and writes the encrypted media to a file on disk
func DecodeMedia(data []byte) (string, error) {
	// Create the Media protobuf
	media := &pb.Media{}
	if err := proto.Unmarshal(data, media); err != nil {
		log.Fatalln("Failed to parse data:", err)
	}

	// Copying to the encrypted media protobuf
	encryptedMedia := &pb.MediaEncrypted{}
	encryptedMedia.Version = media.Version
	encryptedMedia.GUID = media.GUID
	encryptedMedia.Client = media.Client
	encryptedMedia.LoanType = media.LoanType
	encryptedMedia.OrderNumber = media.OrderNumber
	encryptedMedia.UserName = media.UserName
	encryptedMedia.Latitude = media.Latitude
	encryptedMedia.Longitude = media.Longitude
	encryptedMedia.DateTaken = media.DateTaken
	encryptedMedia.DeviceModel = media.DeviceModel
	encryptedMedia.DeviceOS = media.DeviceOS
	encryptedMedia.DeviceOSVersion = media.DeviceOSVersion
	encryptedMedia.FileName = media.FileName
	encryptedMedia.MimeType = media.MimeType
	encryptedMedia.Application = media.Application
	encryptedMedia.ApplicationID = media.ApplicationID
	encryptedMedia.ApplicationVersion = media.ApplicationVersion
	key, errKey := encrypt.GenerateKey()
	if errKey != nil {
		log.Fatalln("Error generating AES-256 key ", errKey)
	}
	encryptedMedia.SymmetricKey = key
	// Encrypt the media bytes
	encryptedBytes, errEncrypt := encrypt.Encrypt(key, media.Bytes)
	if errEncrypt != nil {
		log.Fatalln("Error encrypting media bytes ", errEncrypt)
	}
	encryptedMedia.EncryptedBytes = encryptedBytes

	// Marshal the encryptedMedia
	encryptedData, errMarshal := proto.Marshal(encryptedMedia)
	if errMarshal != nil {
		log.Fatalln("Error marshaling encryptedMedia: ", errMarshal)
	}
	// Write encrypted media back to disk
	fileName := fmt.Sprint(encryptedMedia.GUID, ".sem")
	if errWrite := ioutil.WriteFile(fileName, encryptedData, 0644); errWrite != nil {
		log.Fatalln("Failed to write file to disk:", errWrite)
	}
	return fileName, nil
}
