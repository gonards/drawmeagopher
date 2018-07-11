package main

import (
	"net/http"
)

func newServer() {
	http.HandleFunc("/", getImageHandler)
	http.HandleFunc("/image", getImageHandler)
	addr := settings.ServerAddress
	if settings == nil {
		addr = ":8080"
	}
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		msg(err)
	}
}
