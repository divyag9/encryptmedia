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
	fileName, errDecode := mediastore.DecodeMedia(data)
	if errDecode != nil {
		log.Fatal("Failed to decode the bytes recieved: ", errDecode)
	}
	fmt.Println("File stored on disk: ", fileName)
}
