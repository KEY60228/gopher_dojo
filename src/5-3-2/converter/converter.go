package converter

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strings"
)

func IsValidExt(ext string) (bool, error) {
	match, err := regexp.MatchString("^(jpe?g|png|gif)$", ext)
	if err != nil {
		return false, err
	}
	return match, nil
}

func Do(path, at string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return err
	}
	result, err := IsValidExt(format)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("拡張子とフォーマットが異なります")
	}

	nfpath := path[:strings.LastIndex(path, ".")+1] + at
	nf, err := os.Create(nfpath)
	if err != nil {
		return err
	}
	defer nf.Close()

	switch at {
	case "jpeg", "jpg":
		err = jpeg.Encode(nf, img, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(nf, img)
	case "gif":
		err = gif.Encode(nf, img, nil)
	}
	if err != nil {
		return err
	}

	return nil
}
