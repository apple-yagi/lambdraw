package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"resize-api/pkg/resizer"
	"resize-api/pkg/s3"
	"testing"
)

type fakeS3Client struct {
	s3.S3

	FakePutImage func(key string, buff *bytes.Buffer) (string, error)
}

func (c *fakeS3Client) PutImage(key string, buff *bytes.Buffer) (string, error) {
	return c.FakePutImage(key, buff)
}

func TestExecute(t *testing.T) {
	c := &fakeS3Client{}
	h := NewHandler(c, &resizer.Resizer{})
	req := Request{}

	actual := h.Execute(req)
	expected := Response{
		StatusCode: 500,
		Body: "must request body",
	}
	if expected.StatusCode != actual.StatusCode || expected.Body != actual.Body {
		t.Errorf("actual: %v; expected: %v", actual, expected)
	}
	
	h.Client = &fakeS3Client{
		FakePutImage: func(key string, buff *bytes.Buffer) (string, error) {
			return "test", nil
		},
	}

	file, _ := os.Open("./testdata/gopher.png")
	defer file.Close()

	fi, _ := file.Stat()
	size := fi.Size()

	data := make([]byte, size)
	file.Read(data)

	enc := base64.StdEncoding.EncodeToString(data)

	req = Request{
		IsBase64Encoded: true,
		Body: enc,
	}
	
	actual = h.Execute(req)
	body, err := json.Marshal(map[string]interface{}{
		"url": "test",
	})
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	expected = *h.newResponse(buf.String(), nil)
	if expected.StatusCode != actual.StatusCode || expected.Body != actual.Body {
		t.Errorf("expected: %v; actual: %v", expected, actual)
	}
}