package main

import (
	"log"
	"net/http"
)

func newServer() {
	http.HandleFunc("/", getImageHandler)
	http.HandleFunc("/image", getImageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}