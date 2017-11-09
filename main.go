package main

import (
	"github.com/mkapnick/go-vimeo/cache"
	_ "github.com/mkapnick/go-vimeo/cache"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// root is the root route
func root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Up and running")
}

// serveBytes serves bytes from a source in a given range
// /serve?s=somesource.domain.com&range=20-50
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

	// if no range is given, set the default to the first 100 bytes 0-100
	if byteRange == "" {
		byteRange = "0-100"
	}

	// regex to verify the byte range format is in the form 10-20
	validRange := regexp.MustCompile(`0[1-9]|1[0-2]`)

	// numErrs is an error value that checks to ensure the byte range is valid
	// the byte range can either be a single number (20) or a byte range (0-100)
	numErrs := 0
	_, err := strconv.Atoi(byteRange)

	// byteRange is not a single number
	if err != nil {
		numErrs += 1
	}

	// byteRange is not a valid range
	if !validRange.MatchString(byteRange) {
		numErrs += 1
	}

	// throw a 400 if invalid byte range given
	if numErrs == 2 {
		http.Error(res, "invalid byte range, must be a single byte number or a range like 0-100", 400)
	}

	// Step 1. Serve from the cache if stored
	id := source + "_" + byteRange
	val, err := cache.Get(id)

	// if no error, then the result exists in the cache
	if err == nil {
		// convert the string back into an array of bytes
		res.Write([]byte(val))
	}

	// otherwise we need to do some grunt work
	// Step 1. Ensure the server accepts byte-ranges
	// Step 2. Ensure the byte range is valid
	// Step 3: Request the bytes from the server
	// Step 4: Cache the response if there is room
	// Step 5: Return the bytes
}

// main starts here. Creates a server listening on port 4000
func main() {

	// root route
	http.HandleFunc("/", root)

	// the route that handles serving a range of bytes from a specific source
	http.HandleFunc("/serve", serveBytes)

	// expose the web server on port 4000
	http.ListenAndServe(":4000", nil)
}
