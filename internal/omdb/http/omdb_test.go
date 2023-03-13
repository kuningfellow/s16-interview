package omdb_http

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
)

func TestOMDBAuth(t *testing.T) {
	c, err := NewHTTPOMDB("https://www.omdbapi.com/", "put your key here")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
