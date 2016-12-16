package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/divyag9/encryptmedia/packages"
	"github.com/divyag9/encryptmedia/packages/mediastore"
	"github.com/divyag9/encryptmedia/packages/protobuf"
	"github.com/spf13/viper"
)

func main() {
	config()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Setting up the configuration
func config() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}

// handle the incoming requet
func handler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read the request body", err)
	}

	media := &encryptMedia.Media{}
	// Decode the recieved media
	err = protobuf.UnmarshalMedia(reqBody, media)
	if err != nil {
		log.Println("Failed to decode Media: ", err)
	}
	mediaEncrypted := &encryptMedia.MediaEncrypted{}
	// Get the media encrypted bytes and write to file on disk
	mediaEncryptedBytes, err := mediastore.GetMediaEncryptedBytes(media, mediaEncrypted, viper.GetString("pemFilePath"))
	if err != nil {
		log.Println("Failed to get media encrypted bytes: ", err)
	}
	fmt.Fprintf(w, "%s", mediaEncryptedBytes)
}
