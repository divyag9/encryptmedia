package protobuf

import (
	"github.com/divyag9/encryptmedia/packages"
	"github.com/golang/protobuf/proto"
)

// UnmarshalMedia decodes Media from binary data
func UnmarshalMedia(data []byte, media *encryptMedia.Media) error {
	var pb Media
	proto.Unmarshal(data, &pb)

	media.Version = pb.Version
	media.GUID = pb.GUID
	media.Client = pb.Client
	media.LoanType = pb.LoanType
	media.OrderNumber = pb.OrderNumber
	media.UserName = pb.UserName
	media.Latitude = pb.Latitude
	media.Longitude = pb.Longitude
	media.DateTaken = pb.DateTaken
	media.DeviceModel = pb.DeviceModel
	media.DeviceOS = pb.DeviceOS
	media.DeviceOSVersion = pb.DeviceOSVersion
	media.FileName = pb.FileName
	media.MimeType = pb.MimeType
	media.Application = pb.Application
	media.ApplicationID = pb.ApplicationID
	media.ApplicationVersion = pb.ApplicationVersion
	media.Bytes = pb.Bytes

	return nil
}

// MarshalMedia recieves the Media protobuf encodes to binary format
func MarshalMedia(media *encryptMedia.Media) ([]byte, error) {

	return proto.Marshal(&Media{
		Version:            media.Version,
		GUID:               media.GUID,
		Client:             media.Client,
		LoanType:           media.LoanType,
		OrderNumber:        media.OrderNumber,
		UserName:           media.UserName,
		Latitude:           media.Latitude,
		Longitude:          media.Longitude,
		DateTaken:          media.DateTaken,
		DeviceModel:        media.DeviceModel,
		DeviceOS:           media.DeviceOS,
		DeviceOSVersion:    media.DeviceOSVersion,
		FileName:           media.FileName,
		MimeType:           media.MimeType,
		Application:        media.Application,
		ApplicationID:      media.ApplicationID,
		ApplicationVersion: media.ApplicationVersion,
		Bytes:              media.Bytes,
	})
}
