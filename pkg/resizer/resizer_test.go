package resizer

import (
	"bytes"
	"os"
	"testing"
)

func TestResizePng(t *testing.T) {
	encodedImage := encodeImage()
	resizer := NewResizer()

	buff, err := resizer.Resize([]byte(encodedImage))
	if err != nil {
		t.Errorf("Failed resize: %s", err.Error())
	}

	outputImage(buff)
}

func outputImage(buff *bytes.Buffer) {
	f, err := os.Create("./testdata/output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(buff.Bytes())
	if err != nil {
		panic(err)
	}
}

func encodeImage() []byte {
	file, _ := os.Open("./testdata/gopher.png")
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	buff := make([]byte, size)
	file.Read(buff)

	return buff
}