package handler

import (
	"encoding/base64"
	"os"
)

func NewSuccessRequest() *Request {
	file, _ := os.Open("./testdata/gopher.png")
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	enc := base64.StdEncoding.EncodeToString(data)

	return &Request{
		IsBase64Encoded: true,
		Body: enc,
	}
}

func NewEmptyRequest() *Request {
	return &Request{}
}