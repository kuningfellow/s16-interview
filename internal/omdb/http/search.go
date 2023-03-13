package omdb_http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

func (o *OMDBHTTP) createSearchRequest(ctx context.Context, req omdb.SearchRequest) (*http.Request, error) {
	httpreq, err := http.NewRequestWithContext(ctx, "GET", o.url, nil)
	if err != nil {
		return nil, err
	}

	urlParam := httpreq.URL.Query()
	urlParam.Add("s", req.Query)
	httpreq.URL.RawQuery = urlParam.Encode()

	return httpreq, nil
}

func (o *OMDBHTTP) Search(ctx context.Context, req omdb.SearchRequest) (*omdb.SearchResponse, error) {
	httpreq, err := o.createSearchRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err := handleResponseError(o.RoundTrip(httpreq))
	if err != nil {
		return nil, err
	}

	body, err := readAndClose(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseSearchResponse(body)
}

func parseSearchResponse(body []byte) (*omdb.SearchResponse, error) {
	ret := &omdb.SearchResponse{}
	err := json.Unmarshal(body, &ret)

	return ret, err
}
