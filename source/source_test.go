package source_test

import (
	"github.com/mkapnick/go-vimeo/source"
	"github.com/stretchr/testify/assert"
	"testing"
)

// should be valid on valid source
func Test_Initialize_Valid(t *testing.T) {
	Source, _ := source.New("http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4")
	assert.Equal(t, Source.IsValid, true)
}

// should not be valid on invalid source
func Test_Initialize_Invalid(t *testing.T) {
	Source, _ := source.New("http://google.com")
	assert.Equal(t, Source.IsValid, false)
}

// should throw an error on a bad source
func Test_Initialize_Error(t *testing.T) {
	_, err := source.New("http://i-am-a-bad-source")

	assert.Equal(t, err.Error(), "Get http://i-am-a-bad-source: dial tcp: lookup i-am-a-bad-source: no such host")
}
