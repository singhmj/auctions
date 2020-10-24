package common

import (
	"io"
	"io/ioutil"
)

// ReadHTTPBody : Reads the content from body and returns the bytes and error
func ReadHTTPBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return ioutil.ReadAll(body)
}
