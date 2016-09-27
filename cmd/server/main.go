package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/divyag9/encryptmedia/packages/mediastore"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// handle the incoming requet
func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Failed to read request body: ", err)
	}
	media, _ := mediastore.GetMedia(data)
	encryptedMedia, _ := media.GetEncryptedMedia()
	fileName, err := encryptedMedia.WriteEncryptedMediaToFile()
	if err != nil {
		log.Fatal("Failed to write media to disk: ", err)
	}
	fmt.Println("File stored on disk: ", fileName)
}
