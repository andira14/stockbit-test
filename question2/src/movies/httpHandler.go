package movies

import (
	"net/http"
	"questions/question2/src/httpHandler/response"
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

	result, err := getMovies(page, search)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	if strings.ToLower(result.Response) != "true" {
		response.BadRequest(w, "QUERY_ERROR", result.Error)
		return
	}

	response.Ok(w, result.Search)
}
