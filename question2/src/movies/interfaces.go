package movies

type data struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	Id        string `json:"imdbID"`
	MovieType string `json:"Type"`
	Poster    string `json:"Poster"`
}

type searchData struct {
	Search      []data `json:"Search"`
	ResultCount string `json:"totalResults"`
	Response    string `json:"Response"`
	Error       string `json:"Error,omitempty"`
}

type rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
type detail struct {
	Title     string   `json:"Title"`
	Year      string   `json:"Year"`
	Id        string   `json:"imdbID"`
	MovieType string   `json:"Type"`
	Poster    string   `json:"Poster"`
	Released  string   `json:"Released"`
	Genre     string   `json:"Genre"`
	Director  string   `json:"Director"`
	Writer    string   `json:"Writer"`
	Actors    string   `json:"Actors"`
	Ratings   []rating `json:"Ratings"`
	Response  string   `json:"Response"`
	Error     string   `json:"Error,omitempty"`
}
