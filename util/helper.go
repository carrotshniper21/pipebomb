// pipebomb/util/helper.go
package util

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error, message string, status int) {
	if err != nil {
		http.Error(w, message, status)
		return
	}
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	responseBytes, err := json.Marshal(data)
	if err != nil {
		HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBytes)
	if err != nil {
		HandleError(w, err, "Error writing response: ", http.StatusInternalServerError)
	}
}
