package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sanity-io/litter"
)

// Unmarshal artwork from raw json
func unmarshalArtwork(body []byte) (*Artwork, error) {
	var artwork = new(Artwork)
	err := json.Unmarshal(body, &artwork)
	return artwork, err
}

// Download and override a file
func downloadFile(filepath string, url string) error {
	// Make sure parent directory exists
	os.MkdirAll(path.Dir(filepath), os.ModePerm)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	c := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := c.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

// Iterate over the collection and download files
func refreshLibrary(artwork *Artwork) error {
	for i := 0; i < len(artwork.Categories); i++ {
		category := artwork.Categories[i]
		for j := 0; j < len(category.Images); j++ {
			image := category.Images[j]

			msg(category.Name, " [", j+1, "/", len(category.Images), "]")
			msg("\nDownloading... ", image.Href)

			path := path.Join("artwork", category.Name, strings.Replace(image.Name, " ", "_", -1)+".png")
			litter.Dump(path)
			err := downloadFile(path, image.Href)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Calculate time elapsed
func elapsed() func() {
	start := time.Now()
	return func() {
		msg(fmt.Printf("%s took %v\n", "Crawling", time.Since(start)))
	}
}

func crawl() {
	// Construct http client
	url := "https://gopherize.me/api/artwork/"
	c := http.Client{
		Timeout: time.Second * 10,
	}

	// Execute http call
	res, err := c.Get(url)
	if err != nil {
		msg(err.Error())
		return
	}
	defer res.Body.Close()

	// Read result
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		msg(err.Error())
		return
	}

	// Parse artwork as json
	artwork, err := unmarshalArtwork(body)
	if err != nil {
		msg(err.Error())
		return
	}

	// Refresh library by downloading images
	defer elapsed()()
	err = refreshLibrary(artwork)
	if err != nil {
		msg(err.Error())
		return
	}
}
