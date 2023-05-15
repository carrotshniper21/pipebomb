// pipebomb/logging/logging.go
package logging

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/tidwall/pretty"
)

// LoggingMiddleware is a middleware function that logs the request and response details.
// It logs the request method, URL, body, and the response status code and body.
// The body of the request and response is pretty-printed for better readability.
// If the body is too long, it is truncated for brevity in the logs.
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
			if string(prettyBody) != "" {
				color.Cyan("Request body: \n%s", string(prettyBody))
			} else {
				color.Cyan("Request body: None\n")
			}
			// Replace the request body with a new reader, so it can be read again by the handlers
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Record the status code and body using a custom response writer
		recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)

		// Log the response status code
		color.Magenta("Response status: %d", recorder.statusCode)

		// Log the response body with pretty-printed JSON
		prettyBody := pretty.Color(pretty.Pretty(recorder.body.Bytes()), pretty.TerminalStyle)
		if len(prettyBody) > 0 {
			// Truncate the response body if it's too long
			const maxBodyLength = 1000
			prettyBodyStr := string(prettyBody)
			if len(prettyBodyStr) > maxBodyLength {
				prettyBodyStr = prettyBodyStr[:maxBodyLength] + "..."
			}
			color.Cyan("Response body: \n%s", prettyBodyStr)
		} else {
			color.Cyan("Response body: None\n")
		}
		color.Black("-------------------")
	})
}

// statusRecorder is a custom response writer that records the status code and body of the response.
// It embeds http.ResponseWriter and adds fields for the status code and body.
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

// WriteHeader sets the status code for the HTTP response and records it.
func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// Write writes the data to the connection as part of an HTTP reply and records the body.
func (r *statusRecorder) Write(body []byte) (int, error) {
	r.body.Write(body)
	return r.ResponseWriter.Write(body)
}
