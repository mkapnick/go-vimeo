package main

import (
	_ "github.com/mkapnick/go-vimeo/cache"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ServeBytes_InvalidSource(t *testing.T) {
	req, err := http.NewRequest("GET", "/serve", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeBytes)

	handler.ServeHTTP(res, req)

	status := res.Code
	message := res.Body.String()

	assert.Equal(t, status, 400)
	assert.Equal(t, message, "s query param is a required field\n")
}

func Test_ServeBytes_InvalidByte(t *testing.T) {
	req, err := http.NewRequest("GET", "/serve?s=test.com&range=bla", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeBytes)

	handler.ServeHTTP(res, req)

	status := res.Code
	message := res.Body.String()

	assert.Equal(t, status, 400)
	assert.Equal(t, message, "invalid byte range, must be a single byte number or a range like 0-100\n")
}

// invalid range 10*10
func Test_ServeBytes_InvalidByteRange(t *testing.T) {
	req, err := http.NewRequest("GET", "/serve?s=test.com&range=10*10", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeBytes)

	handler.ServeHTTP(res, req)

	status := res.Code
	message := res.Body.String()

	assert.Equal(t, status, 400)
	assert.Equal(t, message, "invalid byte range, must be a single byte number or a range like 0-100\n")
}

// invalid range 10-0
func Test_ServeBytes_InvalidByteRangeTwo(t *testing.T) {
	req, err := http.NewRequest("GET", "/serve?s=test.com&range=10-0", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeBytes)

	handler.ServeHTTP(res, req)

	status := res.Code
	message := res.Body.String()

	assert.Equal(t, status, 400)
	assert.Equal(t, message, "invalid byte range, must be a single byte number or a range like 0-100\n")
}

// invalid range 10-15-20
func Test_ServeBytes_InvalidByteRangeThree(t *testing.T) {
	req, err := http.NewRequest("GET", "/serve?s=test.com&range=10-15-20", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeBytes)

	handler.ServeHTTP(res, req)

	status := res.Code
	message := res.Body.String()

	assert.Equal(t, status, 400)
	assert.Equal(t, message, "invalid byte range, must be a single byte number or a range like 0-100\n")
}

// invalid range 10-bla
func Test_ServeBytes_InvalidByteRangeFour(t *testing.T) {
	req, err := http.NewRequest("GET", "/serve?s=test.com&range=10-bla", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeBytes)

	handler.ServeHTTP(res, req)

	status := res.Code
	message := res.Body.String()

	assert.Equal(t, status, 400)
	assert.Equal(t, message, "invalid byte range, must be a single byte number or a range like 0-100\n")
}
