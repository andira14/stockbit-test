package response

import (
	"encoding/json"
	"net/http"
)

func InternalServerError(w http.ResponseWriter) {
	generalResponse(w, http.StatusInternalServerError,
		httpResponse{Code: "INTERNAL_SERVER_ERROR", Message: "Internal server error"},
	)
}

func BadRequest(w http.ResponseWriter, code string, message string) {
	generalResponse(w, http.StatusInternalServerError,
		httpResponse{Code: code, Message: message},
	)

}

func Ok(w http.ResponseWriter, data interface{}) {
	generalResponse(w, http.StatusOK,
		httpResponse{Code: "OK", Data: data},
	)
}

func generalResponse(w http.ResponseWriter, code int, response httpResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
