package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"questions/question2/src/mapper"
)

func getRepo(page int, searchWord string) (searchData, error) {
	url := fmt.Sprintf("%s&s=%s&page=%d", imdbUrlSetup(), searchWord, page)
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

func getDetailRepo(imdbId string) (detail, error) {
	url := fmt.Sprintf("%s&i=%s", imdbUrlSetup(), imdbId)
	resp, err := http.Get(url)
	if err != nil {
		return detail{}, err
	}

	bodyBytes, jsonError := mapper.JsonToBytes(resp.Body)
	if jsonError != nil {
		return detail{}, jsonError
	}

	var result detail

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return detail{}, err
	}

	return result, nil
}

func imdbUrlSetup() string {
	return fmt.Sprintf("%s?apikey=%s", os.Getenv("IMDB_URL"), os.Getenv("IMDB_KEY"))
}
