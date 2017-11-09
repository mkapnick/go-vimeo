package main

import (
	"io"
	"net/http"
	"regexp"
)

// serveBytes serves bytes from a source in a given range
// /endpoint?s=somesource.domain.com&range=20-50
// queryParams:
//  1) s: the source
//  2) range: the range of bytes to serve. End range is inclusive and
// 	optional
func serveBytes(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	source := queryParams.Get("s")

	// throw a bad request if the client did not pass in a source
	if source == "" {
		http.Error(res, "s query param is a required field", 400)
	}

	byteRange := queryParams.Get("range")

	// if no range is given, set the default to 0-10
	if byteRange == "" {
		byteRange = "0-10"
	}

	// byte range must be in the format 10-20
	validRange := regexp.MustCompile(`0[1-9]|1[0-2]`)

	// return a 400 if the byte range is in an invalid format
	if !validRange.MatchString(byteRange) {
		http.Error(res, "invalid byte range format, should be 10-20", 400)
	}

	// Step 1. Serve from the cache to see if the source and frame range exist
	// in the cache

	io.WriteString(res, source)
}

// main starts here. Creates a server listening on port 4000
func main() {
	http.HandleFunc("/", serveBytes)

	// expose the web server on port 4000
	http.ListenAndServe(":4000", nil)
}
