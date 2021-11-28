package mapper

import (
	"io"
	"io/ioutil"
)

func JsonToBytes(body io.ReadCloser) ([]byte, error) {

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
