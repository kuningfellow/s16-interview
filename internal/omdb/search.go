package omdb

type SearchRequest struct {
	Query string
}

type SearchResponse struct {
	Search       []SearchItems
	TotalResults int `json:"totalResults,string"`
	Response     string
}

type SearchItems struct {
	Title  string
	Year   string
	IMDBID string `json:"imdbID"`
	Type   string
	Poster string
}
