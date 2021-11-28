package movies

type movieData struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	Id        string `json:"imdbID"`
	MovieType string `json:"Type"`
	Poster    string `json:"Poster"`
}

type searchData struct {
	Search      []movieData `json:"Search"`
	ResultCount string      `json:"totalResults"`
	Response    string      `json:"Response"`
	Error       string      `json:"Error"`
}
