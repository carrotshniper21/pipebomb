// pipebomb/logging/logging.go
package logging

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request information
		color.Blue("Request method: %s", r.Method)
		color.Blue("Request URL: %s", r.URL.String())

		// Read and log the request body
		if r.Body != nil {
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}

			// Log the request body with pretty-printed JSON
			prettyBody := pretty.Color(pretty.Pretty(bodyBytes), pretty.TerminalStyle)
			color.Cyan("Request body: \n%s", string(prettyBody))

			// Replace the request body with a new reader, so it can be read again by the handlers
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Record the status code using a custom response writer
		recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)

		// Log the response status code
		color.Magenta("Response status: %d", recorder.statusCode)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
