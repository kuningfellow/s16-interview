package omdb

type DetailRequest struct {
	ID string
}

type DetailResponse struct {
	Title      string
	Year       int `json:",string"`
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []DetailRatings
	Metascore  string
	IMDBRating string `json:"imdbRating"`
	IMDBID     string `json:"imdbID"`
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}

type DetailRatings struct {
	Source string
	Value  string
}
