package movies

import (
	"net/http"
	"questions/question2/src/httpHandler/response"
	"strconv"
	"strings"
)

func GetMovies(w http.ResponseWriter, req *http.Request) {
	queries := req.URL.Query()
	page := queries.Get("pagination")
	search := queries.Get("searchword")

	if page == "" {
		response.BadRequest(w, "EMPTY_PAGINATION", "Pagination cannot be empty")
		return
	}

	if search == "" {
		response.BadRequest(w, "EMPTY_SEARCH", "Searchword cannot be empty")
		return
	}

	convPage, convErr := strconv.Atoi(page)
	if convErr != nil {
		response.BadRequest(w, "INVALID_PAGE", "Page should be a number")
		return
	}

	result, err := getService(convPage, search)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	if result.Error != "" {
		response.BadRequest(w, "QUERY_ERROR", result.Error)
		return
	}

	response.Ok(w, result.Search)
}

func GetMovieDetail(w http.ResponseWriter, req *http.Request) {
	urlSplit := strings.Split(req.URL.String(), "/")
	if urlSplit[2] == "" {
		response.BadRequest(w, "EMPTY_MOVIEID", "Movie ID cannot be empty")
		return
	}
	movieId := urlSplit[2]

	result, err := getDetailService(movieId)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	if strings.ToLower(result.Response) != "true" {
		response.BadRequest(w, "QUERY_ERROR", "Incorrect ID")
		return
	}
	response.Ok(w, result)
}
