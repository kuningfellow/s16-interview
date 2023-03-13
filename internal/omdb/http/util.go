package omdb_http

import (
	"io"
	"io/ioutil"
)

func readAndClose(r io.ReadCloser) ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}
