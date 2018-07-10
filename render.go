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

func render() {
	folders := []string{
		"artwork/010-Body",
		"artwork/020-Eyes",
		"artwork/021-Shirts",
		"artwork/022-Hair",
		"artwork/023-Facial_Hair",
		"artwork/024-Glasses",
		"artwork/025-Hats_and_Hair_Accessories",
		"artwork/027-Extras",
	}

	var images []string
	var file string

	fmt.Println("Picked Up Images :")
	for _, folder := range folders {
		file = getrandomfile(folder)
		images = append(images, file)
	}

	var buf bytes.Buffer

	if err := generateimage(&buf, images...); err != nil {
		fmt.Println("An error occured")
	}

	saveimg(&buf)
}

func getrandomfile(folder string) string {
	files := getfiles(folder)
	rand.Seed(time.Now().Unix())
	file := files[rand.Intn(len(files)-1)]
	fmt.Println("  - " + file)
	return folder + "/" + file
}

func getfiles(folder string) []string {
	var files []string
	filesInfo, err := ioutil.ReadDir("./" + folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range filesInfo {
		files = append(files, fileInfo.Name())
	}

	return files
}

func saveimg(r io.Reader) {
	finalimg, _ := os.Create("test.png")
	defer finalimg.Close()
	bytes, _ := ioutil.ReadAll(r)
	finalimg.Write(bytes)
}

func generateimage(w io.Writer, images ...string) error {
	imgObjects := loadimages(images...)
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

func loadimages(names ...string) []image.Image {
	imagesList := make([]image.Image, len(names))
	for i, name := range names {
		if len(name) == 0 {
			continue
		}
		fImg, _ := os.Open(name)
		defer fImg.Close()
		img, _ := png.Decode(fImg)
		imagesList[i] = img
		fImg.Close()
	}
	return imagesList
}
