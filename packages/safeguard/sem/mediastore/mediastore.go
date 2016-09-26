package mediastore

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/divyag9/encryptmedia/packages/safeguard/sem/protobuf"
	"github.com/golang/protobuf/proto"
)

//DecodeMedia takes the bytes, decodes to Media protobuf and writes to a file on disk
func DecodeMedia(data []byte) string {
	// Create the Media protobuf
	media := &pb.Media{}
	if err := proto.Unmarshal(data, media); err != nil {
		log.Fatalln("Failed to parse data:", err)
	}

	//Copying to another protobuf
	//What all should go into the newMedia buffer??
	newMedia := &pb.Media{}
	newMedia.Version = media.Version
	newMedia.GUID = media.GUID
	newMedia.Client = media.Client
	newMedia.LoanType = media.LoanType
	newMedia.OrderNumber = media.OrderNumber
	newMedia.UserName = media.UserName
	newMedia.Latitude = media.Latitude
	newMedia.Longitude = media.Longitude
	newMedia.DateTaken = media.DateTaken
	newMedia.DeviceModel = media.DeviceModel
	newMedia.DeviceOS = media.DeviceOS
	newMedia.DeviceOSVersion = media.DeviceOSVersion
	newMedia.FileName = media.FileName
	newMedia.ImageBytes = media.ImageBytes // will need to encrypt this
	newMedia.MimeType = media.MimeType
	newMedia.Application = media.Application
	newMedia.ApplicationID = media.ApplicationID
	newMedia.ApplicationVersion = media.ApplicationVersion

	//Write media back to disk
	fileName := fmt.Sprint(media.GUID, ".pack")
	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		log.Fatalln("Failed to write file to disk:", err)
	}
	return fileName
}
