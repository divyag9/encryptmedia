package main

import (
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read request body: ", err)
	}

	media := &encryptMedia.Media{}
	// Decode the recieved media
	errUnmarshal := protobuf.UnmarshalMedia(data, media)
	if errUnmarshal != nil {
		log.Fatalln("Failed to decode Media:", errUnmarshal)
	}

	mediaEncrypted := &encryptMedia.MediaEncrypted{}
	// Get the media encrypted bytes and write to file on disk
	mediaEncryptedBytes, errBytes := mediastore.GetMediaEncryptedBytes(media, mediaEncrypted)
	if errBytes != nil {
		log.Fatalln("Failed to get media encrypted bytes: ", errBytes)
	}
	errSave := mediastore.SaveMediaEncrypted(mediaEncryptedBytes, fmt.Sprint(mediaEncrypted.GUID, ".sem"))
	if errSave != nil {
		log.Fatalln("Failed to write media encrypted bytes to disk: ", errSave)
	}
	fmt.Println("File stored on disk")
}
