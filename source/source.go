package source

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Source struct {
	Url           string
	Response      *http.Response
	IsValid       bool
	ContentLength string
	ByteRange     string
}

// New a new httpok source
func New(source string, byteRange string) (*Source, error) {
	return Initialize(source, byteRange)
}

// FetchBytes request wrapper
func (s *Source) FetchBytes() ([]byte, error) {
	req, err := http.NewRequest("GET", s.Url, nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%s", s.ByteRange))

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

// Initialize checks to see if the server accepts byte ranges
// and that the byte range is a valid range
func Initialize(source string, byteRange string) (*Source, error) {
	req, err := http.NewRequest("GET", source, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	rangeVal := res.Header.Get("Accept-Ranges")

	if rangeVal == "" || rangeVal != "bytes" {
		// source does not accept byte ranges
		return &Source{
			Url:      source,
			Response: res,
			IsValid:  false,
		}, nil
	}

	// confirm the byteRange is within the content length
	isValid := true
	contentLength := res.Header.Get("Content-Length")

	// TODO what if content length is a very large number?
	contentLengthInt, _ := strconv.Atoi(contentLength)
	ranges := strings.Split(byteRange, "-")

	if len(ranges) == 2 {
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])

		if start > contentLengthInt || end > contentLengthInt {
			isValid = false
		}
	} else {
		startByte, _ := strconv.Atoi(byteRange)
		if startByte > contentLengthInt {
			isValid = false
		}
		byteRange = byteRange + "-" + contentLength
	}

	return &Source{
		Url:           source,
		Response:      res,
		IsValid:       isValid,
		ContentLength: contentLength,
		ByteRange:     byteRange,
	}, nil
}
