package httpHandler

import (
	"net/http"
	"questions/question2/src/movies"
)

var handlerMapper = map[string]http.HandlerFunc{
	"GET/movies":  movies.GetMovies,
	"GET/movies/": movies.GetMovieDetail,
}
