package omdb_http

import (
	"net/http"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

type OMDBHTTP struct {
	apiKey       string
	roundTripper http.RoundTripper
	url          string
}

var _ omdb.OMDB = (*OMDBHTTP)(nil) // interface satisfaction check
// NewHTTPOMDB creates a new HTTP OMDB client.
func NewHTTPOMDB(url string, apikey string) (*OMDBHTTP, error) {
	return &OMDBHTTP{
		apiKey:       apikey,
		roundTripper: http.DefaultTransport,
		url:          url,
	}, nil
}

func (o *OMDBHTTP) RoundTrip(req *http.Request) (*http.Response, error) {
	urlParam := req.URL.Query()
	urlParam.Set("apikey", o.apiKey)
	req.URL.RawQuery = urlParam.Encode()

	return o.roundTripper.RoundTrip(req)
}
