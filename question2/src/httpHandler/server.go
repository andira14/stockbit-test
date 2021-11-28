package httpHandler

import (
	"fmt"
	"net/http"
	"strings"
)

func Request(pattern string, method string, mux *http.ServeMux) {
	fmt.Println(pattern)
	mux.Handle(pattern, mapFunction(pattern, strings.ToUpper(method)))
}

func mapFunction(pattern string, method string) http.HandlerFunc {
	return handlerMapper[fmt.Sprintf("%s%s", method, pattern)]
}
