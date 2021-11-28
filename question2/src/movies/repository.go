package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"questions/question2/src/mapper"
)

func getMovies(page string, searchWord string) (searchData, error) {
	url := fmt.Sprintf("%s&s=%s&page=%s", imdbUrlSetup(), searchWord, page)
	resp, err := http.Get(url)
	if err != nil {
		return searchData{}, err
	}

	bodyBytes, jsonError := mapper.JsonToBytes(resp.Body)
	if jsonError != nil {
		return searchData{}, jsonError
	}

	var result searchData

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return searchData{}, err
	}

	return result, nil
}

func imdbUrlSetup() string {
	return fmt.Sprintf("%s?apikey=%s", os.Getenv("IMDB_URL"), os.Getenv("IMDB_KEY"))
}
