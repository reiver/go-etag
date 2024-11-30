package etag

import (
	"net/http"
	"strings"
)

// Handle handles the HTTP ETag & HTTP If-None-Match — it return true if HTTP request was handled, and returns false if it wasn't.
//
// Example usage:
//
//	func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
//
//		// ...
//
//		var handled bool = etag.Handle(responseWriter, request, eTag)
//		if handled {
//			return
//		}
//
//		// ...
//
//	}
func Handle(responseWriter http.ResponseWriter, request *http.Request, eTag string) bool {
	if nil == responseWriter {
		return false
	}
	if nil == request {
		httpError(responseWriter, http.StatusInternalServerError)
		return false
	}

	var responseHeader http.Header = responseWriter.Header()
	{
		if nil == responseHeader {
			httpError(responseWriter, http.StatusInternalServerError)
			return false
		}
	}

	var value string = `"` + eTag + `"`

	{
		responseHeader.Add("ETag", value)
	}

	var ifNoneMatch string
	{
		var requestHeader http.Header = request.Header
		if nil != requestHeader {
			ifNoneMatch = requestHeader.Get("If-None-Match")
		}
	}

	if "" != ifNoneMatch {
		switch {
		case "*" == ifNoneMatch:
			fallthrough
		case value == ifNoneMatch:
			fallthrough
		case strings.Contains(ifNoneMatch, value):
			responseWriter.WriteHeader(http.StatusNotModified)
			return true
		}
	}

	return false
}
