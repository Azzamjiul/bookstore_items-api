package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/Azzamjiul/bookstore_utils-go/error_utils"
)

func ResponseAsJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseErrorAsJSON(w http.ResponseWriter, err *error_utils.RestErr) {
	ResponseAsJSON(w, err.Status, err)
}
