package source_test

import (
	"github.com/mkapnick/go-vimeo/source"
	"github.com/stretchr/testify/assert"
	"testing"
)

// should be valid on valid source
func Test_Initialize_Valid(t *testing.T) {
	Source, _ := source.New("http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4", "0")
	assert.Equal(t, Source.IsValid, true)
}

// should not be valid on invalid source
func Test_Initialize_Invalid(t *testing.T) {
	Source, _ := source.New("http://google.com", "0")
	assert.Equal(t, Source.IsValid, false)
}

// should throw an error on a bad source
func Test_Initialize_Error(t *testing.T) {
	_, err := source.New("http://i-am-a-bad-source", "0")

	assert.Equal(t, err.Error(), "Get http://i-am-a-bad-source: dial tcp: lookup i-am-a-bad-source: no such host")
}

// should fetch bytes from source url in the specified range
func Test_FetchBytesRange(t *testing.T) {
	Source, _ := source.New("http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4", "0-100")

	bytes, _ := Source.FetchBytes()

	assert.Equal(t, len(bytes), 101)
}

// should fetch all bytes from source url
// commenting this out since the test video is large and it
// takes 60 seconds to fetch all bytes
/*
func Test_FetchBytesAll(t *testing.T) {
	Source, _ := source.New("http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4", "0")

	bytes, _ := Source.FetchBytes()

	assert.Equal(t, len(bytes), 125433572)
}
*/

// fetch bytes from the source url in a specific range
func Test_FetchBytesRangeTwo(t *testing.T) {
	Source, _ := source.New("http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4", "500-800")

	bytes, _ := Source.FetchBytes()

	assert.Equal(t, len(bytes), 301)
}
