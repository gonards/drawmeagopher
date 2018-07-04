package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	var buf bytes.Buffer
	images := []string{
		"010-Body/blue_gopher.png",
		"020-Eyes/looking_up_no_lashes.png",
		"021-Shirts/pink_rainbow_shirt.png",
		"022-Hair/pink_hair_blue_ears.png",
		"023-Facial_Hair/full_ash_blonde_beard.png",
		"024-Glasses/movie_glasses.png",
		"025-Hats_Hair_Accessories/pirate_hat.png",
		"027-Extras/steampunk_glasses.png",
	}

	if err := generateimage(&buf, images...); err != nil {
		fmt.Println("An error occured")
	}

	saveimg(&buf)
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
	prefix := "artwork/" 
	imagesList := make([]image.Image, len(names))
	for i, name := range names {
		if len(name) == 0 {
			continue
		}
		fImg, _ := os.Open(prefix + name)
    	defer fImg.Close()
    	img, _ := png.Decode(fImg)
		imagesList[i] = img
		fImg.Close()
	}
	return imagesList
}