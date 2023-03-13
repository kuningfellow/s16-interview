package omdb_http

import (
	"context"
	"fmt"
	"net/http/httputil"
	"testing"

	"github.com/kuningfellow/s16-interview/internal/omdb"
)

func TestDetailCreateRequest(t *testing.T) {
	c, err := NewHTTPOMDB("testendpoint", "secretkey")
	if err != nil {
		t.Fatal(err)
	}
	req, err := c.createDetailRequest(context.Background(), omdb.DetailRequest{
		ID: "someid",
	})
	if err != nil {
		t.Fatal(err)
	}

	b, err := httputil.DumpRequest(req, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestDetailParseResponse(t *testing.T) {
	response := `{
		"Title": "Batman: Under the Red Hood",
		"Year": "2010",
		"Rated": "PG-13",
		"Released": "27 Jul 2010",
		"Runtime": "75 min",
		"Genre": "Animation, Action, Crime, Drama, Mystery, Sci-Fi, Thriller",
		"Director": "Brandon Vietti",
		"Writer": "Judd Winick, Bob Kane (Batman created by)",
		"Actors": "Bruce Greenwood, Jensen Ackles, John DiMaggio, Neil Patrick Harris",
		"Plot": "There's a mystery afoot in Gotham City, and Batman must go toe-to-toe with a mysterious vigilante, who goes by the name of Red Hood. Subsequently, old wounds reopen and old, once buried memories come into the light.",
		"Language": "English",
		"Country": "USA",
		"Awards": "1 nomination.",
		"Poster": "https://m.media-amazon.com/images/M/MV5BNmY4ZDZjY2UtOWFiYy00MjhjLThmMjctOTQ2NjYxZGRjYmNlL2ltYWdlL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
		"Ratings": [
			{
				"Source": "Internet Movie Database",
				"Value": "8.1/10"
			},
			{
				"Source": "Rotten Tomatoes",
				"Value": "100%"
			}
		],
		"Metascore": "N/A",
		"imdbRating": "8.1",
		"imdbVotes": "56,401",
		"imdbID": "tt1569923",
		"Type": "movie",
		"DVD": "27 Jul 2010",
		"BoxOffice": "N/A",
		"Production": "N/A",
		"Website": "N/A",
		"Response": "True"
	}`

	resp, err := parseDetailResponse([]byte(response))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
