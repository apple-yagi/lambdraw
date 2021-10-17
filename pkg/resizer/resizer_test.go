package resizer

import (
	"bytes"
	"os"
	"testing"
)

func TestResizePng(t *testing.T) {
	encodedImage := encodeImage("./testdata/gopher.png")
	resizer := NewResizer()

	buff, err := resizer.Resize([]byte(encodedImage))
	if err != nil {
		t.Errorf("failed resize: %s", err.Error())
	}

	outputImage(buff, "./testdata/output/gopher.png")
}

func TestResizeJpeg(t *testing.T) {
	encodedImage := encodeImage("./testdata/gopher.jpg")
	resizer := NewResizer()

	buff, err := resizer.Resize([]byte(encodedImage))
	if err != nil {
		t.Errorf("failed resize: %s", err.Error())
	}

	outputImage(buff, "./testdata/output/gopher.jpg")
}

func outputImage(buff *bytes.Buffer, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(buff.Bytes())
	if err != nil {
		panic(err)
	}
}

func encodeImage(path string) []byte {
	file, _ := os.Open(path)
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	buff := make([]byte, size)
	file.Read(buff)

	return buff
}