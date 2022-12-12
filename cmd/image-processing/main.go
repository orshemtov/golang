package main

import (
	"image"
	"log"
	"path"

	"github.com/orshemtov/golang/pkg/images"
)

func processImage(im image.Image) image.Image {
	// Do some processing here
	return im
}

func main() {
	im, err := images.LoadImage(path.Join("data", "images", "messi.jpg"))
	if err != nil {
		log.Fatal(err)
	}

	im = processImage(im)

	err = images.ExportImage(path.Join("data", "images", "messi.processed.jpg"), im)
	if err != nil {
		log.Fatal(err)
	}
}
