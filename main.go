package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/divyag9/encryptmedia/packages/safeguard/sem/mediastore"
)

func main() {
	// media := &pb.Media{Version: 1,
	// 	GUID:               "test",
	// 	Client:             "test",
	// 	LoanType:           "test",
	// 	OrderNumber:        "test",
	// 	UserName:           "test",
	// 	Latitude:           1.0,
	// 	Longitude:          2.0,
	// 	DateTaken:          "test",
	// 	DeviceModel:        "test",
	// 	DeviceOS:           "test",
	// 	DeviceOSVersion:    "test",
	// 	FileName:           "test",
	// 	ImageBytes:         []byte{10, 20, 30},
	// 	MimeType:           "test",
	// 	Application:        "test",
	// 	ApplicationID:      "test",
	// 	ApplicationVersion: "test"}
	//
	// // decode Media and return the filename stored on disk
	// data, err := proto.Marshal(media)
	// if err != nil {
	// 	log.Fatal("marshaling error: ", err)
	// }
	// fmt.Println("data: ", data)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

// handle the incoming requet
func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Failed to read request body: ", err)
	}
	origChecksum := md5.Sum(data)
	fileName := mediastore.DecodeMedia(data)
	fmt.Println("File stored on disk: ", fileName)
	if getChecksum(fileName) == origChecksum {
		fmt.Println("File match")
	}
}

// Returns checksum of the file contents
func getChecksum(fileName string) [16]byte {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file: ", err)
	}
	checksum := md5.Sum(in)
	return checksum
}
