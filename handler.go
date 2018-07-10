package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
)

// getImage - Function called by the handler to get a random image.
func getImageHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer

	if err := generateImage(&buf); err != nil {
		log.Println("Unable to generate the image")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("Unable to write image.")
	}
}
