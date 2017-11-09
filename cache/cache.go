package cache

import (
	"github.com/go-redis/redis"
	"strconv"
	"strings"
)

// client caches the frames using redis
var client *redis.Client

// MB size
const MB int = 1000000

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB,
	})

	// set the cache to a size of 64MB
	client.ConfigSet("maxmemory", "64MB")
}

// Client the redis client
func Client() *redis.Client {
	return client
}

// Set a key/value pair in the redis store
func Set(key string, value string) (bool, error) {
	err := client.Set(key, value, 0).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}

// Get a value from the store associated with a key
func Get(key string) (string, error) {
	val, err := client.Get(key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

// IsValidSize checks if the size of the cache is valid or not. In
// this example, the size of the cache must not be greater than 64MB
func IsValidSize() bool {
	byteSize := Size()
	mb := byteSize / MB

	if mb > 64 {
		return false
	}

	return true
}

// Size human readable size of the cache so far
func Size() int {
	info := Info("Memory")
	arr := strings.Split(info, "\n")
	size := 0

out:
	for _, entry := range arr {
		s := strings.Split(entry, ":")
		key := s[0]

		if key == "used_memory" {
			size, _ = strconv.Atoi(s[1])
			break out
		}
	}

	return size
}

// Info associated with redis db
func Info(key string) string {
	info := client.Info(key)
	return info.Val()
}
