package util

import (
	"encoding/json"
	"fmt"
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
	HandleError(w, err, err.Error(), http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, s := w.Write(responseBytes)
	HandleError(w, s, "Error writing response: ", s)
}

func LoggingError(context string, err error) {
	if err != nil {
		fmt.Println("Error in", context, err)
	}
}
