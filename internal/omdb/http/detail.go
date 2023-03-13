package omdb_http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

func (o *OMDBHTTP) createDetailRequest(ctx context.Context, req omdb.DetailRequest) (*http.Request, error) {
	httpreq, err := http.NewRequestWithContext(ctx, "GET", o.url, nil)
	if err != nil {
		return nil, err
	}

	urlParam := httpreq.URL.Query()
	urlParam.Add("i", req.ID)
	httpreq.URL.RawQuery = urlParam.Encode()

	return httpreq, nil
}

func (o *OMDBHTTP) Detail(ctx context.Context, req omdb.DetailRequest) (*omdb.DetailResponse, error) {
	httpreq, err := o.createDetailRequest(ctx, req)
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

	return parseDetailResponse(body)
}

func parseDetailResponse(body []byte) (*omdb.DetailResponse, error) {
	ret := &omdb.DetailResponse{}
	err := json.Unmarshal(body, &ret)

	return ret, err
}
