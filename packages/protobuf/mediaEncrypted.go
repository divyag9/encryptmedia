package protobuf

import (
	"github.com/divyag9/encryptmedia/packages"
	"github.com/golang/protobuf/proto"
)

// MarshalMediaEncrypted recieves the MediaEncrypted protobuf encodes to binary format
func MarshalMediaEncrypted(mediaEncrypted *encryptMedia.MediaEncrypted) ([]byte, error) {

	return proto.Marshal(&MediaEncrypted{
		Version:            mediaEncrypted.Version,
		GUID:               mediaEncrypted.GUID,
		Client:             mediaEncrypted.Client,
		LoanType:           mediaEncrypted.LoanType,
		OrderNumber:        mediaEncrypted.OrderNumber,
		UserName:           mediaEncrypted.UserName,
		Latitude:           mediaEncrypted.Latitude,
		Longitude:          mediaEncrypted.Longitude,
		DateTaken:          mediaEncrypted.DateTaken,
		DeviceModel:        mediaEncrypted.DeviceModel,
		DeviceOS:           mediaEncrypted.DeviceOS,
		DeviceOSVersion:    mediaEncrypted.DeviceOSVersion,
		FileName:           mediaEncrypted.FileName,
		MimeType:           mediaEncrypted.MimeType,
		Application:        mediaEncrypted.Application,
		ApplicationID:      mediaEncrypted.ApplicationID,
		ApplicationVersion: mediaEncrypted.ApplicationVersion,
		EncryptedBytes:     mediaEncrypted.EncryptedBytes,
		EncryptedKey:       mediaEncrypted.EncryptedKey,
		PrivateKey:         mediaEncrypted.PrivateKey,
	})
}

// UnmarshalMediaEncrypted decodes MediaEncrypted from binary data
func UnmarshalMediaEncrypted(data []byte, mediaEncrypted *encryptMedia.MediaEncrypted) (err error) {
	var pb MediaEncrypted
	err = proto.Unmarshal(data, &pb)
	mediaEncrypted.Version = pb.Version
	mediaEncrypted.GUID = pb.GUID
	mediaEncrypted.Client = pb.Client
	mediaEncrypted.LoanType = pb.LoanType
	mediaEncrypted.OrderNumber = pb.OrderNumber
	mediaEncrypted.UserName = pb.UserName
	mediaEncrypted.Latitude = pb.Latitude
	mediaEncrypted.Longitude = pb.Longitude
	mediaEncrypted.DateTaken = pb.DateTaken
	mediaEncrypted.DeviceModel = pb.DeviceModel
	mediaEncrypted.DeviceOS = pb.DeviceOS
	mediaEncrypted.DeviceOSVersion = pb.DeviceOSVersion
	mediaEncrypted.FileName = pb.FileName
	mediaEncrypted.MimeType = pb.MimeType
	mediaEncrypted.Application = pb.Application
	mediaEncrypted.ApplicationID = pb.ApplicationID
	mediaEncrypted.ApplicationVersion = pb.ApplicationVersion
	mediaEncrypted.EncryptedBytes = pb.EncryptedBytes
	mediaEncrypted.EncryptedKey = pb.EncryptedKey
	mediaEncrypted.PrivateKey = pb.PrivateKey

	return
}
