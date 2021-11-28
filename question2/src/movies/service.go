package movies

import (
	"strings"
)

func getService(page int, search string) (searchData, error) {
	result, err := getRepo(page, search)
	if err != nil {
		return searchData{}, err
	}

	if strings.ToLower(result.Response) != "true" {
		return searchData{Error: result.Error}, nil
	}

	return result, nil
}

func getDetailService(imdbId string) (detail, error) {
	result, err := getDetailRepo(imdbId)
	if err != nil {
		return detail{}, err
	}

	if strings.ToLower(result.Response) != "true" {
		return detail{Error: result.Error}, nil
	}

	return result, nil
}
