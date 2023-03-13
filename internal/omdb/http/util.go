package omdb_http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

func handleResponseError(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, omdb.ErrorEmptyResponse
	}
	if err = HandleHttpError(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func HandleHttpError(resp *http.Response) error {
	// handle codes that we can handle
	switch resp.StatusCode {
	case http.StatusNotFound:
		return omdb.ErrorNotFound
	case http.StatusUnauthorized,
		http.StatusForbidden:
		return omdb.ErrorUnauthorized
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		return fmt.Errorf("%w: unexpected HTTP code %v", omdb.ErrorInternal, resp.StatusCode)
	}

	// we consider 4xx error as our internal error as it should not happen
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("%w: unexpected HTTP code %v", omdb.ErrorInternal, resp.StatusCode)
	}

	return fmt.Errorf("%w: external server returned %v", omdb.ErrorServerError, resp.StatusCode)
}

func readAndClose(r io.ReadCloser) ([]byte, error) {
	if r == nil {
		return nil, nil
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}
