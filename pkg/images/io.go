package images

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func LoadImage(fp string) (image.Image, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return im, nil
}

func ExportImage(fp string, im image.Image) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}

	ext := filepath.Ext(fp)
	switch ext {
	case ".jpg", ".jpeg":
		err := jpeg.Encode(f, im, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	case ".png":
		err := png.Encode(f, im)
		if err != nil {
			return err
		}
	}

	return nil
}
