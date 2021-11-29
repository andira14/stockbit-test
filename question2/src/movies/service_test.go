package movies

import (
	"os"
	"testing"
)

func TestGetService(t *testing.T) {
	srv := serverMock(mockGetMovies)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	result, err := getService(5, "batman")
	if err != nil {
		t.Fatal("Should not return error")
	}
	if result.Response != "True" {
		t.Fatal("Response should be true")
	}
	if len(result.Search) == 0 {
		t.Fatal("Search result should not be empty")
	}
}

func TestGetServiceFailedResponse(t *testing.T) {
	srv := serverMock(mockFailedResponse)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	result, err := getService(5, "batman")
	if err != nil {
		t.Fatal("Should not return error")
	}
	if result.Error == "" {
		t.Fatal("Error should have a message")
	}
	if len(result.Search) != 0 {
		t.Fatal("Search result should be empty")
	}
}

func TestGetDetailService(t *testing.T) {
	srv := serverMock(mockGetMovieDetail)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	result, err := getDetailService("aaaa")
	if err != nil {
		t.Fatal("Should not return error")
	}
	if result.Response == "False" {
		t.Fatal("Response should be True")
	}
}

func TestGetDetailServiceFailedResponse(t *testing.T) {
	srv := serverMock(mockFailedResponse)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	result, err := getDetailService("aaaa")
	if err != nil {
		t.Fatal("Should not return error")
	}
	if result.Response == "True" {
		t.Fatal("Response should be False")
	}
}
