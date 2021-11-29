package movies

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testUrl = "https://www.omdbapi.com"
var testKey = "aaaa"

func TestImdbUrlSetup(t *testing.T) {
	os.Setenv("IMDB_URL", testUrl)
	os.Setenv("IMDB_KEY", testKey)
	testResult := fmt.Sprintf("%s?apikey=%s", testUrl, testKey)

	url := imdbUrlSetup()
	if url != testResult {
		t.Fatalf("Should return %s", testResult)
	}
}

func TestGetRepo(t *testing.T) {
	srv := serverMock(mockGetMovies)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	_, err := getRepo(5, "batman")
	if err != nil {
		t.Fatal("Should not return error")
	}
}

func TestGetRepoFalseResponse(t *testing.T) {
	srv := serverMock(mockFailedResponse)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	result, _ := getRepo(5, "")
	if result.Response != "False" {
		t.Fatal("Should return false response")
	}
	if result.Error != "Incorrect IMDb ID." {
		t.Fatal(`Should show Error "Incorrect IMDb ID."`)
	}
}

func TestGetDetailRepo(t *testing.T) {
	srv := serverMock(mockGetMovieDetail)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	_, err := getDetailRepo("asdfasdf")
	if err != nil {
		t.Fatal("Should not return error")
	}
}

func TestGetDetailRepoFalseResponse(t *testing.T) {
	srv := serverMock(mockFailedResponse)
	defer srv.Close()

	os.Setenv("IMDB_URL", srv.URL)
	os.Setenv("IMDB_KEY", testKey)

	result, _ := getDetailRepo("")
	if result.Response != "False" {
		t.Fatal("Should return false response")
	}
	if result.Error != "Incorrect IMDb ID." {
		t.Fatal(`Should show Error "Incorrect IMDb ID."`)
	}
}

func serverMock(mockHandle func(http.ResponseWriter, *http.Request)) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", mockHandle)

	srv := httptest.NewServer(handler)

	return srv
}

func mockGetMovies(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"Search":[{"Title":"Bat*21","Year":"1988","imdbID":"tt0094712","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZDRmNjYwZDktOTYxZi00MTdlLWI5ZjYtYWU4MDE5MDc5NGM3L2ltYWdlXkEyXkFqcGdeQXVyNjQzNDI3NzY@._V1_SX300.jpg"},{"Title":"The Bat","Year":"1959","imdbID":"tt0052602","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYWY3N2M0MzktZjkxNi00MjNlLTg2ZjctZGVjZTZhNzZiMDc4XkEyXkFqcGdeQXVyMDI2NDg0NQ@@._V1_SX300.jpg"},{"Title":"Siu ngo gong woo: Dung Fong Bat Bai","Year":"1992","imdbID":"tt0103295","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BN2FiN2RjOTUtNzRkYS00ZTM4LWJmYWUtOWRkNDM1NTdkMmQzXkEyXkFqcGdeQXVyNzgzODI1OTE@._V1_SX300.jpg"},{"Title":"Bat sin fan dim: Yan yuk cha siu bau","Year":"1993","imdbID":"tt0103743","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BM2EyNmU5ZGYtNjRiOC00NWNhLTgzOTctM2NlNWMwNzIyNjRlXkEyXkFqcGdeQXVyNzI1NzMxNzM@._V1_SX300.jpg"},{"Title":"The Devil Bat","Year":"1940","imdbID":"tt0032390","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZWE4MjhjMDItZmQ5Yy00OTQxLWE0M2EtZTJiMTFhMzc1NjJjXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"Title":"The Vampire Bat","Year":"1933","imdbID":"tt0024727","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BY2Q1NWZlOWQtOGQwMS00YjUyLTlkZDctNTQ5ZjRlNGE1ZDI1XkEyXkFqcGdeQXVyNjc0MzMzNjA@._V1_SX300.jpg"},{"Title":"The Bat People","Year":"1974","imdbID":"tt0071198","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZWM3ZGUyZWEtMjQ0OS00ZjJlLTkzNjktMjBjM2E1MzI0N2JlL2ltYWdlXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg"},{"Title":"Dung Fong Bat Bai: Fung wan joi hei","Year":"1993","imdbID":"tt0105655","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYTFiYWY2YWEtZWFmMy00MTJkLTlkZWItZGIwNWM0ZWUxYjZlXkEyXkFqcGdeQXVyNzgzODI1OTE@._V1_SX300.jpg"},{"Title":"Lost in Time","Year":"2003","imdbID":"tt0390272","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BOWE5ZDBjYTItOTNlNy00YjE4LWFlZjUtMTFiNmVhMzUyMTgwXkEyXkFqcGdeQXVyNzI1NzMxNzM@._V1_SX300.jpg"},{"Title":"Bat Thumb","Year":"2001","imdbID":"tt0331189","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTY3ODc3MjMzM15BMl5BanBnXkFtZTcwMjIwNzIyMQ@@._V1_SX300.jpg"}],"totalResults":"239","Response":"True"}`))
}

func mockGetMovieDetail(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"Title":"Bat*21","Year":"1988","Rated":"R","Released":"21 Oct 1988","Runtime":"105 min","Genre":"Drama, War","Director":"Peter Markle","Writer":"William C. Anderson, George Gordon","Actors":"Gene Hackman, Danny Glover, Jerry Reed","Plot":"During the Vietnam War, Colonel Hambleton's aircraft is shot down over enemy territory and a frantic rescue operation ensues.","Language":"English","Country":"United States","Awards":"1 nomination","Poster":"https://m.media-amazon.com/images/M/MV5BZDRmNjYwZDktOTYxZi00MTdlLWI5ZjYtYWU4MDE5MDc5NGM3L2ltYWdlXkEyXkFqcGdeQXVyNjQzNDI3NzY@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.5/10"},{"Source":"Rotten Tomatoes","Value":"81%"},{"Source":"Metacritic","Value":"58/100"}],"Metascore":"58","imdbRating":"6.5","imdbVotes":"8,574","imdbID":"tt0094712","Type":"movie","DVD":"24 Apr 2007","BoxOffice":"$3,966,256","Production":"N/A","Website":"N/A","Response":"True"}`))
}

func mockFailedResponse(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"Response":"False","Error":"Incorrect IMDb ID."}`))
}
