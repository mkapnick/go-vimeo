package main

import (
	"io"
	"net/http"
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
		http.Error(res, "Error", 400)
	}

	// byteRange := queryParams.Get("range")

	io.WriteString(res, source)
}

// validateSource validate that the incoming source can accept ranges
func validateSource() {

}

// cache caches the frames using redis
// TODO create separate dir to handle the caching mechanism
// specs: 64MB cache size
func cache() {

}

// main starts here. Creates a server listening on port 4000
func main() {
	http.HandleFunc("/", serveBytes)

	http.ListenAndServe(":4000", nil)
}
