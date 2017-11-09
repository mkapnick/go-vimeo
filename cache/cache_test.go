package cache_test

import (
	"github.com/mkapnick/go-vimeo/cache"
	_ "github.com/mkapnick/go-vimeo/cache"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_Init(t *testing.T) {
	client := cache.Client()
	config := client.ConfigGet("maxmemory")

	key := config.Val()[0].(string)
	value := config.Val()[1].(string)

	assert.Equal(t, key, "maxmemory")
	// about 64MB comes out to 67MB with redis :/
	assert.Equal(t, value, "67108864")
}

func Test_Set(t *testing.T) {
	val, _ := cache.Set("testkey", "testvalue")
	assert.Equal(t, true, val)
}

func Test_Get(t *testing.T) {
	val, _ := cache.Get("testkey")
	assert.Equal(t, "testvalue", val)
}

func Test_Info(t *testing.T) {
	val := cache.Info("Memory")
	arr := strings.Split(val, "\n")

	for index, entry := range arr {
		s := strings.Split(entry, ":")
		key := s[0]

		if index == 1 {
			assert.Equal(t, "used_memory", key)
		}
	}
}

// TODO make a test that goes over the cache limit
func Test_IsValidSize(t *testing.T) {
	isValid := cache.IsValidSize()
	assert.Equal(t, isValid, true)
}
