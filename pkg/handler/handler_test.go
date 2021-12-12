package handler

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/apple-yagi/lambdraw/pkg/resizer"
	"github.com/apple-yagi/lambdraw/pkg/s3"
)

type fakeS3Client struct {
	s3.S3

	FakePutImage func(key string, buff *bytes.Buffer) (string, error)
}

func (c *fakeS3Client) PutImage(key string, buff *bytes.Buffer) (string, error) {
	return c.FakePutImage(key, buff)
}

func TestExecuteWhenSuccessRequest(t *testing.T) {
	c := &fakeS3Client{
		FakePutImage: func(key string, buff *bytes.Buffer) (string, error) {
			return "test", nil
		},
	}
	h := NewHandler(c, &resizer.Resizer{})
	req := NewSuccessRequest()

	actual, err := h.Execute(*req)
	if err != nil {
		panic(err)
	}

	body, err := json.Marshal(map[string]interface{}{
		"url": "test",
	})
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	expected := *h.newResponse(buf.String())
	if expected.StatusCode != actual.StatusCode || expected.Body != actual.Body {
		t.Errorf("expected: %v; actual: %v", expected, actual)
	}
}

func TestExecuteWhenEmptyRequest(t *testing.T) {
	c := &fakeS3Client{}
	h := NewHandler(c, &resizer.Resizer{})
	req := NewEmptyRequest()

	actual, err := h.Execute(*req)
	if err != nil {
		panic(err)
	}

	expected := Response{
		StatusCode: 500,
		Body:       "must request body",
	}
	if expected.StatusCode != actual.StatusCode || expected.Body != actual.Body {
		t.Errorf("actual: %v; expected: %v", actual, expected)
	}
}
