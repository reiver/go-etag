package etag

import (
        "net/http"
)

// httpError deals with writing a basic HTTP response error.
func httpError(responseWriter http.ResponseWriter, code int) {
	http.Error(responseWriter, http.StatusText(code), code)
}
