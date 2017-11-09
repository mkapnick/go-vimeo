package main

import (
	"github.com/mkapnick/go-vimeo/cache"
	_ "github.com/mkapnick/go-vimeo/cache"
	"github.com/mkapnick/go-vimeo/source"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// isNumber checks if a string is a number
func isNumber(val string) (int, bool) {
	value, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}
	return value, true
}

// isValidByteRange validates the incoming byte range
func isValidByteRange(byteRange string) bool {
	_, isValid := isNumber(byteRange)
	if isValid {
		return true
	}

	// not a single number, validate the byte range
	arr := strings.Split(byteRange, "-")
	if len(arr) != 2 {
		return false
	}

	first, isValid := isNumber(arr[0])
	if !isValid {
		return false
	}
	second, isValid := isNumber(arr[1])
	if !isValid {
		return false
	}

	if first > second {
		return false
	}

	return true
}

// Root is the root route
func Root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Up and running")
}

// ServeBytes serves bytes from a source in a given range
// /serve?s=somesource.domain.com&range=20-50
// queryParams:
//  1) s: the source
//  2) range: the range of bytes to serve. End range is inclusive and
// 	optional
func ServeBytes(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	s := queryParams.Get("s")

	// throw a bad request if the client did not pass in a source
	if s == "" {
		http.Error(res, "s query param is a required field", 400)
	}

	byteRange := queryParams.Get("range")

	// if no range is given, set the default to the first 100 bytes 0-100
	if byteRange == "" {
		byteRange = "0-100"
	}

	isValid := isValidByteRange(byteRange)
	if !isValid {
		http.Error(res, "invalid byte range, must be a single byte number or a range like 0-100", 400)
	}

	// Step 1. Serve from the cache if stored
	id := s + "_" + byteRange
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
	Source, err := source.New(s, byteRange)
	if err != nil || !Source.IsValid {
		http.Error(res, "Invalid source url and/or byte range", 400)
	}

	// fetch the bytes from the source
	bytes, err := Source.FetchBytes()

	if err != nil {
		http.Error(res, "Error fetching bytes from the source server", 400)
	}

	// cache the bytes
	cache.Set(id, string(bytes))

	// send the bytes to the client!
	res.Write(bytes)
}

// main starts here. Creates a server listening on port 4000
func main() {

	// root route
	http.HandleFunc("/", Root)

	// the route that handles serving a range of bytes from a specific source
	http.HandleFunc("/serve", ServeBytes)

	// expose the web server on port 4000
	http.ListenAndServe(":4000", nil)
}
