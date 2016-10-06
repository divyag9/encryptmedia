package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/mediastore"
	"github.com/divyag9/encryptmedia/packages/protobuf"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// handle the incoming requet
func handler(w http.ResponseWriter, r *http.Request) {
	//Need to convert to bytes???
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	fmt.Print("byte array: ", bufBody)

	media := &encryptMedia.Media{}
	// Decode the recieved media
	err := protobuf.UnmarshalMedia(bufBody.Bytes(), media)
	if err != nil {
		log.Fatalln("Failed to decode Media: ", err)
	}
	mediaEncrypted := &encryptMedia.MediaEncrypted{}
	// Get the media encrypted bytes and write to file on disk
	mediaEncryptedBytes, err := mediastore.GetMediaEncryptedBytes(media, mediaEncrypted)
	if err != nil {
		log.Fatalln("Failed to get media encrypted bytes: ", err)
	}
	ems := mediastore.EncryptedMediaService{}
	saveMedia(ems, mediaEncryptedBytes, fmt.Sprint(mediaEncrypted.GUID, ".sem"))
}

// Function to save the media
func saveMedia(ms encryptMedia.MediaService, mediaEncryptedBytes []byte, fileName string) {
	err := ms.SaveMediaEncrypted(mediaEncryptedBytes, fileName)
	if err != nil {
		log.Fatalln("Failed to write media encrypted bytes to disk: ", err)
	}
	fmt.Println("File stored on disk")
}
