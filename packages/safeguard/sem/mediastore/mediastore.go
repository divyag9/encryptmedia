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
	//Write media back to disk
	fileName := fmt.Sprint(media.GUID, ".pack")
	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		log.Fatalln("Failed to write file to disk:", err)
	}
	return fileName
}
