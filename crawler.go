package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"
)

// Image struct
type Image struct {
	Href          string `json:"href"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	ThumbnailHref string `json:"thumbnail_href"`
}

// Category struct
type Category struct {
	ID     string  `json:"id"`
	Images []Image `json:"images"`
	Name   string  `json:"name"`
}

// Artwork struct
type Artwork struct {
	TotalCombinations int        `json:"totalCombination"`
	Categories        []Category `json:"categories"`
}

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

			clear()
			fmt.Println(category.Name, " [", j+1, "/", len(category.Images), "]")
			fmt.Println("\nDownloading... ", image.Href)

			err := downloadFile(image.ID, image.Href)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Calculate time elapsed
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		clear()
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

// Clear terminal (only for Linux)
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// Construct http client
	url := "https://gopherize.me/api/artwork/"
	c := http.Client{
		Timeout: time.Second * 10,
	}

	// Execute http call
	res, err := c.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	// Read result
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	// Parse artwork as json
	artwork, err := unmarshalArtwork(body)
	if err != nil {
		panic(err.Error())
	}

	// Refresh library by downloading images
	defer elapsed("Refreshing library")()
	err = refreshLibrary(artwork)
	if err != nil {
		panic(err.Error())
	}
}
