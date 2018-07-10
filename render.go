package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

// render - Function to render and save a random image
func render() {
	var buf bytes.Buffer

	if err := generateImage(&buf); err != nil {
		log.Println("Unable to generate the image")
	}

	saveImg(&buf)
}

// getImages - Get random images from folders.
// This imageswil l constitute the final image.
func getImages() []string {
	folders := getFolders()
	var images []string
	var file string
	fmt.Println("Picked Up Images :")
	for _, folder := range folders {
		file = getRandomFile(folder)
		images = append(images, file)
	}

	return images
}

// getFolders - Get folders associated to categories.
// The weight is used to implement the probability that a category is not used.
func getFolders() []string {
	var folders []string
	categories := Conf.Categories

	for _, category := range categories {
		rand.Seed(time.Now().UnixNano())
		nb := rand.Intn(10)
		if nb < category.Weight {
			folders = append(folders, "artwork/"+category.Name)
		}
	}

	return folders
}

// getRandomFile - Get a random file from a folder.
func getRandomFile(folder string) string {
	files := getFiles(folder)
	rand.Seed(time.Now().UnixNano())
	file := files[rand.Intn(len(files)-1)]
	fmt.Println("  - " + file)
	return folder + "/" + file
}

// getFiles - Get all the files in a folder.
func getFiles(folder string) []string {
	var files []string
	filesInfo, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range filesInfo {
		files = append(files, fileInfo.Name())
	}

	return files
}

// saveImg - Save the final Image to current directory.
func saveImg(r io.Reader) {
	finalImg, _ := os.Create("test.png")
	defer finalImg.Close()
	bytes, _ := ioutil.ReadAll(r)
	finalImg.Write(bytes)
	fmt.Println("Image Saved Properly")
}

// generateImage - Write the final image to a buffer
func generateImage(w io.Writer) error {
	images := getImages()
	imgObjects := loadImages(images...)
	var first image.Image
	for _, img := range imgObjects {
		if img == nil {
			continue
		}
		first = img
		break
	}

	output := image.NewRGBA(first.Bounds())
	for _, img := range imgObjects {
		if img == nil {
			continue
		}
		draw.Draw(output, output.Bounds(), img, image.ZP, draw.Over)
	}

	// encode into a buffer
	if err := png.Encode(w, output); err != nil {
		return err
	}

	return nil
}

// loadImages - Load images.
func loadImages(names ...string) []image.Image {
	imagesList := make([]image.Image, len(names))
	for i, name := range names {
		if len(name) == 0 {
			continue
		}
		fImg, _ := os.Open(name)
		defer fImg.Close()
		img, _ := png.Decode(fImg)
		imagesList[i] = img
	}

	return imagesList
}
