package imageProcessor

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"
)

func ImageToBytes(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(f)
	return buf.Bytes()
}

func BytesToImage(data []byte, filepath string) error {
	reader := bytes.NewReader(data)

	image, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	opt := jpeg.Options{
		Quality: 50,
	}
	err = jpeg.Encode(f, image, &opt)
	if err != nil {
		return err
	}
	return nil
}
